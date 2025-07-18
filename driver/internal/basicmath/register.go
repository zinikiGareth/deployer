package basicmath

import "ziniki.org/deployer/driver/pkg/driverbottom"

func RegisterAll(tools *driverbottom.CoreTools) {
	tools.Register.Register("function-defn", "*", MakeMultiplyFunc(tools))
	tools.Register.Register("function-defn", "/", MakeDivFunc(tools))
	tools.Register.Register("function-defn", "+", MakeAddFunc(tools))
	tools.Register.Register("function-defn", "-", MakeSubFunc(tools))
}
