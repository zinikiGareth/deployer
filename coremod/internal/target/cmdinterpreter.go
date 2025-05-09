package target

import (
	"reflect"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type commandScope struct {
	tools     *pluggable.Tools
	container pluggable.ContainingContext
}

func (cs *commandScope) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	ok, toks, assignTo := cs.splitOnArrow(tokens)
	if !ok { //
		return interpreters.IgnoreInnerScope()
	}
	if len(toks) < 1 {
		cs.tools.Reporter.Reportf(0, "must have a command")
		return interpreters.IgnoreInnerScope()
	}

	tok0 := toks[0]
	cmd, ok := tok0.(pluggable.Identifier)
	if !ok {
		cs.tools.Reporter.Reportf(tok0.Loc().Offset, "invalid command: %s", tok0.String())
		return interpreters.IgnoreInnerScope()
	}

	action, ok := cs.tools.Recall.Find(reflect.TypeFor[pluggable.TargetCommand](), cmd.Id()).(pluggable.TargetCommand)
	if !ok {
		cs.tools.Reporter.Reportf(tok0.Loc().Offset, "unknown command: %s", cmd.Id())
		return interpreters.IgnoreInnerScope()
	}

	cc := cs.container
	if assignTo != nil {
		cc = &WithAssignTo{tools: cs.tools, container: cc, assignTo: assignTo}
	}

	innerScope := action.Handle(cc, toks)
	return innerScope
}

func (b *commandScope) Completed() {
}

func TargetCommandInterpreter(tools *pluggable.Tools, container pluggable.ContainingContext) pluggable.Interpreter {
	return &commandScope{tools: tools, container: container}
}

func (b *commandScope) splitOnArrow(tokens []pluggable.Token) (bool, []pluggable.Token, pluggable.Identifier) {
	for i, t := range tokens {
		arrow, ok := t.(pluggable.Operator)
		if ok && arrow.Is("=>") {
			if i+2 != len(tokens) {
				b.tools.Reporter.Reportf(arrow.Loc().Offset, "invalid use of =>")
				return false, nil, nil
			}
			id, ok := tokens[i+1].(pluggable.Identifier)
			if !ok {
				b.tools.Reporter.Reportf(tokens[i+1].Loc().Offset, "can only assign to a variable")
				return false, nil, nil
			}
			return true, tokens[0:i], id
		}
	}
	return true, tokens, nil
}

type WithAssignTo struct {
	tools     *pluggable.Tools
	assignTo  pluggable.Identifier
	container pluggable.ContainingContext
}

func (wat *WithAssignTo) Add(d pluggable.Action) {
	wat.container.Add(&DoAssign{tools: wat.tools, assignTo: wat.assignTo, action: d})
}

type DoAssign struct {
	tools    *pluggable.Tools
	assignTo pluggable.Identifier
	resolved pluggable.Describable
	action   pluggable.Action
}

func (d *DoAssign) Loc() *errors.Location {
	return d.assignTo.Loc()
}

func (d *DoAssign) DumpTo(w pluggable.IndentWriter) {
	w.Intro("AssignTo")
	w.AttrsWhere(d.assignTo)
	w.TextAttr("assignTo", d.assignTo.Id())
	d.action.DumpTo(w)
	w.EndAttrs()
}

func (d *DoAssign) Resolve(r pluggable.Resolver, b pluggable.Binder) {
	// d.resolved = r.MakeNew(d.assignTo)
	d.action.Resolve(r, d)
}

func (d *DoAssign) ShortDescription() string {
	return "DoAssign[" + d.assignTo.Id() + "<-" + d.action.ShortDescription() + "]"
}

func (d *DoAssign) Prepare(pres pluggable.ValuePresenter) {
	d.action.Prepare(d)
}

func (d *DoAssign) Execute() {
	d.action.Execute()
}

func (d *DoAssign) MayBind(val pluggable.Describable) {
	if d.assignTo != nil {
		d.resolved = val
		d.tools.Repository.IntroduceSymbol(pluggable.SymbolName(d.assignTo.Id()), val)
	}
}

func (d *DoAssign) MustBind(val pluggable.Describable) {
	if d.assignTo == nil {
		panic("assignTo is not specified") // should be an error
	}
	d.MayBind(val)
}

func (d *DoAssign) Present(value any) {
	if d.resolved == nil { // can't do anything if we didn't resolve it
		return
	}
	d.tools.Storage.Bind(d.resolved, value)
}
