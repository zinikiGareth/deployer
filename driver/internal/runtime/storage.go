package runtime

import (
	"fmt"
	"io"
	"log"
	"maps"
	"slices"

	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/utils"
)

type SymbolProvenance struct {
	to driverbottom.Identifier
}

func NewProvenance(to driverbottom.Identifier) *SymbolProvenance {
	return &SymbolProvenance{to: to}
}

type Storage struct {
	registry    driverbottom.Recall
	repo        driverbottom.Repository
	sink        errorsink.ErrorSink
	mode        int
	currentStep string
	drivers     map[string]any
	runtime     map[driverbottom.Describable]any
	symbols     map[string]map[string]*SymbolProvenance
}

func (s *Storage) Bind(v driverbottom.Describable, value any) {
	// So I think the steps here are:
	// 1. For the describable, convert it into a provenance
	// 2. That should have some version that exists for this step in resolve or something
	// 3. The value (in %p form) should not be anywhere in our memory, because that would be an update
	// 4. Keep track of it ...
	s.runtime[v] = value
}

func (s *Storage) Get(v driverbottom.Var) any {
	// TODO: here we are looking for the "correct" version of the VAR, which must be the one that was in operation at the time in this mode
	// I think that may be quite complicated.
	// We also need to think about what it means to "find" and "desire" two different things.
	return s.runtime[v.Binding()]
}

func (s *Storage) Read(name driverbottom.SymbolName) any {
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

func (s *Storage) EvalAsStringer(e driverbottom.Expr) (fmt.Stringer, bool) {
	val := s.Eval(e)
	str, ok := val.(string)
	if ok {
		return utils.AsStringer(str)
	}
	stringer, ok := val.(fmt.Stringer)
	if ok {
		return stringer, true
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

func (s *Storage) SetStepName(stepName string) {
	s.currentStep = stepName
	// log.Printf("mode %d: set step name to %s\n", s.mode, s.currentStep)
	s.symbols[s.currentStep] = make(map[string]*SymbolProvenance)
}

func (s *Storage) EnableSymbol(to driverbottom.Identifier) {
	if s.mode != 0 {
		panic("can only enable symbols during resolution")
	}
	if s.symbols[s.currentStep][to.Id()] != nil {
		panic("cannot define symbol more than once: " + to.Id())
	}
	// log.Printf("adding %s to %s\n", to.Id(), s.currentStep)
	s.symbols[s.currentStep][to.Id()] = NewProvenance(to)
}

func (s *Storage) ExportSymbolsTo(iw driverbottom.IndentWriter) {
	keys := slices.Collect(maps.Keys(s.symbols))
	slices.Sort(keys)
	for _, k := range keys {
		iw.Intro("Step %s\n", k)
		iw.Indent()
		syms := s.symbols[k]
		for _, p := range syms {
			iw.IndPrintf("%s\n", p.to.Id())
		}
		iw.UnIndent()
	}
}

func (s *Storage) NewObjId(loc *errorsink.Location) driverbottom.Identifier {
	l := len(s.symbols[s.currentStep])
	kk := fmt.Sprintf("*%s-%d", s.currentStep, l)
	id := lexicator.NewIdentifierToken(loc.Line, loc.Offset, kk)
	s.symbols[s.currentStep][kk] = NewProvenance(id)
	return id
}

func NewRuntimeStorage(registry driverbottom.Recall, repo driverbottom.Repository, sink errorsink.ErrorSink) driverbottom.RuntimeStorage {
	ret := &Storage{sink: sink, registry: registry, repo: repo, drivers: make(map[string]any), runtime: make(map[driverbottom.Describable]any), symbols: make(map[string]map[string]*SymbolProvenance)}
	return ret
}
