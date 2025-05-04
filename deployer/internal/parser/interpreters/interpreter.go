package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ScopeInterpreter struct {
	tools  *pluggable.Tools
	scoper pluggable.Scoper
}

func (si *ScopeInterpreter) HaveTokens(tokens []pluggable.Token) pluggable.Interpreter {
	// There are probably a "number" of cases here, but the two I am aware of are:
	// <verb> <arg>...
	// <var> "<-" <verb> <arg> ...  ||  <var> "<-" <expr>
	// And it should be fairly easy to tell this with little more than a verb-scoping thing ...

	verb, ok := tokens[0].(pluggable.Identifier)
	if !ok {
		si.tools.Reporter.Report(0, "first token must be an identifier")
	}
	action := si.scoper.FindTopCommand(verb)
	if action == nil {
		si.tools.Reporter.Reportf(0, "there is no error handler for %s", verb)
	}
	return action.Handle(tokens, nil) // Will need other things as well as time goes on ...
}

func (b *ScopeInterpreter) Completed() {
}

func NewInterpreter(tools *pluggable.Tools, s pluggable.Scoper) pluggable.Interpreter {
	return &ScopeInterpreter{tools: tools, scoper: s}
}
