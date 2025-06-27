package files

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CopyModel struct {
	loc  *errorsink.Location
	Src  corebottom.FileSource
	Dest corebottom.DestHolder
}

func (c *CopyModel) Loc() *errorsink.Location {
	return c.loc
}

// ShortDescription implements driverbottom.Describable.
func (c *CopyModel) ShortDescription() string {
	return "CopyModel[]"
}

func (c *CopyModel) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("CopyModel\n")
}

var _ driverbottom.Describable = &CopyModel{}
