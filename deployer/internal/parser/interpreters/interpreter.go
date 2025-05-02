package interpreters

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type mayNotAddToParentOfTop struct {
}

func (cc *mayNotAddToParentOfTop) Add(entry pluggable.Definition) {
	panic("cannot add to parent of top level; store in repo instead")
}

type ScopeInterpreter struct {
	repo   pluggable.Repository
	scoper pluggable.Scoper
}

func (si *ScopeInterpreter) HaveTokens(tools *pluggable.Tools, tokens []pluggable.Token) pluggable.Interpreter {
	// There are probably a "number" of cases here, but the two I am aware of are:
	// <verb> <arg>...
	// <var> "<-" <verb> <arg> ...  ||  <var> "<-" <expr>
	// And it should be fairly easy to tell this with little more than a verb-scoping thing ...

	verb, ok := tokens[0].(pluggable.Identifier)
	if !ok {
		tools.Reporter.Report(0, "first token must be an identifier")
	}
	action := si.scoper.FindAction(verb)
	if action == nil {
		tools.Reporter.Reportf(0, "there is no error handler for %s", verb)
	}
	return action.Handle(tools, &mayNotAddToParentOfTop{}, tokens) // Will need other things as well as time goes on ...
}

func (b *ScopeInterpreter) Completed(tools *pluggable.Tools) {
}

func NewInterpreter(repo pluggable.Repository, s pluggable.Scoper) pluggable.Interpreter {
	return &ScopeInterpreter{repo: repo, scoper: s}
}
