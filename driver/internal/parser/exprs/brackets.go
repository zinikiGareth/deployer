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

func (l AsList) Loc() *errorsink.Location {
	return l.Tokens[0].Loc()
}

func (l AsList) String() string {
	strs := make([]string, len(l.Tokens))
	for i := 0; i < len(l.Tokens); i++ {
		strs[i] = l.Tokens[i].String()
	}
	return strings.Join(strs, " ")
}

type AsMap struct {
	Tokens []driverbottom.Token
}

// Loc implements driverbottom.Token.
func (m AsMap) Loc() *errorsink.Location {
	return m.Tokens[0].Loc()
}

// String implements driverbottom.Token.
func (m AsMap) String() string {
	strs := make([]string, len(m.Tokens))
	for i := 0; i < len(m.Tokens); i++ {
		strs[i] = m.Tokens[i].String()
	}
	return strings.Join(strs, " ")
}

