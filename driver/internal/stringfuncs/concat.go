package stringfuncs

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type ConcatExpr struct {
	tools *driverbottom.CoreTools
	lhs   driverbottom.Expr
	rhs   driverbottom.Expr
}

func (c *ConcatExpr) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
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

func (c *ConcatExpr) Loc() *errorsink.Location {
	panic("unimplemented")
}

func (c *ConcatExpr) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	panic("unimplemented")
}

func (c *ConcatExpr) ShortDescription() string {
	panic("unimplemented")
}

func (c *ConcatExpr) String() string {
	panic("unimplemented")
}

type ConcatFunc struct {
	tools *driverbottom.CoreTools
}

func (c *ConcatFunc) Associativity() bool {
	panic("unimplemented")
}

func (c *ConcatFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_INFIX
}

func (c *ConcatFunc) Precedence() int {
	panic("unimplemented")
}

func (c *ConcatFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) (driverbottom.Expr, bool) {
	if len(before) != 1 || len(after) != 1 {
		panic("not handled")
	}
	return &ConcatExpr{tools: c.tools, lhs: before[0], rhs: after[0]}, true
}

func MakeConcatFunc(tools *driverbottom.CoreTools) driverbottom.Function {
	return &ConcatFunc{tools: tools}
}
