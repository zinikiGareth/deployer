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

type NumberToken struct {
	BaseToken
	value float64
}

type OperatorToken struct {
	BaseToken
	op string
}

type StringToken struct {
	BaseToken
	text string
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

func (tok *NumberToken) Value() float64 {
	return tok.value
}

func (tok *NumberToken) String() string {
	return fmt.Sprintf("%s %v", tok.BaseToken.String(), tok.value)
}

func (tok *OperatorToken) Is(op string) bool {
	return tok.op == op
}

func (tok *OperatorToken) Op() string {
	return tok.op
}

func (tok *OperatorToken) String() string {
	return fmt.Sprintf("%s %s", tok.BaseToken.String(), tok.op)
}

func NewIdentifierToken(file string, line, offset int, text string) pluggable.Identifier {
	return &IdentifierToken{BaseToken: BaseToken{loc: pluggable.NewLocation(file, line, offset)}, id: text}
}

func NewNumberToken(file string, line, offset int, value float64) pluggable.Number {
	return &NumberToken{BaseToken: BaseToken{loc: pluggable.NewLocation(file, line, offset)}, value: value}
}

func NewOperatorToken(file string, line, offset int, text string) pluggable.Operator {
	return &OperatorToken{BaseToken: BaseToken{loc: pluggable.NewLocation(file, line, offset)}, op: text}
}

func (tok *StringToken) Text() string {
	return tok.text
}

func (tok *StringToken) String() string {
	return fmt.Sprintf("%s %s", tok.BaseToken.String(), tok.text)
}

func NewStringToken(file string, line, offset int, text string) pluggable.String {
	return &StringToken{BaseToken: BaseToken{loc: pluggable.NewLocation(file, line, offset)}, text: text}
}
