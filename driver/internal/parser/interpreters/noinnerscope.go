package interpreters

import "ziniki.org/deployer/driver/pkg/driverbottom"

type disallowInnerScope struct {
	tools *driverbottom.CoreTools
}

func (b *disallowInnerScope) HaveTokens(scope driverbottom.Scope, tokens []driverbottom.Token) driverbottom.Interpreter {
	b.tools.Reporter.Report(0, "nested content is not allowed here")
	return &ignoreInnerScope{}
}

func (b *disallowInnerScope) Completed() {
}
