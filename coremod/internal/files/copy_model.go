package files

import "ziniki.org/deployer/coremod/pkg/corebottom"

type CopyModel struct {
	Src  corebottom.FileSource
	Dest corebottom.DestHolder
}
