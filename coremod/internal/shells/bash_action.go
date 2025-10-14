package shells

import (
	"log"
	"os/exec"

	"ziniki.org/deployer/coremod/internal/files"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/coremod/pkg/corepkg"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type bashAction struct {
	scope  driverbottom.Scope
	script driverbottom.Expr
}

func (b *bashAction) ShortDescription() string {
	return "bashCmd[" + b.script.ShortDescription() + "]"
}

func (b *bashAction) DumpArgs(iw driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (b *bashAction) Resolve(resolver driverbottom.Resolver) driverbottom.BindingRequirement {
	return b.script.Resolve(resolver)
}

func (b *bashAction) UpdateReality(tools *corebottom.Tools) {
	// TODO: this needs a lot of work to support arguments, environment, providing input, capturing output, etc.
	// Also needs to support @Teardown path to handle tearing down...
	dp := tools.Storage.Eval(b.script)
	if dp == nil {
		tools.Reporter.ReportAtf(b.script.Loc(), "script did not evaluate to anything")
		return
	}
	err, isErr := dp.(error)
	if isErr {
		tools.Reporter.ReportAtf(b.script.Loc(), "script evaluated to error: %v", err)
		return
	}
	dm, ok := dp.(*files.DirModel)
	if !ok {
		tools.Reporter.ReportAtf(b.script.Loc(), "script did not return a DirModel")
		return
	}
	pourer, err := dm.ObtainPourer()
	if err != nil {
		tools.Reporter.ReportAtf(b.script.Loc(), "error obtaining pourer: %v", err)
		return
	}

	// now run the script
	cmd := exec.Command(pourer.Path)
	bs, err := cmd.CombinedOutput()
	log.Printf("Output was:\n%s", string(bs))
	if err != nil {
		tools.Reporter.ReportAtf(b.script.Loc(), "error running script: %v", err)
		return
	}
	log.Printf("%s ran successfully\n", pourer.Path)
}

var _ corepkg.RealityUpdaterStrategy = &bashAction{}
