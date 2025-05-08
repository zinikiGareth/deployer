package pluggable

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errors"
)

type ProvideLine interface {
	BeginFile(file string)
	HaveLine(lineNo int, text string)
	EndFile()
}

type Interpreter interface {
	HaveTokens(tokens []Token) Interpreter
	Completed()
}

type Scoper interface {
	FindTopCommand(v Identifier) TopCommand
}

type Token interface {
	Locatable
	fmt.Stringer
}

type Identifier interface {
	Token
	Id() string
}

type Number interface {
	Token
	Expr
	Value() float64
}

type String interface {
	Token
	Expr
	Text() string
}

type Operator interface {
	Token
	Is(op string) bool
	Op() string
}

type Punc interface {
	Token
	Is(punc rune) bool
	Which() rune
}

type TopCommand interface {
	Handle(tokens []Token) Interpreter
}

type TargetCommand interface {
	Handle(parent ContainingContext, tokens []Token) Interpreter
}

// Replace this with a notion of minting, blanks, dies ... I think this would be a blank
type Blank interface {
	ShortDescription() string
	Mint(tools *Tools, loc *errors.Location, named string) any
}

type Function interface {
	Eval(me Token, before []Expr, after []Expr) Expr
}

type Expr interface {
	fmt.Stringer
	Locatable
	Eval(s RuntimeStorage) any
}

type ExprParser interface {
	Parse(tokens []Token) (Expr, bool)
	ParseMultiple(tokens []Token) ([]Expr, bool)
}
