package interpreters

import (
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type verbCommandInterpreter struct {
	tools        *driverbottom.CoreTools
	attacher     driverbottom.AttachResult
	forExtension string
	allowAssign  bool
}

func (si *verbCommandInterpreter) HaveTokens(scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	ok, toks, assignTo := si.splitOnArrow(tokens)
	if !ok { //
		return NewIgnoreInnerScope()
	}
	if len(toks) < 1 {
		si.tools.Reporter.Reportf(0, "must have a command")
		return NewIgnoreInnerScope()
	}
	verb, ok := toks[0].(driverbottom.Identifier)
	if !ok {
		si.tools.Reporter.Report(0, "first token must be an identifier")
		return NewIgnoreInnerScope()
	}
	cmd := si.tools.Recall.Find(si.forExtension, verb.Id())
	if cmd == nil && assignTo != nil {
		_, succeeded := si.tools.Parser.Parse(scope, toks)
		if succeeded {
			cmd = si.tools.Recall.Find(si.forExtension, "eval")
		}
	}
	if cmd != nil {
		action, ok := cmd.(driverbottom.VerbCommand)
		if !ok {
			si.tools.Reporter.Reportf(0, "%s is not a command", verb.Id())
			return NewIgnoreInnerScope()
		}
		a := si.attacher
		if assignTo != nil {
			a = WillAssignTo(si.tools, a, assignTo)
		}
		return action.Handle(a, scope, toks)
	} else {
		si.tools.Reporter.Reportf(0, "%s is not a command", verb.Id())
		return NewIgnoreInnerScope()
	}
}

func (b *verbCommandInterpreter) Completed() {
}

func (b *verbCommandInterpreter) splitOnArrow(tokens []driverbottom.Token) (bool, []driverbottom.Token, driverbottom.Identifier) {
	if !b.allowAssign {
		return true, tokens, nil
	}
	for i, t := range tokens {
		arrow, ok := t.(driverbottom.Operator)
		if ok && arrow.Is("=>") {
			if i+2 != len(tokens) {
				b.tools.Reporter.Reportf(arrow.Loc().Offset, "invalid use of =>")
				return false, nil, nil
			}
			id, ok := tokens[i+1].(driverbottom.Identifier)
			if !ok {
				b.tools.Reporter.Reportf(tokens[i+1].Loc().Offset, "can only assign to a variable")
				return false, nil, nil
			}
			return true, tokens[0:i], id
		}
	}
	return true, tokens, nil
}

func NewVerbCommandInterpreter(tools *driverbottom.CoreTools, attacher driverbottom.AttachResult, forExtensionPoint string, allowAssignments bool) driverbottom.Interpreter {
	return &verbCommandInterpreter{tools: tools, attacher: attacher, forExtension: forExtensionPoint, allowAssign: allowAssignments}
}
