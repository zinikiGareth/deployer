package exprs

import "ziniki.org/deployer/deployer/pkg/pluggable"

type Apply struct {
	pluggable.Locatable
	Func pluggable.Function
	Args []pluggable.Expr
}

func (a Apply) Eval(s pluggable.RuntimeStorage) any {
	panic("not implemented")
}

func (a Apply) String() string {
	return "APPPLLY"
}
