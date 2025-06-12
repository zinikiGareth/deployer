package runtime

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Storage struct {
	registry pluggable.Recall
	sink     errorsink.ErrorSink
	mode     int
	drivers  map[string]any
	runtime  map[pluggable.Describable]any
}

func (s *Storage) Bind(v pluggable.Describable, value any) {
	s.runtime[v] = value
}

func (s *Storage) Get(v pluggable.Var) any {
	return s.runtime[v.Binding()]
}

func (s *Storage) Errorf(loc *errorsink.Location, msg string, args ...any) {
	s.sink.Reportf(loc, msg, args...)
}

func (s *Storage) SetMode(mode int) {
	s.mode = mode
}

func (s *Storage) IsMode(mode int) bool {
	return s.mode == mode
}

func (s *Storage) Eval(e pluggable.Expr) any {
	return e.Eval(s)
}

func (s *Storage) EvalAsString(e pluggable.Expr) string {
	val := s.Eval(e)
	str, ok := val.(string)
	if ok {
		return str
	}
	stok, ok := val.(pluggable.String)
	if ok {
		return stok.Text()
	}
	stringer, ok := val.(fmt.Stringer)
	if ok {
		return stringer.String()
	}
	return fmt.Sprintf("%v", val)
}

func NewRuntimeStorage(registry pluggable.Recall, sink errorsink.ErrorSink) pluggable.RuntimeStorage {
	return &Storage{sink: sink, registry: registry, drivers: make(map[string]any), runtime: make(map[pluggable.Describable]any)}
}
