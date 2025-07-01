package interpreters

import (
	"ziniki.org/deployer/driver/internal/parser/exprs"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type collectListInterpreter struct {
	tools  *driverbottom.CoreTools
	parent driverbottom.PropertyParent
	prop   driverbottom.Identifier
	addTo  driverbottom.ValueParent
	exprs  []driverbottom.Expr
}

func (cli *collectListInterpreter) HaveTokens(scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	expr, ok := cli.tools.Parser.Parse(scope, tokens)
	if !ok {
		return NewIgnoreInnerScope()
	}
	switch expr := expr.(type) {
	case *exprs.ListExpr:
		if expr.IsEmpty() {
			return NewCollectListInnerScope(cli.tools, nil, nil, cli)
		}
	case *exprs.MapExpr:
		if expr.IsEmpty() {
			return NewCollectMapInnerScope(cli.tools, nil, nil, cli)
		}
	}

	cli.exprs = append(cli.exprs, expr)
	return NewDisallowInnerScope(cli.tools)
}

func (cli *collectListInterpreter) Completed() {
	cli.parent.AddProperty(cli.prop, exprs.NewListExpr(cli.exprs))
	// cli.parent.Completed()
}

func (cli *collectListInterpreter) Add(expr driverbottom.Expr) {
	cli.exprs = append(cli.exprs, expr)
}
