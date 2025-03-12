package runner

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"plugin"

	"ziniki.org/deployer/deployer/pkg/creator"
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/utils"
	"ziniki.org/deployer/golden/internal/errors"
	"ziniki.org/deployer/golden/internal/lsnrs"
)

type TestRunner struct {
	tracker    *errors.CaseTracker
	deployer   deployer.Deployer
	symbolLsnr *lsnrs.RepoListener
	base       string
	test       string
	out        string
	scripts    string
	scopes     string
	repoIn     string
	repoOut    string
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
	err = utils.EnsureCleanDir(r.repoOut)
	if err != nil {
		return err
	}

	r.tracker.NewCase(r.test, r.out)
	r.symbolLsnr, err = lsnrs.NewRepoListener(r.repoOut)
	if err != nil {
		return err
	}
	r.deployer.AddSymbolListener(r.symbolLsnr)

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
	return init.(func(deployer.Deployer) error)(r.deployer)
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
			eh.Writef("failed running vscode-tmgrammar-snap: %v\n", err)
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
	r.compareGoldenFiles(r.repoIn, r.repoOut)
}

// TODO: I would like to make this its own thing
func (r *TestRunner) compareGoldenFiles(golden, gen string) {
	eh := r.tracker.ErrorHandlerFor("golden")
	eh.Writef("comparing %s to %s\n", golden, gen)
	goldenFiles, err1 := utils.FindFiles(golden, "")
	genFiles, err2 := utils.FindFiles(gen, "")
	if err1 != nil && err2 != nil {
		// it's "safe" to assume it's an empty case
		return
	}
	if err1 != nil {
		// OK, it should be there.  Create it and we'll copy the files in
		utils.EnsureDir(golden)
	}
	if err2 != nil {
		// Presumably if there is a golden dir, there should be a gen dir
		eh.Writef("error collecting generated files from %s\n", gen)
		return
	}

	// Go through the golden files, comparing to the generated ones
	genmap := make(map[string]int)
	for k, g := range genFiles {
		genmap[g] = k + 1
	}
	for _, f := range goldenFiles {
		if genmap[f] != 0 {
			if !utils.CompareFiles(filepath.Join(gen, f), filepath.Join(golden, f)) {
				eh.Writef("generated file %s did not match golden file\n", f)
				eh.Fail()
			}
			delete(genmap, f)
		} else { // if there is no generated file, complain: that's a failure
			eh.Writef("there is no gen file for %s\n", f)
			eh.Fail()
		}
	}
	// If there are any generated files which don't have golden files, let the user know and copy them
	if len(genmap) > 0 {
		eh.Writef("generated files were not present ... copying\n")
		for f := range genmap {
			fmt.Printf("  %s\n", f)
			utils.CopyFile(filepath.Join(gen, f), filepath.Join(golden, f))
		}
	}
}

func (r *TestRunner) WrapUp() {
	r.symbolLsnr.Close()
	r.tracker.Done()
}

func (r *TestRunner) ErrorHandlerFor(purpose string) deployer.ErrorHandler {
	return r.tracker.ErrorHandlerFor(purpose)
}

func NewTestRunner(tracker *errors.CaseTracker, root, test string) (*TestRunner, error) {
	base := filepath.Join(root, test)
	outdir := filepath.Join(base, "out")
	repoin := filepath.Join(base, "repository")
	repoout := filepath.Join(base, "repository-gen")
	scripts := filepath.Join(base, "scripts")
	scopes := filepath.Join(base, "scopes")

	deployerInst := creator.NewDeployer()

	return &TestRunner{tracker: tracker, base: base, out: outdir, test: test, scripts: scripts, scopes: scopes, repoIn: repoin, repoOut: repoout, deployer: deployerInst}, nil
}
