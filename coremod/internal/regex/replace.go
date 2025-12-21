package regex

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type StrReplace struct {
	driverbottom.Locatable
	Tools   *corebottom.Tools
	Value   driverbottom.Expr
	Pattern driverbottom.Expr
	Replace driverbottom.Expr
}

func (s *StrReplace) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (s *StrReplace) Eval(storage driverbottom.RuntimeStorage) any {
	patt, worked := storage.EvalAsStringer(s.Pattern)
	if !worked {
		s.Tools.Reporter.ReportAtf(s.Pattern.Loc(), "pattern must be a string")
		return nil
	}
	re, err := regexp.Compile(patt.String())
	if err != nil {
		s.Tools.Reporter.ReportAtf(s.Pattern.Loc(), "invalid pattern %s: %v", patt, err)
		return err
	}
	str, worked := storage.EvalAsStringer(s.Value)
	if !worked {
		s.Tools.Reporter.ReportAtf(s.Value.Loc(), "value must be a string")
		return nil
	}
	match := re.FindStringSubmatchIndex(str.String())
	if match == nil {
		s.Tools.Reporter.ReportAtf(s.Value.Loc(), "value '%s' did not match pattern '%s'", str.String(), patt.String())
		return nil
	}
	replace, worked := storage.EvalAsStringer(s.Replace)
	if !worked {
		s.Tools.Reporter.ReportAtf(s.Replace.Loc(), "replace must be a string")
		return nil
	}

	rs := []rune(replace.String())

	ret := strings.Builder{}
	for i := 0; i < len(rs); i++ {
		if rs[i] == '\\' {
			i++
			if rs[i] == '{' {
				toCap := "none"
				i++
				if rs[i] == '^' {
					i++
					toCap = "cap"
					if rs[i] == '^' {
						i++
						toCap = "upper"
					}
				}
				from := i
				for i < len(rs) && unicode.IsDigit(rs[i]) {
					i++
				}
				if i == from {
					panic("no number present")
				}
				if i >= len(rs) {
					panic("past end of string")
				}
				if rs[i] != '}' {
					panic("no closing }")
				}
				fld, err := strconv.Atoi(string(rs[from:i]))
				if err != nil {
					panic(err)
				}
				if len(match) < 2*fld+2 {
					panic(fmt.Sprintf("arg out of range: %d", fld))
				}
				have := str.String()[match[fld*2]:match[fld*2+1]]
				switch toCap {
				case "cap":
					have = strings.ToUpper(have[0:1]) + have[1:]
				case "upper":
					have = strings.ToUpper(have)
				default:

				}
				ret.WriteString(have)
			} else {
				panic("unimplemented")
			}
		} else {
			ret.WriteRune(rs[i])
		}
	}
	return ret.String()
}

func (s *StrReplace) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	br := s.Pattern.Resolve(r)
	if br != driverbottom.MAY_BE_BOUND {
		s.Tools.Reporter.ReportAtf(s.Loc(), "pattern must be a value")
	}
	br = s.Value.Resolve(r)
	if br != driverbottom.MAY_BE_BOUND {
		s.Tools.Reporter.ReportAtf(s.Loc(), "value must be a value")
	}
	br = s.Replace.Resolve(r)
	if br != driverbottom.MAY_BE_BOUND {
		s.Tools.Reporter.ReportAtf(s.Loc(), "replace must be a value")
	}
	return driverbottom.MUST_BE_BOUND
}

func (s *StrReplace) ShortDescription() string {
	return fmt.Sprintf("regex.match[%s,%s]", s.Value, s.Pattern)
}

func (s *StrReplace) String() string {
	panic("unimplemented")
}

type ReplaceFunc struct {
	tools *corebottom.Tools
}

func (m *ReplaceFunc) Associativity() bool {
	panic("unimplemented")
}

func (m *ReplaceFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_PREFIX
}

func (m *ReplaceFunc) Precedence() int {
	panic("unimplemented")
}

func (m *ReplaceFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) (driverbottom.Expr, bool) {
	rep := m.tools.Reporter
	if len(before) != 0 || len(after) != 3 {
		rep.Report(me.Loc().Offset, "regex.replace <string> <pattern> <replace-with>")
		return nil, false
	}
	str := after[0]
	patt := after[1]
	replace := after[2]
	return &StrReplace{Tools: m.tools, Locatable: me, Value: str, Pattern: patt, Replace: replace}, true
}

func MakeReplaceFunc(tools *corebottom.Tools) *ReplaceFunc {
	return &ReplaceFunc{tools: tools}
}

var _ driverbottom.Expr = &StrReplace{}

var _ driverbottom.Function = &ReplaceFunc{}
