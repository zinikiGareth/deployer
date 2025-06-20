package lexicator

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type BaseToken struct {
	loc *errorsink.Location
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

type AdverbToken struct {
	BaseToken
	name string
}

type StringToken struct {
	BaseToken
	text string
}

func (tok *BaseToken) Loc() *errorsink.Location {
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

func (tok *IdentifierToken) ShortDescription() string {
	panic("not implemented")
}

func (tok *IdentifierToken) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("Identifier")
	iw.AttrsWhere(tok)
	iw.IndPrintf("\"%s\"\n", tok.id)
	iw.EndAttrs()
}

func (tok *NumberToken) Value() float64 {
	return tok.value
}

func (tok *NumberToken) Resolve(r driverbottom.Resolver) {
}

func (tok *NumberToken) Eval(s driverbottom.RuntimeStorage) any {
	return tok.Value()
}

func (t *NumberToken) ShortDescription() string {
	panic("not implemented")
}

func (t *NumberToken) DumpTo(iw driverbottom.IndentWriter) {
	panic("not implemented")
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

func (tok *AdverbToken) Name() string {
	return tok.name
}

func (tok *AdverbToken) String() string {
	return fmt.Sprintf("%s %s", tok.BaseToken.String(), tok.name)
}

func (tok *StringToken) Text() string {
	return tok.text
}

func (tok *StringToken) ShortDescription() string {
	return fmt.Sprintf("%s %s", tok.BaseToken.String(), tok.text)
}

func (t *StringToken) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("String")
	iw.AttrsWhere(t)
	iw.IndPrintf("\"%s\"\n", t.text)
	iw.EndAttrs()
}

func (tok *StringToken) String() string {
	return tok.text
}

func (tok *StringToken) Resolve(r driverbottom.Resolver) {
}

func (tok *StringToken) Eval(s driverbottom.RuntimeStorage) any {
	// log.Printf("Eval(String) = %s\n", tok.Text())
	return tok.text
}

func NewIdentifierToken(line *errorsink.LineLoc, offset int, text string) driverbottom.Identifier {
	return &IdentifierToken{BaseToken: BaseToken{loc: line.Location(offset)}, id: text}
}

func NewNumberToken(line *errorsink.LineLoc, offset int, value float64) driverbottom.Number {
	return &NumberToken{BaseToken: BaseToken{loc: line.Location(offset)}, value: value}
}

func NewOperatorToken(line *errorsink.LineLoc, offset int, text string) driverbottom.Operator {
	return &OperatorToken{BaseToken: BaseToken{loc: line.Location(offset)}, op: text}
}

func NewPuncToken(line *errorsink.LineLoc, offset int, text rune) driverbottom.Punc {
	return &PuncToken{BaseToken: BaseToken{loc: line.Location(offset)}, punc: text}
}

func NewAdverbToken(line *errorsink.LineLoc, offset int, text string) driverbottom.Adverb {
	return &AdverbToken{BaseToken: BaseToken{loc: line.Location(offset)}, name: text}
}

func NewStringToken(line *errorsink.LineLoc, offset int, text string) driverbottom.String {
	return &StringToken{BaseToken: BaseToken{loc: line.Location(offset)}, text: text}
}
