package coretop

import (
	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/coremod/internal/files"
	"ziniki.org/deployer/coremod/internal/lists"
	"ziniki.org/deployer/coremod/internal/methods"
	"ziniki.org/deployer/coremod/internal/policy"
	"ziniki.org/deployer/coremod/internal/runmain"
	"ziniki.org/deployer/coremod/internal/target"
	"ziniki.org/deployer/coremod/internal/time"
	"ziniki.org/deployer/coremod/internal/vars"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

var testRunner driverbottom.TestRunner

func ProvideTestRunner(runner driverbottom.TestRunner) error {
	testRunner = runner
	return nil
}

func RegisterWithDriver(driver driverbottom.Driver) error {
	if testRunner != nil {
		eh := testRunner.ErrorHandlerFor("log")
		eh.WriteMsg("Installing things from coremod\n")
	}

	ct := driver.ObtainCoreTools()
	tools := corebottom.NewTools(ct, &corebottom.Options{})
	tools.CoreTools.StoreOther("coremod", tools)

	// Logically, I think, these three have to go in "deployer", not in a module.
	tools.Register.ExtensionPoint("main-args")
	tools.Register.ExtensionPoint("attacher")
	tools.Register.ExtensionPoint("top-level")
	tools.Register.ExtensionPoint("function-defn")

	tools.Register.ExtensionPoint("target")
	tools.Register.ExtensionPoint("policy-statements")
	tools.Register.ExtensionPoint("policy-inner")

	tools.Register.ExtensionPoint("blank")

	tools.Register.ExtensionPoint("prop-interpreter")

	// TODO: I think we should be registering "creation methods" not the created objects ...

	// we need to register something that handles the main arguments, in our case targets to execute
	tools.Register.Register("main-args", "main", runmain.MakeMainHandler(tools))

	// for variables, we need something that can attach these to the top level scope (do we, in fact?)
	tools.Register.Register("attacher", "top-level", vars.NewMakeTopLevelAttacher(tools))

	// top commands
	tools.Register.Register("top-level", "target", target.MakeCoreTargetVerb(tools))

	// target commands
	tools.Register.Register("target", "ensure", basic.NewEnsureCommandHandler(tools))
	tools.Register.Register("target", "find", basic.NewFindCommandHandler(tools))
	tools.Register.Register("target", "coin", basic.NewMemoryCoinCommandHandler(tools))
	tools.Register.Register("target", "env", basic.NewEnvCommandHandler(tools))
	tools.Register.Register("target", "show", basic.NewShowCommandHandler(tools))

	tools.Register.Register("target", "files.dir", files.NewDirCommandHandler(tools))
	tools.Register.Register("target", "files.copy", files.NewCopyCommandHandler(tools))

	tools.Register.Register("target", "policy", policy.NewPolicyCommandHandler(tools))
	tools.Register.Register("target", "attachPolicy", policy.NewAttachPolicyCommandHandler(tools))

	tools.Register.Register("policy-statements", "allow", policy.NewPolicyAllowCommandHandler(tools))
	tools.Register.Register("policy-inner", "action", policy.NewPolicyActionCommandHandler(tools))
	tools.Register.Register("policy-inner", "condition", policy.NewPolicyConditionCommandHandler(tools))
	tools.Register.Register("policy-inner", "principal", policy.NewPolicyPrincipalCommandHandler(tools))

	// functions
	tools.Register.Register("function-defn", "->", methods.MakeInvokeFunc(tools))
	tools.Register.Register("function-defn", "hours", time.MakeHoursFunc(tools))
	tools.Register.Register("function-defn", "sum", lists.MakeSumFunc(tools))

	return nil
}
