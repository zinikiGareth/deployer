package interpreters

import (
	"ziniki.org/deployer/driver/internal/parser/exprs"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type collectListInterpreter struct {
	tools  *driverbottom.CoreTools
	parent driverbottom.PropertyParent
	prop   driverbottom.Identifier
	exprs  []driverbottom.Expr
}

func (cli *collectListInterpreter) HaveTokens(tokens []driverbottom.Token) driverbottom.Interpreter {
	expr, ok := cli.tools.Parser.Parse(tokens)
	if !ok {
		return NewIgnoreInnerScope()
	}
	cli.exprs = append(cli.exprs, expr)

	return NewDisallowInnerScope(cli.tools)
}

func (cli *collectListInterpreter) Completed() {
	cli.parent.AddProperty(cli.prop, exprs.NewListExpr(cli.exprs))
	cli.parent.Completed()
}
