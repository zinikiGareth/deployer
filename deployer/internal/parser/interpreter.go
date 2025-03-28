package parser

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type ScopeInterpreter struct {
	repo   pluggable.Repository
	scoper pluggable.Scoper
}

func (si *ScopeInterpreter) HaveTokens(reporter *errors.ErrorReporter, tokens []pluggable.Token) pluggable.Interpreter {
	// There are probably a "number" of cases here, but the two I am aware of are:
	// <verb> <arg>...
	// <var> "<-" <verb> <arg> ...  ||  <var> "<-" <expr>
	// And it should be fairly easy to tell this with little more than a verb-scoping thing ...

	verb, ok := tokens[0].(pluggable.Identifier)
	if !ok {
		panic("first token must be an identifier")
	}
	action := si.scoper.FindVerb(verb)
	if action == nil {
		panic("this is obvs an error, but I don't have an error handler")
	}
	return action.Handle(reporter, si.repo, tokens) // Will need other things as well as time goes on ...
}

func NewInterpreter(repo pluggable.Repository, s pluggable.Scoper) pluggable.Interpreter {
	return &ScopeInterpreter{repo: repo, scoper: s}
}
