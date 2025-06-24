package files

import (
	"fmt"
	"path/filepath"

	"ziniki.org/deployer/coremod/pkg/corebottom"
)

type DirModel struct {
	paths  []any
	pourer *DirectoryPourer
}

func (d *DirModel) ObtainPourer() {
	var err error
	for _, v := range d.paths {
		if d.pourer == nil {
			p, ok := v.(*DirectoryPourer)
			if ok {
				d.pourer = p
			} else {
				tmp, ok := v.(fmt.Stringer)
				if ok {
					v = tmp.String()
				}
				s, ok := v.(string)
				if ok {
					if filepath.IsAbs(s) {
						d.pourer, err = NewDirectoryPourer(s)
						if err != nil {
							panic(err)
						}
					} else {
						panic(fmt.Sprintf("cannot use non-abs path here: %v\n", v))
					}
				} else {
					panic(fmt.Sprintf("cannot handle base path %T\n", v))
				}
			}
		} else {
			s, ok := v.(string)
			if ok {
				if !filepath.IsAbs(s) {
					d.pourer, err = d.pourer.Relative(s)
					if err != nil {
						panic(err)
					}
				} else {
					panic(fmt.Sprintf("cannot use abs path here: %v\n", v))
				}
			} else {
				panic(fmt.Sprintf("cannot handle nested path %T\n", v))
			}
		}
	}
}

func (dp *DirModel) PourAll(into corebottom.FileDest) {
	dp.ObtainPourer()
	dp.pourer.PourAll(into)
}

func (dp *DirModel) PourOut(name string, into corebottom.FileDest) {
	dp.ObtainPourer()
	dp.pourer.PourOut(name, into)
}

func NewDirModel(paths []any) *DirModel {
	return &DirModel{paths: paths}
}

var _ corebottom.FileSource = &DirModel{}
