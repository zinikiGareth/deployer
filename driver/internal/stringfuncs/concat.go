package stringfuncs

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type ConcatExpr struct {
	driverbottom.Locatable
	tools *driverbottom.CoreTools
	lhs   driverbottom.Expr
	rhs   driverbottom.Expr
}

func (c *ConcatExpr) DumpTo(to driverbottom.IndentWriter) {
	to.Intro("concat")
	to.AttrsWhere(c)
	to.NestedAttr("lhs", c.lhs)
	to.NestedAttr("rhs", c.rhs)
	to.EndAttrs()
}

func (c *ConcatExpr) Eval(s driverbottom.RuntimeStorage) any {
	l, ok := c.tools.Storage.EvalAsStringer(c.lhs)
	if !ok {
		panic("error eval lhs")
	}
	r, ok := c.tools.Storage.EvalAsStringer(c.rhs)
	if !ok {
		panic("error eval rhs")
	}
	return l.String() + r.String()
}

func (c *ConcatExpr) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	br := c.lhs.Resolve(r)
	if br == driverbottom.NO_VALUE {
		c.tools.Reporter.ReportAtf(c.lhs.Loc(), "lhs of ++ must return a value")
	}
	br = c.rhs.Resolve(r)
	if br == driverbottom.NO_VALUE {
		c.tools.Reporter.ReportAtf(c.rhs.Loc(), "rhs of ++ must return a value")
	}
	return driverbottom.MAY_BE_BOUND
}

func (c *ConcatExpr) ShortDescription() string {
	return fmt.Sprintf("Concat[%s,%s]", c.lhs.ShortDescription(), c.rhs.ShortDescription())
}

func (c *ConcatExpr) String() string {
	panic("unimplemented")
}

type ConcatFunc struct {
	tools *driverbottom.CoreTools
}

func (c *ConcatFunc) Associativity() bool {
	return false
}

func (c *ConcatFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_INFIX
}

func (c *ConcatFunc) Precedence() int {
	return 5
}

func (c *ConcatFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) (driverbottom.Expr, bool) {
	if len(before) != 1 || len(after) != 1 {
		panic("not handled")
	}
	return &ConcatExpr{Locatable: me, tools: c.tools, lhs: before[0], rhs: after[0]}, true
}

func MakeConcatFunc(tools *driverbottom.CoreTools) driverbottom.Function {
	return &ConcatFunc{tools: tools}
}
