package policy

import "ziniki.org/deployer/deployer/pkg/pluggable"

type policyAllowNestedInterpreter struct {
	tools  *pluggable.Tools
	action *PolicyAllowAction
}

func (pani *policyAllowNestedInterpreter) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) >= 1 {
		tok, ok := tokens[0].(pluggable.Identifier)
		if !ok {
			panic("not an identifier")
		}
		switch tok.Id() {
		case "condition":
			ah := NewPolicyConditionCommandHandler(pani.tools)
			return ah.Handle(pani.action, tokens)
		default:
			panic("unimplemented")
		}
	} else {
		panic("unimplemented")
	}
}

func (p *policyAllowNestedInterpreter) Completed() {
}

func NewPolicyAllowNestedInterpreter(tools *pluggable.Tools, paa *PolicyAllowAction) pluggable.Interpreter {
	return &policyAllowNestedInterpreter{tools: tools, action: paa}
}
