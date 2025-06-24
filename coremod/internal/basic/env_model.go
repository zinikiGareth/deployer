package basic

import (
	"fmt"
	"os"

	"ziniki.org/deployer/driver/pkg/errorsink"
)

type EnvModel struct {
	reporter errorsink.ErrorRepI
	loc      errorsink.Location
	from     fmt.Stringer
}

func (e *EnvModel) String() string {
	str := e.from.String()
	val := os.Getenv(str)
	if val == "" {
		e.reporter.At(e.loc.Line)
		e.reporter.Reportf(e.loc.Offset, "the env var %s is not set", str)
	}
	return val
}

func NewEnvModel(r errorsink.ErrorRepI, loc errorsink.Location, from fmt.Stringer) *EnvModel {
	return &EnvModel{reporter: r, loc: loc, from: from}
}
