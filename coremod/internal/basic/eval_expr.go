package basic

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
)

type EvalModel struct {
	driverbottom.Locatable
	tools *corebottom.Tools
	expr  driverbottom.Expr
}

func (e *EvalModel) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	e.expr.Resolve(r)
	return driverbottom.MUST_BE_BOUND
}

func (e *EvalModel) DetermineInitialState(pres corebottom.ValuePresenter) {
	pres.Present(e.expr.Eval(e.tools.Storage))
}

func (e *EvalModel) DetermineDesiredState(pres corebottom.ValuePresenter) {
	val := e.expr.Eval(e.tools.Storage)
	e.tools.Storage.IgnoreDuplicate(val)
	pres.Present(val)
}

func (e *EvalModel) DumpTo(to driverbottom.IndentWriter) {
	to.Intro("EvalAction")
	to.AttrsWhere(e)
	to.EndAttrs()
}

func (e *EvalModel) ShortDescription() string {
	panic("unimplemented")
}

type EvalCommandHandler struct {
	Tools *corebottom.Tools
}

func (ech *EvalCommandHandler) Handle(attacher driverbottom.AttachResult, scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	expr, ok := ech.Tools.Parser.Parse(scope, tokens)
	if !ok {
		panic("eval failed")
	}
	ea := &EvalModel{Locatable: expr, tools: ech.Tools, expr: expr}
	attacher.Attach(ea)

	return drivertop.NewDisallowInnerScope(ech.Tools.CoreTools)

}

func NewEvalCommandHandler(tools *corebottom.Tools) driverbottom.VerbCommand {
	return &EvalCommandHandler{Tools: tools}
}

var _ driverbottom.Describable = &EvalModel{}

var _ corebottom.ModelBuilder = &EvalModel{}
