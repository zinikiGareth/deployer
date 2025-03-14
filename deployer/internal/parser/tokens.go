package parser

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type BaseToken struct {
	loc pluggable.Location
}

type IdentifierToken struct {
	BaseToken
	id string
}

func (tok *BaseToken) Loc() pluggable.Location {
	return tok.loc
}

func (tok *BaseToken) String() string {
	return tok.loc.String()
}

func (tok *IdentifierToken) Id() string {
	return tok.id
}

func (tok *IdentifierToken) String() string {
	return fmt.Sprintf("%s %s", tok.BaseToken.String(), tok.id)
}

func NewIdentifierToken(file string, line, offset int, text string) pluggable.Identifier {
	return &IdentifierToken{BaseToken: BaseToken{loc: pluggable.NewLocation(file, line, offset)}, id: text}
}
