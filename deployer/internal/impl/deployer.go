package impl

import (
	"fmt"

	"ziniki.org/deployer/deployer/internal/repo"
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/utils"
)

type DeployerImpl struct {
	repo  repo.Repository
	input []string
}

func (d *DeployerImpl) ReadScriptsFrom(indir string) error {
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
	}
	return nil
}

// Mainly support for the test harness, but do with them as you will (not used in $cmd$)
func (d *DeployerImpl) AddSymbolListener(lsnr deployer.SymbolListener) {
	d.repo.AddSymbolListener(lsnr)
}

func NewDeployer() deployer.Deployer {
	return &DeployerImpl{repo: repo.NewRepository()}
}
