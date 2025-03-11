package runner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"

	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/utils"
)

type TestRunner struct {
	deployer *deployer.Deployer
	base     string
	test     string
	scripts  string
	scopes   string
}

func (r *TestRunner) Module(mod string) error {
	p, err := plugin.Open(mod)
	if err != nil {
		return err
	}
	init, err := p.Lookup("RegisterWithDeployer")
	if err != nil {
		log.Printf("ignoring module " + mod + " as it does not have RegisterWithDeployer")
		return nil
	}
	return init.(func(*deployer.Deployer) error)(r.deployer)
}

func (r *TestRunner) Run() {
	fmt.Printf("%s:\n", r.test)
	r.TestScopes()
	err := r.deployer.ReadScriptsFrom(r.scripts)
	if err != nil {
		fmt.Printf("Error reading scripts from %s: %v\n", r.scripts, err)
		return
	}
	err = r.deployer.Deploy()
	if err != nil {
		fmt.Printf("Error deploying: %v\n", err)
		return
	}
}

func (r *TestRunner) TestScopes() {
	testIn := filepath.Join(r.base, "scope-test")

	// Make sure clean directory exists
	_, err := os.Stat(testIn)
	if err == nil {
		err = os.RemoveAll(testIn)
		if err != nil {
			fmt.Printf("error deleting %s: %v\n", testIn, err)
			return
		}
	}
	err = os.MkdirAll(testIn, 0777)
	if err != nil {
		fmt.Printf("error creating %s: %v\n", testIn, err)
		return
	}

	// Now copy all the files across
	nin, err := utils.CopyFilesFrom(r.scripts, testIn, ".dply")
	if err != nil {
		fmt.Printf("error copying files from %s to %s: %v\n", r.scripts, testIn, err)
		return
	}
	nout, err := utils.CopyFilesFrom(r.scopes, testIn, ".snap")
	if err != nil {
		fmt.Printf("error copying files from %s to %s: %v\n", r.scopes, testIn, err)
		return
	}
	fmt.Printf("%d, %d\n", nin, nout)
	if nin == 0 && nout == 0 {
		fmt.Printf("no input or output files in %s\n", r.test)
	} else if nin == nout {
		fmt.Printf("run test\n")
		cmd := exec.Command("vscode-tmgrammar-snap", "--config", "../../../vsix/package.json", testIn+"/*.dply")
		cmd.Dir = r.base
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Printf("error running scopes:test: %v\n", err)
			return
		}
	} else {
		fmt.Printf("run update in %s\n", r.base)
		cmd := exec.Command("vscode-tmgrammar-snap", "--config", "../../../vsix/package.json", "--updateSnapshot", testIn+"/*.dply")
		cmd.Dir = r.base
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Printf("error running scopes:test: %v\n", err)
			return
		}
		_, err = utils.CopyFilesFrom(testIn, r.scopes, ".snap")
		if err != nil {
			fmt.Printf("error copying resultant snap files from %s to %s: %v\n", testIn, r.scopes, err)
			return
		}
	}
}

func NewTestRunner(root, test string) (*TestRunner, error) {
	base := filepath.Join(root, test)
	scripts := filepath.Join(base, "scripts")
	scopes := filepath.Join(base, "scopes")
	deployer := deployer.NewDeployer()

	return &TestRunner{base: base, test: test, scripts: scripts, scopes: scopes, deployer: deployer}, nil
}
