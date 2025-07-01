package vars

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type TopLevelAttacher struct {
	tools *corebottom.Tools
	scope driverbottom.Scope
}

func (a *TopLevelAttacher) MakeAssign(holder driverbottom.Holder, assignTo driverbottom.Identifier, action any) any {
	ret := MakeDoAssign(a.tools, holder, assignTo, action)
	return ret
}

func (a *TopLevelAttacher) Attach(tlf any) error {
	top, ok := tlf.(driverbottom.TopLevelForm)
	if !ok {
		return fmt.Errorf("cannot attach %T to top level; not a TLF", tlf)
	}
	a.tools.Repository.TopLevel(top)
	if err := a.scope.IntroduceSymbol(top.Name(), top); err != nil {
		return fmt.Errorf("duplicate target %s", top.Name())
	}

	return nil
}

type TLACreator struct {
	tools *corebottom.Tools
}

func (c *TLACreator) Create(scope driverbottom.Scope) driverbottom.AttachResult {
	return &TopLevelAttacher{tools: c.tools, scope: scope}
}

func NewMakeTopLevelAttacher(tools *corebottom.Tools) *TLACreator {
	return &TLACreator{tools: tools}
}

var _ driverbottom.AttacherCreator = &TLACreator{}
