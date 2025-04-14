package runtime

import (
	"reflect"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Storage struct {
	sink    errors.ErrorSink
	mode    int
	actions map[pluggable.Executable]pluggable.ExecuteAction
}

func (s *Storage) Bind(name pluggable.SymbolName, value any) {
}

func (s *Storage) Errorf(loc pluggable.Location, msg string, args ...any) {
	s.sink.Reportf(loc.Line, loc.Offset, "", msg, args...)
}

func (s *Storage) SetMode(mode int) {
	s.mode = mode
}

func (s *Storage) IsMode(mode int) bool {
	return s.mode == mode
}

func (s *Storage) ObtainDriver(forType reflect.Type) any {
	ret := reflect.New(forType).Interface()
	im, ok := ret.(pluggable.InitMe)
	if ok {
		im.InitMe(s)
	}
	return ret
}

func (s *Storage) BindAction(a pluggable.Executable, av pluggable.ExecuteAction) {
	s.actions[a] = av
}

func (s *Storage) RetrieveAction(a pluggable.Executable) pluggable.ExecuteAction {
	return s.actions[a]
}

func NewRuntimeStorage(sink errors.ErrorSink) pluggable.RuntimeStorage {
	return &Storage{sink: sink, actions: make(map[pluggable.Executable]pluggable.ExecuteAction)}
}
