package creator

import (
	"io"

	"ziniki.org/deployer/deployer/internal/impl"
	"ziniki.org/deployer/deployer/internal/parser/interpreters"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func NewDriver(sink errorsink.ErrorSink, userErrorsTo io.StringWriter) deployer.Driver {
	return impl.NewDriver(sink, userErrorsTo)
}

func MakeVar(name string) pluggable.Holder {
	// id := lexicator.NewIdentifierToken(nil, 0, name)
	return &interpreters.VarHolder{}
}

func MakeString(str string) pluggable.String {
	return lexicator.NewStringToken(nil, 0, str)
}
