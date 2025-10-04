package driverbottom

import (
	"fmt"
)

type Fixity string

const (
	OP_PREFIX  = Fixity("prefix")
	OP_INFIX   = Fixity("infix")
	OP_POSTFIX = Fixity("postfix")
)

type ProvideLine interface {
	BeginFile(file string)
	HaveLine(lineNo int, text string)
	EndFile()
}

type Interpreter interface {
	HaveTokens(scope Scope, tokens []Token) Interpreter
	Completed()
}

type CreateInterpreter func(tools *CoreTools, scope Scope, parent PropertyParent, prop Identifier, tokens []Token) Interpreter

type PropertyParent interface {
	AddProperty(name Identifier, expr Expr)
	AddAdverb(adverb Adverb, args []Token) Interpreter
	Completed()
}

type ValueParent interface {
	Add(expr Expr)
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
	MakeAssign(holder Holder, assignTo Identifier, action any) any
	Attach(item any) error
}

type VerbCommand interface {
	Handle(attacher AttachResult, scope Scope, tokens []Token) Interpreter
}

type Function interface {
	ReduceExpr(me Token, before []Expr, after []Expr) (Expr, bool)
	Precedence() int
	Associativity() bool
	Fixity() Fixity
}

type Expr interface {
	fmt.Stringer
	Describable
	Resolve(r Resolver) BindingRequirement
	Eval(s RuntimeStorage) any
}

type List interface {
	Expr
	IsEmpty() bool
	Length() int
	Members() []Expr
}

type MapEntry interface {
	Describable
	Key() Identifier
	Value() Expr
}

type Map interface {
	Expr
	IsEmpty() bool
	Size() int
	Members() []MapEntry
}

type ExprParser interface {
	Parse(scope Scope, tokens []Token) (Expr, bool)
	ParseMultiple(scope Scope, tokens []Token) ([]Expr, bool)
}
