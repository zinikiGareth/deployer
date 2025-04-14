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
	registry     *registry.Registry
	repo         pluggable.Repository
	sink         errors.ErrorSink
	userErrorsTo io.StringWriter
	srcdir       string
	input        []string
}

func (d *DeployerImpl) ObtainRegister() pluggable.Register {
	return d.registry
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
		d.repo.ReadingFile(f)
		parser.Parse(d.registry, d.repo, d.sink, f, from)
	}
	if d.sink.HasErrors() {
		return fmt.Errorf("errors during parsing")
	}
	d.repo.ResolveAll(d.sink, d.registry)
	if d.sink.HasErrors() {
		return fmt.Errorf("errors during resolving")
	}
	targets, err := d.findTargets(targetNames...)
	if err != nil {
		return err
	}
	storage := runtime.NewRuntimeStorage(d.sink)

	storage.SetMode(pluggable.DRYRUN_MODE)
	for _, t := range targets {
		t.Execute(storage)
	}

	storage.SetMode(pluggable.EXECUTE_MODE)
	for _, t := range targets {
		t.Execute(storage)
	}

	return nil
}

// Mainly support for the test harness, but do with them as you will (not used in $cmd$)
func (d *DeployerImpl) AddSymbolListener(lsnr pluggable.SymbolListener) {
	d.repo.AddSymbolListener(lsnr)
}

func (d *DeployerImpl) Traverse(lsnr pluggable.RepositoryTraverser) {
	d.repo.Traverse(lsnr)
}

func (d *DeployerImpl) findTargets(names ...string) ([]pluggable.Target, error) {
	var targets []pluggable.Target
	var ue error
	for _, n := range names {
		t := d.repo.FindTarget(pluggable.SymbolName(n))
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
	return &DeployerImpl{registry: registry.NewRegistry(), repo: repo.NewRepository(), sink: sink, userErrorsTo: userErrorsTo}
}
