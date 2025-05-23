package creator

import (
	"io"

	"ziniki.org/deployer/deployer/internal/impl"
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/errors"
)

func NewDeployer(sink errors.ErrorSink, userErrorsTo io.StringWriter) deployer.Deployer {
	return impl.NewDeployer(sink, userErrorsTo)
}
