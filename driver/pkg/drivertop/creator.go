package drivertop

import (
	"io"

	"ziniki.org/deployer/driver/internal/impl"
	"ziniki.org/deployer/driver/internal/parser/exprs"
	"ziniki.org/deployer/driver/internal/parser/interpreters"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func NewDriver(sink errorsink.ErrorSink, userErrorsTo io.StringWriter) driverbottom.Driver {
	return impl.NewDriver(sink, userErrorsTo)
}

func NewIdentifierToken(loc *errorsink.Location, name string) driverbottom.Identifier {
	return lexicator.NewIdentifierToken(loc.Line, loc.Offset, name)
}

func MakeVar(name string) driverbottom.Holder {
	// id := lexicator.NewIdentifierToken(nil, 0, name)
	return &interpreters.VarHolder{}
}

func MakeString(loc *errorsink.Location, str string) driverbottom.String {
	return lexicator.NewStringToken(loc.Line, loc.Offset, str)
}

func NewListExpr(loc *errorsink.Location, es []driverbottom.Expr) driverbottom.List {
	return exprs.NewListExpr(loc, es)
}

func NewAnyExpr(loc *errorsink.Location, value any) driverbottom.Expr {
	return exprs.NewAnyExpr(loc, value)
}
