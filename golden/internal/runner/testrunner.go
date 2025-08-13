package runner

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/coremod/pkg/coretop"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/drivertop"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/utils"
	"ziniki.org/deployer/golden/internal/errors"
	"ziniki.org/deployer/golden/internal/lsnrs"
	"ziniki.org/deployer/golden/pkg/testing"
	"ziniki.org/deployer/testmod/pkg/testmod"
)

type TestRunner struct {
	tracker    *errors.CaseTracker
	driver     driverbottom.Driver
	deployer   corebottom.Deployer
	symbolLsnr *lsnrs.RepoListener
	golden     *goldenComparator
	RunnerPaths
}

func (r *TestRunner) Tools() *driverbottom.CoreTools {
	return r.driver.ObtainCoreTools()
}

func (r *TestRunner) Run(modules []string) {
	err := r.Setup(modules)
	if err != nil {
		fmt.Printf("Error during setup: %v\n", err)
		return
	}

	ignoreFile := filepath.Join(r.scripts, "ignore")
	_, err = os.Stat(ignoreFile)
	if err == nil {
		log.Printf("ignoring test in %s\n", r.scripts)
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
	err = utils.EnsureCleanDir(r.resolveOut)
	if err != nil {
		return err
	}
	err = utils.EnsureCleanDir(r.findOut)
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
	r.driver.AddSymbolListener(r.symbolLsnr)

	tools := r.deployer.ObtainTools()
	tsl, err := testing.NewTestStepLogger(tools, filepath.Join(r.resolveOut, "steps.txt"), filepath.Join(r.findOut, "steps.txt"), filepath.Join(r.prepOut, "steps.txt"), filepath.Join(r.execOut, "steps.txt"))
	if err != nil {
		return err
	}
	tools.Register.ProvideDriver("testhelpers.TestStepLogger", tsl)

	ts, err := testing.NewTestStorer(tools, filepath.Join(r.resolveOut, "vars.txt"), filepath.Join(r.findOut, "vars.txt"), filepath.Join(r.prepOut, "vars.txt"), filepath.Join(r.execOut, "vars.txt"))
	if err != nil {
		return err
	}

	tools.Register.ProvideDriver("testhelpers.TestStorer", ts)

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
		err = test.(func(driverbottom.TestRunner) error)(r)
		if err != nil {
			return err
		}
	}
	init, err := p.Lookup("RegisterWithDriver")
	if err != nil {
		log.Printf("ignoring module %s as it does not have RegisterWithDriver", mod)
		return nil
	}
	return init.(func(driverbottom.Driver) error)(r.driver)
}

func (r *TestRunner) loadCoreMod() error {
	err := coretop.ProvideTestRunner(r)
	if err != nil {
		return err
	}
	return coretop.RegisterWithDriver(r.driver)
}

func (r *TestRunner) loadTestMod() error {
	err := testmod.ProvideTestRunner(r)
	if err != nil {
		return err
	}
	return testmod.RegisterWithDriver(r.driver)
}

func NewTestRunner(tracker *errors.CaseTracker, root, test string) (*TestRunner, error) {
	paths := ConfigurePaths(root, test)

	err := utils.EnsureCleanDir(paths.errorsOut)
	if err != nil {
		panic(fmt.Sprintf("error creating error dir %s: %v", paths.errorsOut, err))
	}
	ueTxt := filepath.Join(paths.errorsOut, "usererrors.txt")
	userErrorsTo := utils.NewLazyFileCreator(ueTxt)
	sink := errorsink.NewFileSink(paths.errorFile)

	driverInst := drivertop.NewDriver(sink, userErrorsTo)
	tools := corebottom.NewTools(driverInst.ObtainCoreTools(), &corebottom.Options{})
	deployerInst := coretop.NewDeployer(driverInst, tools)

	gc := newGoldenComparator(tracker, paths)

	return &TestRunner{tracker: tracker, golden: gc, RunnerPaths: paths, driver: driverInst, deployer: deployerInst}, nil
}
