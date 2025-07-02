package runtime

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"

	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/utils"
)

type SymbolProvenance struct {
	name   driverbottom.Identifier
	values map[int]map[string]any
}

func NewProvenance(name driverbottom.Identifier) *SymbolProvenance {
	return &SymbolProvenance{name: name, values: make(map[int]map[string]any)}
}

type Storage struct {
	registry    driverbottom.Recall
	repo        driverbottom.Repository
	sink        errorsink.ErrorSink
	mode        int
	currentStep string
	drivers     map[string]any
	// TODO: I think I want to replace this with some version from "symbols", but I'm not entirely sure I know how yet ...
	// runtime   map[driverbottom.Holder]any
	stepNames []string
	symbols   map[string]map[driverbottom.Holder]*SymbolProvenance

	// Track the values we bind to be sure it doesn't get bound more than once
	unique map[driverbottom.Describable]any
}

func (s *Storage) Bind(v driverbottom.Holder, value any) {
	if v == nil || v.VarName() == nil {
		panic("need a var")
	}
	desc, ok := value.(driverbottom.Describable)
	if ok {
		if s.unique[desc] == true {
			panic("duplicate")
		}
		s.unique[desc] = true
	}

	log.Printf("binding var %s to %v in mode %d, step %s\n", v.VarName().Id(), value, s.mode, s.currentStep)
	// So I think the steps here are:
	// 1. For the var, figure out the associated provenance
	curr := s.findCoin(v)
	if curr.values[s.mode] == nil {
		curr.values[s.mode] = make(map[string]any)
	}
	curr.values[s.mode][s.currentStep] = value
	// 2. That should have some version that exists for this step in resolve or something
	// 3. The value (in %p form) should not be anywhere in our memory, because that would be an update
	// 4. Keep track of it ...
	// s.runtime[v] = value
}

func (s *Storage) IgnoreDuplicate(value any) {
	desc, ok := value.(driverbottom.Describable)
	if ok {
		s.unique[desc] = false
	}
}

func (s *Storage) findCoin(v driverbottom.Holder) *SymbolProvenance {
	proveni := s.symbols[s.currentStep]
	curr := proveni[v]
	if curr == nil {
		// we need to back up the list and find it BEFORE here
		for k := s.stepIndex(); curr == nil && k >= 0; k-- {
			proveni = s.symbols[s.stepNames[k]]
			curr = proveni[v]
		}
		if curr == nil {
			log.Printf("could not find symbol %s defined before step %s (%d)\n", v.VarName().Id(), s.currentStep, s.stepIndex())
			panic("symbol not found")
		}
	}
	return curr
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
	var curr *SymbolProvenance
	// we need to back up the list and find it BEFORE here
	for k := s.stepIndex(); curr == nil && k >= 0; k-- {
		proveni := s.symbols[s.stepNames[k]]
		curr = proveni[v]
	}
	if curr != nil {
		for mode := s.mode; mode >= 0; mode-- {
			values := curr.values[mode]
			for k := s.stepIndex(); k >= 0; k-- {
				val := values[s.stepNames[k]]
				if val != nil {
					return val
				}
			}
		}
	}
	// if curr == nil {
	// 	log.Printf("could not find symbol %s defined before step %s (%d)\n", v.VarName().Id(), s.currentStep, s.stepIndex())
	// 	panic("symbol not found")
	// }
	// }
	// return curr

	// I would have thought this was serious, but apparently copy_files needs to cope with "no bucket" in the initial phase
	log.Printf("no value has been set for %v in mode %d", v, s.CurrentMode())
	iw := utils.NewIndentWriter(os.Stdout)
	s.ExportSymbolsTo(iw)
	if s.CurrentMode() == 2 || s.CurrentMode() == 3 {
		panic("and this cannot be right in this mode")
	}
	return nil
}

func (s *Storage) GetCoin(coin driverbottom.Holder, mode int) any {
	// recent := false
	if mode == driverbottom.CURRENT_MODE {
		mode = s.CurrentMode()
		// recent = true
	}
	prov := s.findCoin(coin)
	if prov == nil {
		log.Fatalf("no coin %s\n", coin.VarName().Id())
	}
	// if recent {
	// 	log.Printf("found provenance %v\n", prov)
	// }
	val := prov.values[mode]
	if val == nil {
		// log.Printf("no coin found for: %s in mode %d; returning nil\n", coin.VarName().Id(), mode)
		return nil
	}
	for k := s.stepIndex(); k >= 0; k-- {
		if val[s.stepNames[k]] != nil {
			return val[s.stepNames[k]]
		}
	}
	panic("no val found")
}

func (s *Storage) GetCoinFrom(coin driverbottom.Holder, modes []int) any {
	prov := s.findCoin(coin)
	if prov == nil {
		log.Fatalf("no coin %s\n", coin.VarName().Id())
	}
	// log.Printf("found provenance %v\n", prov)
	for _, mode := range modes {
		val := prov.values[mode]
		if val == nil {
			continue
		}
		for k := s.stepIndex(); k >= 0; k-- {
			if val[s.stepNames[k]] != nil {
				return val[s.stepNames[k]]
			}
		}
	}
	panic("no val found")
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
	s.ExportSymbolsTo(utils.NewIndentWriter(w))
}

func (s *Storage) SetStepName(stepName string) {
	// log.Printf("mode %d: set step name to %s\n", s.mode, s.currentStep)
	s.currentStep = stepName
	if s.mode == 0 {
		s.stepNames = append(s.stepNames, s.currentStep)
		s.symbols[s.currentStep] = make(map[driverbottom.Holder]*SymbolProvenance)
	}
}

func (s *Storage) EnableSymbol(to driverbottom.Holder) {
	if s.mode != 0 {
		panic("can only enable symbols during resolution")
	}
	if s.symbols[s.currentStep][to] != nil {
		panic("cannot define symbol more than once: " + to.VarName().Id())
	}
	// log.Printf("adding %s to %s\n", to.Id(), s.currentStep)
	s.symbols[s.currentStep][to] = NewProvenance(to.VarName())
}

func (s *Storage) ExportSymbolsTo(iw driverbottom.IndentWriter) {
	for _, k := range s.stepNames {
		iw.Intro("Step %s:\n", k)
		iw.Indent()
		var syms []string
		m := make(map[string]driverbottom.Holder)
		for h := range s.symbols[k] {
			s := h.VarName().Id()
			syms = append(syms, s)
			m[s] = h
		}
		slices.Sort(syms)
		for _, y := range syms {
			p := s.symbols[k][m[y]]
			iw.IndPrintf("%s:\n", y)
			iw.Indent()
			modeValues := p.values[s.mode]
			for _, sn := range s.stepNames {
				ps := modeValues[sn]
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
			iw.UnIndent()
		}
		iw.UnIndent()
	}
}

func (s *Storage) NewObjId(loc *errorsink.Location) driverbottom.Holder {
	l := len(s.symbols[s.currentStep])
	ret := &ObjId{loc: loc, stepName: s.currentStep, which: l}
	s.symbols[s.currentStep][ret] = NewProvenance(ret.VarName())
	return ret
}

func NewRuntimeStorage(registry driverbottom.Recall, repo driverbottom.Repository, sink errorsink.ErrorSink) driverbottom.RuntimeStorage {
	ret := &Storage{sink: sink, registry: registry, repo: repo, drivers: make(map[string]any) /*runtime: make(map[driverbottom.Holder]any), */, symbols: make(map[string]map[driverbottom.Holder]*SymbolProvenance), unique: make(map[driverbottom.Describable]any)}
	return ret
}

type ObjId struct {
	loc      *errorsink.Location
	stepName string
	which    int
}

func (o *ObjId) Loc() *errorsink.Location {
	panic("unimplemented")
}

func (o *ObjId) ShortDescription() string {
	panic("unimplemented")
}

func (o *ObjId) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (o *ObjId) VarName() driverbottom.Identifier {
	kk := fmt.Sprintf("*%s-%d", o.stepName, o.which)
	return lexicator.NewIdentifierToken(o.loc.Line, o.loc.Offset, kk)
}

var _ driverbottom.Holder = &ObjId{}
