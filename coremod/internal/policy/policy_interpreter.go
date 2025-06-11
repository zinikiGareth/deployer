package policy

import "ziniki.org/deployer/deployer/pkg/pluggable"

type policyNestedInterpreter struct {
	tools  *pluggable.Tools
	action *PolicyAction
}

func (pni *policyNestedInterpreter) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	if len(tokens) >= 1 {
		tok, ok := tokens[0].(pluggable.Identifier)
		if !ok {
			panic("not an identifier")
		}
		switch tok.Id() {
		case "allow":
			ah := NewPolicyAllowCommandHandler(pni.tools)
			return ah.Handle(pni.action, tokens)
		default:
			panic("unimplemented")
		}
	} else {
		panic("unimplemented")
	}
}

func (pni *policyNestedInterpreter) Completed() {
	// I don't *think* we need to do anything here right now
}

func NewPolicyNestedIntepreter(tools *pluggable.Tools, pa *PolicyAction) pluggable.Interpreter {
	return &policyNestedInterpreter{tools: tools, action: pa}
}
