package files

import (
	"fmt"
	"path/filepath"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type DirModel struct {
	loc    *errorsink.Location
	paths  []any
	pourer *DirectoryPourer
}

func (d *DirModel) Loc() *errorsink.Location {
	return d.loc
}

func (d *DirModel) ShortDescription() string {
	return "DirModel[]"
}

func (d *DirModel) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("DirModel")
	iw.AttrsWhere(d)
	for _, p := range d.paths {
		switch p := p.(type) {
		case string:
			iw.TextAttr("path", p)
		case driverbottom.Describable:
			iw.TextAttr("path", p.ShortDescription())
		case fmt.Stringer:
			iw.TextAttr("path", p.String())
		default:
			iw.TextAttr("path", fmt.Sprintf("%T", p))
		}
	}
	iw.EndAttrs()
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

func NewDirModel(loc *errorsink.Location, paths []any) *DirModel {
	return &DirModel{loc: loc, paths: paths}
}

var _ driverbottom.Describable = &DirModel{}
var _ corebottom.FileSource = &DirModel{}
