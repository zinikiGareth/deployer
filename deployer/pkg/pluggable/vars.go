package pluggable

type Var interface {
	Expr
	Named() Identifier
	Binding() Describable
}

type Holder interface {
	Describable
}
