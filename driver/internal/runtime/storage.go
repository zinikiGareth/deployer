package runtime

import (
	"fmt"
	"io"

	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/pluggable"
)

type Storage struct {
	registry pluggable.Recall
	repo     pluggable.Repository
	sink     errorsink.ErrorSink
	mode     int
	drivers  map[string]any
	runtime  map[pluggable.Describable]any
}

func (s *Storage) Bind(v pluggable.Describable, value any) {
	s.runtime[v] = value
}

func (s *Storage) Get(v pluggable.Var) any {
	// log.Printf("have %v with %v\n", v, v.Binding())
	return s.runtime[v.Binding()]
}

func (s *Storage) Read(name pluggable.SymbolName) any {
	// log.Printf("read %v\n", name)
	return s.repo.GetDefinition(name)
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
	if e == nil {
		return nil
	}
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

func (s *Storage) DumpTo(w io.Writer) {
	fmt.Fprintf(w, "#keys = %d\n", len(s.runtime))
	for k, v := range s.runtime {
		fmt.Fprintf(w, "  var %s = %v\n", k, v)
	}
}

func NewRuntimeStorage(registry pluggable.Recall, repo pluggable.Repository, sink errorsink.ErrorSink) pluggable.RuntimeStorage {
	ret := &Storage{sink: sink, registry: registry, repo: repo, drivers: make(map[string]any), runtime: make(map[pluggable.Describable]any)}
	return ret
}
