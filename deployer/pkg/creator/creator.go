package creator

import (
	"io"

	"ziniki.org/deployer/deployer/internal/impl"
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/errorsink"
)

func NewDeployer(sink errorsink.ErrorSink, userErrorsTo io.StringWriter) deployer.Deployer {
	return impl.NewDeployer(sink, userErrorsTo)
}
