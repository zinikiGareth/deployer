package parser

import "ziniki.org/deployer/driver/pkg/driverbottom"

type TopLevelAttacher struct {
	tools *driverbottom.CoreTools
}

func (a *TopLevelAttacher) Attach(tlf any) {
	top, ok := tlf.(driverbottom.TopLevelForm)
	if !ok {
		panic("not a TLF")
	}
	a.tools.Repository.TopLevel(top)
	a.tools.Repository.IntroduceSymbol(top.Name(), top)

}
func NewTopLevelAttacher(tools *driverbottom.CoreTools) driverbottom.AttachResult {
	return &TopLevelAttacher{tools: tools}
}
