package coremod

import (
	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/coremod/internal/files"
	"ziniki.org/deployer/coremod/internal/methods"
	"ziniki.org/deployer/coremod/internal/policy"
	"ziniki.org/deployer/coremod/internal/target"
	"ziniki.org/deployer/coremod/internal/time"
	"ziniki.org/deployer/deployer/pkg/deployer"
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

	tools.Register.ExtensionPoint("top-level")
	tools.Register.ExtensionPoint("target")
	tools.Register.ExtensionPoint("policy-statements")
	tools.Register.ExtensionPoint("policy-inner")

	tools.Register.ExtensionPoint("blank")

	tools.Register.ExtensionPoint("function-defn")

	// TODO: I think we should be registering "creation methods" not the created objects ...

	// top commands
	tools.Register.Register("top-level", "target", target.MakeCoreTargetVerb(tools))

	// target commands
	tools.Register.Register("target", "ensure", basic.NewEnsureCommandHandler(tools))
	tools.Register.Register("target", "find", basic.NewFindCommandHandler(tools))
	tools.Register.Register("target", "env", basic.NewEnvCommandHandler(tools))
	tools.Register.Register("target", "show", basic.NewShowCommandHandler(tools))

	tools.Register.Register("target", "files.dir", files.NewDirCommandHandler(tools))
	tools.Register.Register("target", "files.copy", files.NewCopyCommandHandler(tools))

	tools.Register.Register("target", "policy", policy.NewPolicyCommandHandler(tools))

	tools.Register.Register("policy-statements", "allow", policy.NewPolicyAllowCommandHandler(tools))
	tools.Register.Register("policy-inner", "condition", policy.NewPolicyConditionCommandHandler(tools))

	// functions
	tools.Register.Register("function-defn", "->", methods.MakeInvokeFunc(tools))
	tools.Register.Register("function-defn", "hours", time.MakeHoursFunc(tools))

	return nil
}
