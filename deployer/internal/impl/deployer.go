package impl

import (
	"fmt"
	"io"
	"path/filepath"

	"ziniki.org/deployer/deployer/internal/parser"
	"ziniki.org/deployer/deployer/internal/registry"
	"ziniki.org/deployer/deployer/internal/repo"
	"ziniki.org/deployer/deployer/internal/runtime"
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/utils"
)

type DeployerImpl struct {
	tools        *pluggable.Tools
	userErrorsTo io.StringWriter
	srcdir       string
	input        []string
}

func (d *DeployerImpl) ObtainTools() *pluggable.Tools {
	return d.tools
}

func (d *DeployerImpl) ReadScriptsFrom(indir string) error {
	d.srcdir = indir
	input, err := utils.FindFiles(indir, ".dply")
	if err != nil {
		return err
	}
	d.input = append(d.input, input...)
	return nil
}

func (d *DeployerImpl) Deploy(targetNames ...string) error {
	for _, f := range d.input {
		fmt.Printf("  %s\n", f)
		from := filepath.Join(d.srcdir, f)
		d.tools.Repository.ReadingFile(f)
		parser.Parse(d.tools, f, from)
	}
	if d.tools.Reporter.HasErrors() {
		return fmt.Errorf("errors during parsing")
	}
	d.tools.Repository.ResolveAll(d.tools)
	if d.tools.Reporter.HasErrors() {
		return fmt.Errorf("errors during resolving")
	}
	targets, err := d.findTargets(targetNames...)
	if err != nil {
		return err
	}

	d.tools.Storage.SetMode(pluggable.PREPARE_MODE)
	for _, t := range targets {
		t.Prepare(d.tools.Storage)
	}

	d.tools.Storage.SetMode(pluggable.EXECUTE_MODE)
	for _, t := range targets {
		t.Execute(d.tools.Storage)
	}

	return nil
}

// Mainly support for the test harness, but do with them as you will (not used in $cmd$)
func (d *DeployerImpl) AddSymbolListener(lsnr pluggable.SymbolListener) {
	d.tools.Repository.AddSymbolListener(lsnr)
}

func (d *DeployerImpl) Traverse(lsnr pluggable.RepositoryTraverser) {
	d.tools.Repository.Traverse(lsnr)
}

func (d *DeployerImpl) findTargets(names ...string) ([]pluggable.TargetThing, error) {
	var targets []pluggable.TargetThing
	var ue error
	for _, n := range names {
		t := d.tools.Repository.FindTarget(pluggable.SymbolName(n))
		if t == nil {
			msg := fmt.Sprintf("there is no target %s\n", n)
			d.userErrorsTo.WriteString(msg)
			if ue == nil {
				ue = deployer.UserError(msg)
			}
		}
		targets = append(targets, t)
	}
	if ue != nil {
		return nil, ue
	}
	return targets, nil
}

func NewDeployer(sink errors.ErrorSink, userErrorsTo io.StringWriter) deployer.Deployer {
	reg := registry.NewRegistry()
	reporter := errors.NewErrorReporter(sink)
	storage := runtime.NewRuntimeStorage(reg, sink)
	repo := repo.NewRepository()
	tools := pluggable.NewTools(reporter, reg, reg, repo, storage)
	return &DeployerImpl{tools: tools, userErrorsTo: userErrorsTo}
}
