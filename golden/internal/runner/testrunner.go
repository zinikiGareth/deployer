package runner

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"plugin"

	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/utils"
	"ziniki.org/deployer/golden/internal/errors"
)

type TestRunner struct {
	tracker  *errors.CaseTracker
	deployer *deployer.Deployer
	base     string
	test     string
	out      string
	scripts  string
	scopes   string
}

func (r *TestRunner) Run(modules []string) {
	err := r.Setup(modules)
	if err != nil {
		fmt.Printf("Error during setup: %v\n", err)
		return
	}

	r.TestScopes(r.tracker.ErrorHandlerFor("scopes").(errors.TestErrorHandler))
	r.TestDeployment(r.tracker.ErrorHandlerFor("deploy").(errors.TestErrorHandler))

	r.WrapUp()
}

func (r *TestRunner) Setup(modules []string) error {
	fmt.Printf("%s:\n", r.test)
	err := utils.EnsureCleanDir(r.out)
	if err != nil {
		return err
	}

	r.tracker.NewCase(r.test, r.out)

	return r.LoadModules(modules)
}

func (r *TestRunner) LoadModules(modules []string) error {
	for _, m := range modules {
		err := r.Module(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *TestRunner) Module(mod string) error {
	p, err := plugin.Open(mod)
	if err != nil {
		return err
	}
	test, err := p.Lookup("ProvideTestRunner")
	if err == nil {
		err = test.(func(deployer.TestRunner) error)(r)
		if err != nil {
			return err
		}
	}
	init, err := p.Lookup("RegisterWithDeployer")
	if err != nil {
		log.Printf("ignoring module " + mod + " as it does not have RegisterWithDeployer")
		return nil
	}
	return init.(func(*deployer.Deployer) error)(r.deployer)
}

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
		cmd := exec.Command("vscode-tmgrammar-snap", "--config", "../../../vsix/package.json", testIn+"/*.dply")
		cmd.Dir = r.base
		cmd.Stdout = eh
		cmd.Stderr = eh
		err := cmd.Run()
		if err != nil {
			eh.Fail()
			return
		}
	} else {
		cmd := exec.Command("vscode-tmgrammar-snap", "--config", "../../../vsix/package.json", "--updateSnapshot", testIn+"/*.dply")
		cmd.Dir = r.base
		cmd.Stdout = eh
		cmd.Stderr = eh
		err := cmd.Run()
		if err != nil {
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

func (r *TestRunner) TestDeployment(eh errors.TestErrorHandler) {
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

func (r *TestRunner) WrapUp() {
	r.tracker.Done()
}

func (r *TestRunner) ErrorHandlerFor(purpose string) deployer.ErrorHandler {
	return r.tracker.ErrorHandlerFor(purpose)
}

// These things belong elsewhere ...
func NewTestRunner(tracker *errors.CaseTracker, root, test string) (*TestRunner, error) {
	base := filepath.Join(root, test)
	outdir := filepath.Join(base, "out")
	scripts := filepath.Join(base, "scripts")
	scopes := filepath.Join(base, "scopes")

	deployerInst := deployer.NewDeployer()

	return &TestRunner{tracker: tracker, base: base, out: outdir, test: test, scripts: scripts, scopes: scopes, deployer: deployerInst}, nil
}
