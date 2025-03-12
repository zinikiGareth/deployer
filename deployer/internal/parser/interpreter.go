package parser

import "fmt"

type ScopeInterpreter struct {
}

func (si *ScopeInterpreter) HaveTokens(tokens []Token) {
	for i, t := range tokens {
		fmt.Printf("token: %d %s\n", i, t)
	}
}

func NewInterpreter(s Scoper) Interpreter {
	return &ScopeInterpreter{}
}
