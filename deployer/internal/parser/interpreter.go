package parser

import (
	"fmt"

	"ziniki.org/deployer/deployer/internal/repo"
)

type ScopeInterpreter struct {
	repo   repo.Repository
	scoper Scoper
}

func (si *ScopeInterpreter) HaveTokens(tokens []Token) {
	for i, t := range tokens {
		fmt.Printf("token: %d %s\n", i, t)
	}

	// There are probably a "number" of cases here, but the two I am aware of are:
	// <verb> <arg>...
	// <var> "<-" <verb> <arg> ...  ||  <var> "<-" <expr>
	// And it should be fairly easy to tell this with little more than a verb-scoping thing ...

	verb, ok := tokens[0].(*IdentifierToken)
	if !ok {
		panic("first token must be an identifier")
	}
	action := si.scoper.FindVerb(verb)
	if action == nil {
		panic("this is obvs an error, but I don't have an error handler")
	}
	action.Handle(si.repo, tokens) // Will need other things as well as time goes on ...
}

func NewInterpreter(repo repo.Repository, s Scoper) Interpreter {
	return &ScopeInterpreter{repo: repo, scoper: s}
}
