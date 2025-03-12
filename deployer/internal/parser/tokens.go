package parser

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/deployer"
)

type BaseToken struct {
	Loc deployer.Location
}

type IdentifierToken struct {
	BaseToken
	Id string
}

func (tok *BaseToken) String() string {
	return tok.Loc.String()
}

func (tok *IdentifierToken) String() string {
	return fmt.Sprintf("%s %s", tok.BaseToken.String(), tok.Id)
}

func NewIdentifierToken(file string, line, offset int, text string) Token {
	return &IdentifierToken{BaseToken: BaseToken{Loc: deployer.NewLocation(file, line, offset)}, Id: text}
}
