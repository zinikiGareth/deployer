package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/utils"
	"ziniki.org/deployer/golden/internal/errors"
	"ziniki.org/deployer/golden/internal/lsnrs"
)

func (r *TestRunner) TestDeployment(eh errors.TestErrorHandler) {
	err := r.deployer.ReadScriptsFrom(r.scripts)
	if err != nil {
		fmt.Printf("Error reading scripts from %s: %v\n", r.scripts, err)
		return
	}
	targetFile := filepath.Join(r.scripts, "targets")
	targets, err := r.ReadTargets(targetFile)
	if err != nil {
		fmt.Printf("Error reading target list from %s: %v\n", targetFile, err)
		return
	}
	err = r.deployer.Deploy(targets...)
	if err != nil {
		// this is really just repeating information
		// should it go in a file?
		eh := r.ErrorHandlerFor("log")
		eh.Writef("Error deploying: %v\n", err)
	}
	storer := lsnrs.NewGoldenRepoStorer()
	r.deployer.Traverse(storer)
	storer.DumpNamesTo(r.repoOut)
	storer.DumpDefnsTo(r.repoOut)
	r.golden.compareAll()
}

func (r *TestRunner) ReadTargets(file string) ([]string, error) {
	lines, err := utils.FileAsLines(file)

	if err != nil {
		pe, ok := err.(*os.PathError)
		if !ok {
			return nil, err
		}
		if pe.Op == "open" && pe.Err == syscall.ENOENT {
			return nil, nil
		}
		return nil, err
	}

	// TODO: I feel we will want to do some cleaning up here
	// Specifically:
	//   * remove "blank" and "comment (#)" lines
	//   * allow multiple targets on one line and break them up
	return lines, nil
}

func (r *TestRunner) WrapUp() {
	r.symbolLsnr.Close()
	r.tracker.Done()
}

func (r *TestRunner) ErrorHandlerFor(purpose string) deployer.ErrorHandler {
	return r.tracker.ErrorHandlerFor(purpose)
}
