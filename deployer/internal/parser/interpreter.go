package parser

import (
	"ziniki.org/deployer/deployer/pkg/errors"
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

func (si *ScopeInterpreter) HaveTokens(reporter errors.ErrorRepI, tokens []pluggable.Token) pluggable.Interpreter {
	// There are probably a "number" of cases here, but the two I am aware of are:
	// <verb> <arg>...
	// <var> "<-" <verb> <arg> ...  ||  <var> "<-" <expr>
	// And it should be fairly easy to tell this with little more than a verb-scoping thing ...

	verb, ok := tokens[0].(pluggable.Identifier)
	if !ok {
		reporter.Report(0, "first token must be an identifier")
	}
	action := si.scoper.FindVerb(verb)
	if action == nil {
		reporter.Reportf(0, "there is no error handler for %s", verb)
	}
	return action.Handle(reporter, si.repo, &mayNotAddToParentOfTop{}, tokens) // Will need other things as well as time goes on ...
}

func (b *ScopeInterpreter) Completed(reporter errors.ErrorRepI) {
}

func NewInterpreter(repo pluggable.Repository, s pluggable.Scoper) pluggable.Interpreter {
	return &ScopeInterpreter{repo: repo, scoper: s}
}
