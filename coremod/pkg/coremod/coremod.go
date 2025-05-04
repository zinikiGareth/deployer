package coremod

import (
	"reflect"

	"ziniki.org/deployer/coremod/internal/target"
	"ziniki.org/deployer/coremod/internal/time"
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

var testRunner deployer.TestRunner

func ProvideTestRunner(runner deployer.TestRunner) error {
	testRunner = runner
	return nil
}

func RegisterWithDeployer(deployer deployer.Deployer) error {
	if testRunner != nil {
		eh := testRunner.ErrorHandlerFor("log")
		eh.WriteMsg("Installing things from coremod\n")
	}
	register := deployer.ObtainRegister()
	register.Register(reflect.TypeFor[pluggable.TargetCommand](), "target", &target.CoreTargetVerb{})
	register.Register(reflect.TypeFor[pluggable.Function](), "hours", &time.HoursFunc{})
	return nil
}
