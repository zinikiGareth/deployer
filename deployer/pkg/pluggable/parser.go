package pluggable

import "fmt"

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
	Value() float64
}

type String interface {
	Token
	Text() string
}

type Operator interface {
	Token
	Is(op string) bool
	Op() string
}

type TopCommand interface {
	Handle(tokens []Token, assignTo Identifier) Interpreter
}

type TargetCommand interface {
	Handle(parent ContainingContext, tokens []Token, assignTo Identifier) Interpreter
}

type Noun interface {
	ShortDescription() string
	CreateWithName(named string) any
}

type Function interface {
	Eval(me Token, before []Expr, after []Expr) Expr
}

type Expr interface {
	fmt.Stringer
	Locatable
}

type ExprParser interface {
	Parse(tokens []Token) (Expr, bool)
}
