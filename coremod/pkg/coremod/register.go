package coremod

import (
	"reflect"

	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/coremod/internal/files"
	"ziniki.org/deployer/coremod/internal/methods"
	"ziniki.org/deployer/coremod/internal/policy"
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

	tools := deployer.ObtainTools()

	// top commands
	tools.Register.Register(reflect.TypeFor[pluggable.VerbCommand](), "target", target.MakeCoreTargetVerb(tools))

	// target commands
	tools.Register.Register(reflect.TypeFor[pluggable.VerbCommand](), "ensure", basic.NewEnsureCommandHandler(tools))
	tools.Register.Register(reflect.TypeFor[pluggable.VerbCommand](), "find", basic.NewFindCommandHandler(tools))
	tools.Register.Register(reflect.TypeFor[pluggable.VerbCommand](), "env", basic.NewEnvCommandHandler(tools))
	tools.Register.Register(reflect.TypeFor[pluggable.VerbCommand](), "show", basic.NewShowCommandHandler(tools))

	tools.Register.Register(reflect.TypeFor[pluggable.VerbCommand](), "files.dir", files.NewDirCommandHandler(tools))
	tools.Register.Register(reflect.TypeFor[pluggable.VerbCommand](), "files.copy", files.NewCopyCommandHandler(tools))

	tools.Register.Register(reflect.TypeFor[pluggable.VerbCommand](), "policy", policy.NewPolicyCommandHandler(tools))

	// functions
	tools.Register.Register(reflect.TypeFor[pluggable.Function](), "->", methods.MakeInvokeFunc(tools))
	tools.Register.Register(reflect.TypeFor[pluggable.Function](), "hours", time.MakeHoursFunc(tools))

	return nil
}
