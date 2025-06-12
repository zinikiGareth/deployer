package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type verbCommandInterpreter struct {
	tools        *pluggable.Tools
	attacher     pluggable.AttachResult
	forExtension string
	allowAssign  bool
}

func (si *verbCommandInterpreter) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	ok, toks, assignTo := si.splitOnArrow(tokens)
	if !ok { //
		return NewIgnoreInnerScope()
	}
	if len(toks) < 1 {
		si.tools.Reporter.Reportf(0, "must have a command")
		return NewIgnoreInnerScope()
	}
	verb, ok := tokens[0].(pluggable.Identifier)
	if !ok {
		si.tools.Reporter.Report(0, "first token must be an identifier")
		return NewIgnoreInnerScope()
	}
	cmd := si.tools.Recall.Find(si.forExtension, verb.Id())
	if cmd == nil {
		si.tools.Reporter.Reportf(0, "there is no command %s", verb.Id())
		return NewIgnoreInnerScope()
	}
	action, ok := cmd.(pluggable.VerbCommand)
	if !ok {
		si.tools.Reporter.Reportf(0, "%s is not a command", verb.Id())
		return NewIgnoreInnerScope()
	}

	a := si.attacher
	if assignTo != nil {
		a = &WithAssignTo{tools: si.tools, container: a, assignTo: assignTo}
	}
	return action.Handle(a, toks)
}

func (b *verbCommandInterpreter) Completed() {
}

func (b *verbCommandInterpreter) splitOnArrow(tokens []pluggable.Token) (bool, []pluggable.Token, pluggable.Identifier) {
	if !b.allowAssign {
		return true, tokens, nil
	}
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

func NewVerbCommandInterpreter(tools *pluggable.Tools, attacher pluggable.AttachResult, forExtensionPoint string, allowAssignments bool) pluggable.Interpreter {
	return &verbCommandInterpreter{tools: tools, attacher: attacher, forExtension: forExtensionPoint, allowAssign: allowAssignments}
}
