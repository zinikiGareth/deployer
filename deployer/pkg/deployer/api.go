package deployer

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/utils"
)

type Deployer struct {
	input []string
}

func (d *Deployer) ReadScriptsFrom(indir string) error {
	input, err := utils.FindFiles(indir, ".dply")
	if err != nil {
		return err
	}
	d.input = append(d.input, input...)
	return nil
}

func (d *Deployer) Deploy() error {
	for _, f := range d.input {
		fmt.Printf("  %s\n", f)
	}
	return nil
}

func NewDeployer() *Deployer {
	return &Deployer{}
}
