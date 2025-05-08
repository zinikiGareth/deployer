package lexicator

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type BaseToken struct {
	loc *errors.Location
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

type PuncToken struct {
	BaseToken
	punc rune
}

type StringToken struct {
	BaseToken
	text string
}

func (tok *BaseToken) Loc() *errors.Location {
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

func (tok *NumberToken) Eval(s pluggable.RuntimeStorage) any {
	return tok
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

func (tok *PuncToken) Is(p rune) bool {
	return tok.punc == p
}

func (tok *PuncToken) Which() rune {
	return tok.punc
}

func (tok *PuncToken) String() string {
	return fmt.Sprintf("%s %c", tok.BaseToken.String(), tok.punc)
}

func (tok *StringToken) Text() string {
	return tok.text
}

func (tok *StringToken) String() string {
	return fmt.Sprintf("%s %s", tok.BaseToken.String(), tok.text)
}

func (tok *StringToken) Eval(s pluggable.RuntimeStorage) any {
	return tok
}

func NewIdentifierToken(line *errors.LineLoc, offset int, text string) pluggable.Identifier {
	return &IdentifierToken{BaseToken: BaseToken{loc: line.Location(offset)}, id: text}
}

func NewNumberToken(line *errors.LineLoc, offset int, value float64) pluggable.Number {
	return &NumberToken{BaseToken: BaseToken{loc: line.Location(offset)}, value: value}
}

func NewOperatorToken(line *errors.LineLoc, offset int, text string) pluggable.Operator {
	return &OperatorToken{BaseToken: BaseToken{loc: line.Location(offset)}, op: text}
}

func NewPuncToken(line *errors.LineLoc, offset int, text rune) pluggable.Punc {
	return &PuncToken{BaseToken: BaseToken{loc: line.Location(offset)}, punc: text}
}

func NewStringToken(line *errors.LineLoc, offset int, text string) pluggable.String {
	return &StringToken{BaseToken: BaseToken{loc: line.Location(offset)}, text: text}
}
