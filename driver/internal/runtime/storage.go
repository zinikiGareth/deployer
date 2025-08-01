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
	stepNames   []string
	symbols     map[string]map[driverbottom.Holder]*SymbolProvenance

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

	// log.Printf("binding var %s to %v in mode %d, step %s\n", v.VarName().Id(), value, s.mode, s.currentStep)
	curr := s.findCoin(v)
	if curr.values[s.mode] == nil {
		curr.values[s.mode] = make(map[string]any)
	}
	curr.values[s.mode][s.currentStep] = value
}

func (s *Storage) Adopt(v driverbottom.Holder, found any) {
	if v == nil || v.VarName() == nil {
		panic("need a var")
	}
	curr := s.findCoin(v)
	if curr.values[s.mode] == nil {
		curr.values[s.mode] = make(map[string]any)
	}
	for i := 0; i < s.mode; i++ {
		if curr.values[i] != nil && curr.values[i][s.currentStep] == found {
			curr.values[s.mode][s.currentStep] = found
			return
		}
	}
	log.Fatalf("cannot adopt a value for %s in mode %d which is not already a value for that var", v.VarName(), s.mode)
}

func (s *Storage) IgnoreDuplicate(value any) {
	desc, ok := value.(driverbottom.Describable)
	if ok {
		s.unique[desc] = false
	}
}

func (s *Storage) findCoin(v driverbottom.Holder) *SymbolProvenance {
	var curr *SymbolProvenance
	// we need to back up the list and find it BEFORE here
	for k := s.stepIndex(); curr == nil && k >= 0; k-- {
		proveni := s.symbols[s.stepNames[k]]
		curr = proveni[v]
	}

	if curr == nil {
		log.Printf("could not find symbol %s defined before step %s (%d)\n", v.VarName().Id(), s.currentStep, s.stepIndex())
		panic("symbol not found")
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
	curr := s.findCoin(v)
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

	if s.CurrentMode() == 2 || s.CurrentMode() == 3 {
		log.Printf("no value has been set for %v in mode %d", v, s.CurrentMode())
		iw := utils.NewIndentWriter(os.Stderr)
		s.ExportSymbolsTo(iw)
		panic("and this cannot be right in this mode")
	}
	return nil
}

func (s *Storage) GetCoin(coin driverbottom.Holder, mode int) any {
	if mode == driverbottom.CURRENT_MODE {
		mode = s.CurrentMode()
	}
	prov := s.findCoin(coin)
	if prov == nil {
		log.Fatalf("no coin %s\n", coin.VarName().Id())
	}
	val := prov.values[mode]
	if val == nil {
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
	if val == nil {
		return nil, false
	}
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
	// log.Printf("mode %d: set step name to %s\n", s.mode, stepName)
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
			for fmode := 0; fmode <= s.mode; fmode++ {
				modeValues := p.values[fmode]
				for _, sn := range s.stepNames {
					ps := modeValues[sn]
					if ps != nil {
						iw.IndPrintf("@%s[%d]\n", sn, fmode)
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

func (s *Storage) PendingObjId(loc *errorsink.Location) driverbottom.ResolvableHolder {
	return &ObjId{loc: loc, pending: true}
}

func (s *Storage) NewObjId(loc *errorsink.Location) driverbottom.ResolvableHolder {
	l := len(s.symbols[s.currentStep])
	ret := &ObjId{loc: loc, stepName: s.currentStep, which: l}
	s.symbols[s.currentStep][ret] = NewProvenance(ret.VarName())
	return ret
}

func NewRuntimeStorage(registry driverbottom.Recall, repo driverbottom.Repository, sink errorsink.ErrorSink) driverbottom.RuntimeStorage {
	ret := &Storage{sink: sink, registry: registry, repo: repo, drivers: make(map[string]any), symbols: make(map[string]map[driverbottom.Holder]*SymbolProvenance), unique: make(map[driverbottom.Describable]any)}
	return ret
}

type ObjId struct {
	loc      *errorsink.Location
	pending  bool
	stepName string
	which    int
}

func (o *ObjId) Loc() *errorsink.Location {
	return o.loc
}

func (o *ObjId) ShortDescription() string {
	return fmt.Sprintf("ObjId[%s]", o.VarName())
}

func (o *ObjId) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("ObjId")
	iw.AttrsWhere(o)
	iw.TextAttr("step", o.stepName)
	iw.TextAttr("mode", fmt.Sprintf("%d", o.which))
	iw.EndAttrs()
}

func (o *ObjId) Resolve(rs driverbottom.RuntimeStorage) {
	if o.pending {
		s := rs.(*Storage)
		o.pending = false
		o.stepName = s.currentStep
		o.which = len(s.symbols[s.currentStep])
		s.symbols[s.currentStep][o] = NewProvenance(o.VarName())
	}
}

func (o *ObjId) VarName() driverbottom.Identifier {
	if o.pending {
		panic("cannot use VarName before Var is resolved")
	}
	kk := fmt.Sprintf("*%s-%d", o.stepName, o.which)
	return lexicator.NewIdentifierToken(o.loc.Line, o.loc.Offset, kk)
}

var _ driverbottom.Holder = &ObjId{}
