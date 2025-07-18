package basicmath

import "ziniki.org/deployer/driver/pkg/driverbottom"

func RegisterAll(tools *driverbottom.CoreTools) {
	tools.Register.Register("function-defn", "*", MakeMultiplyFunc(tools))
}
