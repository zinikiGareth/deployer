package runner

import (
	"fmt"
	"log"
	"path/filepath"
	"plugin"

	"ziniki.org/deployer/deployer/pkg/deployer"
)

type TestRunner struct {
	deployer *deployer.Deployer
	base     string
	test     string
	scripts  string
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

func NewTestRunner(root, test string) (*TestRunner, error) {
	base := filepath.Join(root, test)
	scripts := filepath.Join(base, "scripts")
	deployer := deployer.NewDeployer()

	return &TestRunner{base: base, test: test, scripts: scripts, deployer: deployer}, nil
}
