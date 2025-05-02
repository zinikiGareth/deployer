package exprs

import "ziniki.org/deployer/deployer/pkg/pluggable"

type Apply struct {
	pluggable.Locatable
	Func pluggable.Function
	Args []pluggable.Expr
}
