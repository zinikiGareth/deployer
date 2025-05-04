package runner

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"ziniki.org/deployer/coremod/pkg/coremod"
	"ziniki.org/deployer/deployer/pkg/creator"
	"ziniki.org/deployer/deployer/pkg/deployer"
	sink "ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/utils"
	"ziniki.org/deployer/golden/internal/errors"
	"ziniki.org/deployer/golden/internal/lsnrs"
	"ziniki.org/deployer/golden/pkg/testing"
	"ziniki.org/deployer/testmod/pkg/testmod"
)

type TestRunner struct {
	tracker    *errors.CaseTracker
	deployer   deployer.Deployer
	symbolLsnr *lsnrs.RepoListener
	golden     *goldenComparator
	RunnerPaths
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
	err = utils.EnsureCleanDir(r.prepOut)
	if err != nil {
		return err
	}
	err = utils.EnsureCleanDir(r.execOut)
	if err != nil {
		return err
	}

	r.tracker.NewCase(r.test, r.out)
	r.symbolLsnr, err = lsnrs.NewRepoListener(r.repoOut)
	if err != nil {
		return err
	}
	r.deployer.AddSymbolListener(r.symbolLsnr)

	storage := r.deployer.ObtainStorage()
	register := r.deployer.ObtainRegister()
	tsl, err := testing.NewTestStepLogger(storage, filepath.Join(r.prepOut, "steps.txt"), filepath.Join(r.execOut, "steps.txt"))
	if err != nil {
		return err
	}
	register.ProvideDriver("testhelpers.TestStepLogger", tsl)

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
	if strings.HasSuffix(mod, "coremod.so") {
		return r.loadCoreMod()
	} else if strings.HasSuffix(mod, "testmod.so") {
		return r.loadTestMod()
	}
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

func (r *TestRunner) loadCoreMod() error {
	err := coremod.ProvideTestRunner(r)
	if err != nil {
		return err
	}
	return coremod.RegisterWithDeployer(r.deployer)
}

func (r *TestRunner) loadTestMod() error {
	err := testmod.ProvideTestRunner(r)
	if err != nil {
		return err
	}
	return testmod.RegisterWithDeployer(r.deployer)
}

func NewTestRunner(tracker *errors.CaseTracker, root, test string) (*TestRunner, error) {
	paths := ConfigurePaths(root, test)

	err := utils.EnsureCleanDir(paths.errorsOut)
	if err != nil {
		panic(fmt.Sprintf("error creating error dir %s: %v", paths.errorsOut, err))
	}
	ueTxt := filepath.Join(paths.errorsOut, "usererrors.txt")
	userErrorsTo, err := os.Create(ueTxt)
	if err != nil {
		panic(fmt.Sprintf("error creating error file %s: %v", ueTxt, err))
	}
	sink := sink.NewFileSink(paths.errorFile)
	deployerInst := creator.NewDeployer(sink, userErrorsTo)

	gc := newGoldenComparator(tracker, paths)

	return &TestRunner{tracker: tracker, golden: gc, RunnerPaths: paths, deployer: deployerInst}, nil
}
