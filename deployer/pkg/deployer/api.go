package deployer

import (
	"fmt"
	"os"
	"strings"
)

type Deployer struct {
	input []string
}

func (d *Deployer) ReadScriptsFrom(indir string) error {
	input, err := findFiles(indir)
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

func findFiles(indir string) ([]string, error) {
	files, err := os.ReadDir(indir)
	if err != nil {
		return nil, fmt.Errorf("could not read script directory %s: %v", indir, err)
	}
	deployFiles := make([]string, 0)
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".dply") {
			deployFiles = append(deployFiles, f.Name())
		}
	}
	return deployFiles, nil
}

func NewDeployer() *Deployer {
	return &Deployer{}
}
