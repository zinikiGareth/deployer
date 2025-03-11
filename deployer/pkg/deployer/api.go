package deployer

import (
	"fmt"
	"io"

	"ziniki.org/deployer/deployer/pkg/utils"
)

type Deployer struct {
	input []string
}

type TestRunner interface {
	ErrorHandlerFor(purpose string) ErrorHandler
}

// TODO: the deployer cmd will want a version of this that writes to stdout
type ErrorHandler interface {
	io.Writer
	WriteMsg(msg string)
	Writef(fmt string, args ...any)
	Close()
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
