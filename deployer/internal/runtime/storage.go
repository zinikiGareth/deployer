package runtime

import (
	"log"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Storage struct {
	registry pluggable.Recall
	sink     errors.ErrorSink
	mode     int
	drivers  map[string]any
	actions  map[pluggable.Executable]pluggable.ExecuteAction
	runtime  map[string]any
}

func (s *Storage) Bind(name pluggable.SymbolName, value any) {
	s.runtime[string(name)] = value
}

func (s *Storage) Get(name pluggable.SymbolName) any {
	return s.runtime[string(name)]
}

func (s *Storage) Errorf(loc *errors.Location, msg string, args ...any) {
	s.sink.Reportf(loc, msg, args...)
}

func (s *Storage) SetMode(mode int) {
	s.mode = mode
}

func (s *Storage) IsMode(mode int) bool {
	return s.mode == mode
}

func (s *Storage) ObtainDriver(forType string) any {
	ret := s.drivers[forType]
	if ret != nil {
		return ret
	}
	c := s.registry.ObtainDriver(forType)
	if c == nil {
		panic("there is no driver for " + forType)
	}
	im, ok := c.(pluggable.InitMe)
	if ok {
		ret = im.InitMe(s)
	} else {
		ret = c
	}
	s.drivers[forType] = ret
	return ret
}

func (s *Storage) BindAction(a pluggable.Executable, av pluggable.ExecuteAction) {
	s.actions[a] = av
}

func (s *Storage) RetrieveAction(a pluggable.Executable) pluggable.ExecuteAction {
	return s.actions[a]
}

func (s *Storage) Eval(e pluggable.Expr) any {
	str, ok := e.(pluggable.String)
	if ok {
		return str.Text()
	} else {
		id, ok := e.(pluggable.Identifier)
		if ok {
			return s.Get(pluggable.SymbolName(id.Id()))
		} else {
			log.Fatalf("cannot evaluate %v", e)
			return nil
		}
	}
}

func NewRuntimeStorage(registry pluggable.Recall, sink errors.ErrorSink) pluggable.RuntimeStorage {
	return &Storage{sink: sink, actions: make(map[pluggable.Executable]pluggable.ExecuteAction), registry: registry, drivers: make(map[string]any), runtime: make(map[string]any)}
}
