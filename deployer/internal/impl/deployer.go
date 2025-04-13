package impl

import (
	"fmt"
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
	registry *registry.Registry
	repo     pluggable.Repository
	sink     errors.ErrorSink
	srcdir   string
	input    []string
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
	storage := runtime.NewRuntimeStorage()

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
	for _, n := range names {
		t := d.repo.FindTarget(pluggable.SymbolName(n))
		if t == nil {
			return nil, fmt.Errorf("there is no target %s", n)
		}
		targets = append(targets, t)
	}
	return targets, nil
}

func NewDeployer(sink errors.ErrorSink) deployer.Deployer {
	return &DeployerImpl{registry: registry.NewRegistry(), repo: repo.NewRepository(), sink: sink}
}
