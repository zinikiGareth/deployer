package parser

import "ziniki.org/deployer/deployer/pkg/pluggable"

type TopLevelAttacher struct {
	tools *pluggable.CoreTools
}

func (a *TopLevelAttacher) Attach(tlf any) {
	top, ok := tlf.(pluggable.TopLevelForm)
	if !ok {
		panic("not a TLF")
	}
	a.tools.Repository.TopLevel(top)
	a.tools.Repository.IntroduceSymbol(top.Name(), top)

}
func NewTopLevelAttacher(tools *pluggable.CoreTools) pluggable.AttachResult {
	return &TopLevelAttacher{tools: tools}
}
