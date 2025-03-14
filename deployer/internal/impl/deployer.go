package impl

import (
	"fmt"
	"path/filepath"

	"ziniki.org/deployer/deployer/internal/parser"
	"ziniki.org/deployer/deployer/internal/registry"
	"ziniki.org/deployer/deployer/internal/repo"
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/utils"
)

type DeployerImpl struct {
	registry *registry.Registry
	repo     pluggable.Repository
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

func (d *DeployerImpl) Deploy() error {
	for _, f := range d.input {
		fmt.Printf("  %s\n", f)
		from := filepath.Join(d.srcdir, f)
		d.repo.ReadingFile(f)
		parser.Parse(d.registry, d.repo, f, from)
	}
	return nil
}

// Mainly support for the test harness, but do with them as you will (not used in $cmd$)
func (d *DeployerImpl) AddSymbolListener(lsnr pluggable.SymbolListener) {
	d.repo.AddSymbolListener(lsnr)
}

func NewDeployer() deployer.Deployer {
	return &DeployerImpl{registry: registry.NewRegistry(), repo: repo.NewRepository()}
}
