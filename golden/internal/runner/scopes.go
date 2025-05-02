package runner

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"ziniki.org/deployer/deployer/pkg/utils"
	"ziniki.org/deployer/golden/internal/errors"
)

func (r *TestRunner) TestScopes(eh errors.TestErrorHandler) {
	testIn := filepath.Join(r.base, "scope-test")

	// Make sure clean directory exists
	err := utils.EnsureCleanDir(testIn)
	if err != nil {
		fmt.Printf("error ensuring %s: %v\n", testIn, err)
		return
	}

	// Now copy all the files across
	nin, err := utils.CopyFilesFrom(r.scripts, testIn, ".dply")
	if err != nil {
		fmt.Printf("error copying files from %s to %s: %v\n", r.scripts, testIn, err)
		return
	}
	err = utils.EnsureDir(r.scopes)
	if err != nil {
		fmt.Printf("error ensuring %s: %v\n", r.scopes, err)
		return
	}
	nout, err := utils.CopyFilesFrom(r.scopes, testIn, ".snap")
	if err != nil {
		fmt.Printf("error copying files from %s to %s: %v\n", r.scopes, testIn, err)
		return
	}
	if nin == 0 && nout == 0 {
		// fmt.Printf("no input or output files in %s\n", r.test)
	} else if nin == nout {
		cmd := exec.Command("vscode-tmgrammar-snap", "--config", "../../deployer-vsix/package.json", testIn+"/*.dply")
		// cmd.Dir = r.root
		cmd.Stdout = eh
		cmd.Stderr = eh
		err := cmd.Run()
		if err != nil {
			eh.Writef("failed running vscode-tmgrammar-snap: %v\n", err)
			eh.Fail()
			return
		}
	} else {
		cmd := exec.Command("vscode-tmgrammar-snap", "--config", "../../deployer-vsix/package.json", "--updateSnapshot", testIn+"/*.dply")
		// cmd.Dir = r.root
		cmd.Stdout = eh
		cmd.Stderr = eh
		err := cmd.Run()
		if err != nil {
			eh.Writef("failed running vscode-tmgrammar-snap: %v\n", err)
			eh.Fail()
			return
		}
		_, err = utils.CopyFilesFrom(testIn, r.scopes, ".snap")
		if err != nil {
			eh.Writef("error copying resultant snap files from %s to %s: %v\n", testIn, r.scopes, err)
			eh.Fail()
			return
		}
	}
}
