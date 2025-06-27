package vars

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type TopLevelAttacher struct {
	tools *corebottom.Tools
}

func (a *TopLevelAttacher) MakeAssign(holder driverbottom.Holder, assignTo driverbottom.Identifier, action any) any {
	ret := MakeDoAssign(a.tools, holder, assignTo, action)
	return ret
}

func (a *TopLevelAttacher) Attach(tlf any) {
	top, ok := tlf.(driverbottom.TopLevelForm)
	if !ok {
		panic("not a TLF")
	}
	a.tools.Repository.TopLevel(top)
	a.tools.Repository.IntroduceSymbol(top.Name(), top)

}

type TLACreator struct {
	tools *corebottom.Tools
}

func (c *TLACreator) Create() driverbottom.AttachResult {
	return &TopLevelAttacher{tools: c.tools}
}

func NewMakeTopLevelAttacher(tools *corebottom.Tools) *TLACreator {
	return &TLACreator{tools: tools}
}

var _ driverbottom.AttacherCreator = &TLACreator{}
