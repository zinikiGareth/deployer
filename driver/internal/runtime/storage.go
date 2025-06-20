package runtime

import (
	"fmt"
	"io"
	"log"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/utils"
)

type Storage struct {
	registry driverbottom.Recall
	repo     driverbottom.Repository
	sink     errorsink.ErrorSink
	mode     int
	drivers  map[string]any
	runtime  map[driverbottom.Describable]any
}

func (s *Storage) Bind(v driverbottom.Describable, value any) {
	s.runtime[v] = value
}

func (s *Storage) Get(v driverbottom.Var) any {
	// log.Printf("have %v with %v\n", v, v.Binding())
	return s.runtime[v.Binding()]
}

func (s *Storage) Read(name driverbottom.SymbolName) any {
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

func (s *Storage) Eval(e driverbottom.Expr) any {
	if e == nil {
		return nil
	}
	return e.Eval(s)
}

func (s *Storage) EvalAsStringer(e driverbottom.Expr) fmt.Stringer {
	val := s.Eval(e)
	str, ok := val.(string)
	if ok {
		return utils.AsStringer(str)
	}
	stok, ok := val.(driverbottom.String)
	if ok {
		return stok
	}
	stringer, ok := val.(fmt.Stringer)
	if ok {
		return stringer
	}
	return utils.AsStringer(fmt.Sprintf("%v", val))
}

func (s *Storage) EvalAsNumber(e driverbottom.Expr) driverbottom.AsNumber {
	val := s.Eval(e)
	str, ok := val.(float64)
	if ok {
		return utils.F64AsNumber(str)
	}
	ntok, ok := val.(driverbottom.Number)
	if ok {
		return ntok
	}
	conv, ok := val.(driverbottom.AsNumber)
	if ok {
		return conv
	}
	log.Fatalf("cannot convert to AsNumber: %T", val)
	return nil
}

func (s *Storage) DumpTo(w io.Writer) {
	fmt.Fprintf(w, "#keys = %d\n", len(s.runtime))
	for k, v := range s.runtime {
		fmt.Fprintf(w, "  var %s = %v\n", k, v)
	}
}

func NewRuntimeStorage(registry driverbottom.Recall, repo driverbottom.Repository, sink errorsink.ErrorSink) driverbottom.RuntimeStorage {
	ret := &Storage{sink: sink, registry: registry, repo: repo, drivers: make(map[string]any), runtime: make(map[driverbottom.Describable]any)}
	return ret
}
