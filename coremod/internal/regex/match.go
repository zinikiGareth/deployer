package regex

import (
	"fmt"
	"regexp"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type StrMatch struct {
	driverbottom.Locatable
	Tools   *corebottom.Tools
	Value   driverbottom.Expr
	Pattern driverbottom.Expr
}

func (s *StrMatch) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (s *StrMatch) Eval(storage driverbottom.RuntimeStorage) any {
	patt, worked := storage.EvalAsStringer(s.Pattern)
	if !worked {
		s.Tools.Reporter.ReportAtf(s.Pattern.Loc(), "pattern must be a string")
		return false
	}
	re, err := regexp.Compile(patt.String())
	if err != nil {
		s.Tools.Reporter.ReportAtf(s.Pattern.Loc(), "invalid pattern %s: %v", patt, err)
		return err
	}
	str, worked := storage.EvalAsStringer(s.Value)
	if !worked {
		s.Tools.Reporter.ReportAtf(s.Value.Loc(), "value must be a string")
		return false
	}
	return re.FindStringIndex(str.String()) != nil
}

func (s *StrMatch) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	br := s.Pattern.Resolve(r)
	if br != driverbottom.MAY_BE_BOUND {
		s.Tools.Reporter.ReportAtf(s.Loc(), "pattern must be a value")
	}
	br = s.Value.Resolve(r)
	if br != driverbottom.MAY_BE_BOUND {
		s.Tools.Reporter.ReportAtf(s.Loc(), "value must be a value")
	}
	return driverbottom.MUST_BE_BOUND
}

func (s *StrMatch) ShortDescription() string {
	return fmt.Sprintf("regex.match[%s,%s]", s.Value, s.Pattern)
}

func (s *StrMatch) String() string {
	panic("unimplemented")
}

type MatchFunc struct {
	tools *corebottom.Tools
}

func (m *MatchFunc) Associativity() bool {
	panic("unimplemented")
}

func (m *MatchFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_PREFIX
}

func (m *MatchFunc) Precedence() int {
	panic("unimplemented")
}

func (m *MatchFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) (driverbottom.Expr, bool) {
	rep := m.tools.Reporter
	if len(before) != 0 || len(after) != 2 {
		rep.Report(me.Loc().Offset, "regex.match <string> <pattern>")
		return nil, false
	}
	str := after[0]
	patt := after[1]
	return &StrMatch{Tools: m.tools, Locatable: me, Value: str, Pattern: patt}, true
}

func MakeMatchFunc(tools *corebottom.Tools) *MatchFunc {
	return &MatchFunc{tools: tools}
}

var _ driverbottom.Expr = &StrMatch{}

var _ driverbottom.Function = &MatchFunc{}
