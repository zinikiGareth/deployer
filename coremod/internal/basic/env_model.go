package basic

import (
	"fmt"
	"os"
	"strings"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type EnvModel struct {
	loc      *errorsink.Location
	reporter errorsink.ErrorRepI
	from     fmt.Stringer
	failed   bool
}

// Loc implements driverbottom.Describable.
func (e *EnvModel) Loc() *errorsink.Location {
	return e.loc
}

func (e *EnvModel) ShortDescription() string {
	return "EnvModel[" + e.from.String() + "]"
}

func (e *EnvModel) DumpTo(iw driverbottom.IndentWriter) {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = ""
	}
	iw.Intro("EnvModel")
	iw.AttrsWhere(e)
	str := e.from.String()
	if cwd != "" {
		str = strings.Replace(str, cwd, "<HOME>", -1)
	}
	iw.TextAttr("from", str)
	iw.EndAttrs()

}

func (e *EnvModel) String() string {
	str := e.from.String()
	val := os.Getenv(str)
	if !e.failed && val == "" {
		e.reporter.ReportAtf(e.loc, "the env var %s is not set", str)
		e.failed = true
	}
	return val
}

func NewEnvModel(r errorsink.ErrorRepI, loc *errorsink.Location, from fmt.Stringer) *EnvModel {
	return &EnvModel{reporter: r, loc: loc, from: from}
}

var _ driverbottom.Describable = &EnvModel{}
