package driverbottom

import (
	"fmt"
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
	F64() float64
}

type String interface {
	Token
	Expr
	fmt.Stringer
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
	MakeAssign(holder Describable, assignTo Identifier, action any) any
	Attach(item any)
}

type VerbCommand interface {
	Handle(attacher AttachResult, tokens []Token) Interpreter
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

type List interface {
	Expr
	IsEmpty() bool
	Length() int
	Members() []Expr
}

type ExprParser interface {
	Parse(tokens []Token) (Expr, bool)
	ParseMultiple(tokens []Token) ([]Expr, bool)
}
