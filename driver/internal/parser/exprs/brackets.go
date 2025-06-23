package exprs

import (
	"strings"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type ParenReduction interface {
	ReduceParens(tokens []driverbottom.Token) ([]driverbottom.Token, bool)
}

type Bracketed struct {
	Tokens []driverbottom.Token
}

func (b Bracketed) Loc() *errorsink.Location {
	return b.Tokens[0].Loc()
}

func (b Bracketed) String() string {
	strs := make([]string, len(b.Tokens))
	for i := 0; i < len(b.Tokens); i++ {
		strs[i] = b.Tokens[i].String()
	}
	return strings.Join(strs, " ")
}

type AsList struct {
	Tokens []driverbottom.Token
}

func (b AsList) Loc() *errorsink.Location {
	return b.Tokens[0].Loc()
}

func (b AsList) String() string {
	strs := make([]string, len(b.Tokens))
	for i := 0; i < len(b.Tokens); i++ {
		strs[i] = b.Tokens[i].String()
	}
	return strings.Join(strs, " ")
}
