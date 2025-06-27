package runtime

import (
	"fmt"
	"io"
	"log"
	"maps"
	"slices"

	"ziniki.org/deployer/driver/internal/parser/interpreters"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/utils"
)

type SymbolProvenance struct {
	to     driverbottom.Identifier
	values map[int]map[string]any
}

func NewProvenance(to driverbottom.Identifier) *SymbolProvenance {
	return &SymbolProvenance{to: to, values: make(map[int]map[string]any)}
}

type Storage struct {
	registry    driverbottom.Recall
	repo        driverbottom.Repository
	sink        errorsink.ErrorSink
	mode        int
	currentStep string
	drivers     map[string]any
	runtime     map[driverbottom.Holder]any
	stepNames   []string
	symbols     map[string]map[string]*SymbolProvenance
}

func (s *Storage) Bind(v driverbottom.Holder, value any) {
	if v == nil || v.VarName() == nil {
		panic("need a var")
	}
	// log.Printf("binding var %s in mode %d, step %s\n", v.VarName().Id(), s.mode, s.currentStep)
	// So I think the steps here are:
	// 1. For the var, figure out the associated provenance
	proveni := s.symbols[s.currentStep]
	curr := proveni[v.VarName().Id()]
	if curr == nil {
		// we need to back up the list and find it BEFORE here
		for k := s.stepIndex(); curr == nil && k >= 0; k-- {
			proveni = s.symbols[s.stepNames[k]]
			curr = proveni[v.VarName().Id()]
		}
		if curr == nil {
			log.Fatalf("could not find symbol %s defined before step %s (%d)\n", v.VarName().Id(), s.currentStep, s.stepIndex())
		}
	}
	if curr.values[s.mode] == nil {
		curr.values[s.mode] = make(map[string]any)
	}
	curr.values[s.mode][s.currentStep] = value
	// 2. That should have some version that exists for this step in resolve or something
	// 3. The value (in %p form) should not be anywhere in our memory, because that would be an update
	// 4. Keep track of it ...
	s.runtime[v] = value
}

func (s *Storage) stepIndex() int {
	for k, step := range s.stepNames {
		if step == s.currentStep {
			return k
		}
	}
	panic("did not find current step")
}

func (s *Storage) Get(v driverbottom.Holder) any {
	// TODO: here we are looking for the "correct" version of the VAR, which must be the one that was in operation at the time in this mode
	// I think that may be quite complicated.
	// We also need to think about what it means to "find" and "desire" two different things.
	return s.runtime[v]
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

func (s *Storage) CurrentMode() int {
	return s.mode
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
	if s.mode == 0 {
		s.stepNames = append(s.stepNames, s.currentStep)
		s.symbols[s.currentStep] = make(map[string]*SymbolProvenance)
	}
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
	for _, k := range s.stepNames {
		iw.Intro("Step %s:\n", k)
		iw.Indent()
		syms := slices.Collect(maps.Keys(s.symbols[k]))
		slices.Sort(syms)
		for _, y := range syms {
			p := s.symbols[k][y]
			iw.IndPrintf("%s:\n", y)
			iw.Indent()
			for _, v := range p.values {
				for _, sn := range s.stepNames {
					ps := v[sn]
					if ps != nil {
						iw.IndPrintf("@%s\n", sn)
						iw.Indent()
						switch v := ps.(type) {
						case driverbottom.Describable:
							v.DumpTo(iw)
						case string:
							iw.IndPrintf("%s\n", v)
						case int:
							iw.IndPrintf("%d\n", v)
						default:
							iw.IndPrintf("%T %v\n", v, v)
						}
						iw.UnIndent()
					}
				}
			}
			iw.UnIndent()
		}
		iw.UnIndent()
	}
}

func (s *Storage) NewObjId(loc *errorsink.Location) driverbottom.Holder {
	l := len(s.symbols[s.currentStep])
	kk := fmt.Sprintf("*%s-%d", s.currentStep, l)
	id := lexicator.NewIdentifierToken(loc.Line, loc.Offset, kk)
	s.symbols[s.currentStep][kk] = NewProvenance(id)
	return interpreters.NewVarHolder(id)
}

func NewRuntimeStorage(registry driverbottom.Recall, repo driverbottom.Repository, sink errorsink.ErrorSink) driverbottom.RuntimeStorage {
	ret := &Storage{sink: sink, registry: registry, repo: repo, drivers: make(map[string]any), runtime: make(map[driverbottom.Holder]any), symbols: make(map[string]map[string]*SymbolProvenance)}
	return ret
}
