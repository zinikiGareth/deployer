package parser

import "fmt"

type BaseToken struct {
	LineNo int
	Offset int
}

type IdentifierToken struct {
	BaseToken
	Id string
}

func (tok *BaseToken) String() string {
	return fmt.Sprintf("%d.%d", tok.LineNo, tok.Offset)
}

func (tok *IdentifierToken) String() string {
	return fmt.Sprintf("%s %s", tok.BaseToken.String(), tok.Id)
}

func NewIdentifierToken(line, offset int, text string) Token {
	return &IdentifierToken{BaseToken: BaseToken{LineNo: line, Offset: offset}, Id: text}
}
