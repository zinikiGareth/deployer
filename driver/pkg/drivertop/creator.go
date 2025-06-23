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

func MakeVar(name string) driverbottom.Holder {
	// id := lexicator.NewIdentifierToken(nil, 0, name)
	return &interpreters.VarHolder{}
}

func MakeString(str string) driverbottom.String {
	return lexicator.NewStringToken(nil, 0, str)
}

func NewListExpr(es []driverbottom.Expr) driverbottom.List {
	return exprs.NewListExpr(es)
}
