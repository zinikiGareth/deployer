package pluggable

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errorsink"
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

type PropertyParent interface {
	AddProperty(name Identifier, expr Expr)
	AddAdverb(adverb Adverb, args []Token) Interpreter
	Completed()
}

type Scoper interface {
	FindVerbCommand(v Identifier) VerbCommand
}

type Token interface {
	Locatable
	fmt.Stringer
}

type Identifier interface {
	Token
	Describable
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

type Adverb interface {
	Token
	Name() string
}

type Punc interface {
	Token
	Is(punc rune) bool
	Which() rune
}

type AttachResult interface {
	Attach(item any)
}

type VerbCommand interface {
	Handle(attacher AttachResult, tokens []Token) Interpreter
}

// Replace this with a notion of minting, blanks, dies ... I think this would be a blank
type Blank interface {
	ShortDescription() string
	Find(tools *CoreTools, loc *errorsink.Location, named string) any
	Mint(tools *CoreTools, loc *errorsink.Location, named string, props map[Identifier]Expr, teardown TearDown) any
}

type Function interface {
	ReduceExpr(me Token, before []Expr, after []Expr) Expr
}

type Expr interface {
	fmt.Stringer
	Describable
	Resolve(r Resolver)
	Eval(s RuntimeStorage) any
}

type ExprParser interface {
	Parse(tokens []Token) (Expr, bool)
	ParseMultiple(tokens []Token) ([]Expr, bool)
}
