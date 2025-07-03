package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	cwd, err := os.Getwd()
	if err != nil {
		cwd = ""
	}
	iw.Intro("DirModel")
	iw.AttrsWhere(d)
	for _, p := range d.paths {
		str := ""
		switch p := p.(type) {
		case string:
			str = p
		case driverbottom.Describable:
			str = p.ShortDescription()
		case fmt.Stringer:
			str = p.String()
		default:
			str = fmt.Sprintf("%T", p)
		}
		if cwd != "" {
			str = strings.Replace(str, cwd, "<HOME>", -1)
		}
		iw.TextAttr("path", str)
	}
	iw.EndAttrs()
}

func (d *DirModel) ObtainPourer() {
	if d.pourer != nil {
		return
	}
	for _, v := range d.paths {
		if d.pourer == nil {
			switch p := v.(type) {
			case *DirectoryPourer:
				d.pourer = p
			case fmt.Stringer:
				d.pourer = handleBaseString(p.String())
			case string:
				d.pourer = handleBaseString(p)
			default:
				panic(fmt.Sprintf("cannot handle base path %T\n", v))
			}
		} else {
			switch p := v.(type) {
			case fmt.Stringer:
				d.pourer = handleNestedString(d.pourer, p.String())
			case string:
				d.pourer = handleNestedString(d.pourer, p)
			default:
				panic(fmt.Sprintf("cannot handle nested path %T\n", v))
			}
		}
	}
}

func handleBaseString(s string) *DirectoryPourer {
	if filepath.IsAbs(s) {
		pourer, err := NewDirectoryPourer(s)
		if err != nil {
			panic(err)
		}
		return pourer
	} else {
		panic(fmt.Sprintf("cannot use non-abs path here: %s\n", s))
	}
}

func handleNestedString(dp *DirectoryPourer, s string) *DirectoryPourer {
	if !filepath.IsAbs(s) {
		pourer, err := dp.Relative(s)
		if err != nil {
			panic(err)
		}
		return pourer
	} else {
		panic(fmt.Sprintf("cannot use abs path here: %s\n", s))
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
	// for k, v := range paths {
	// 	log.Printf("path %d %v\n", k, v)
	// }
	ret := &DirModel{loc: loc, paths: paths}
	// TODO: this needs a little bit of care, because there could be issues with things not being resolved yet
	// We need to trap those
	ret.ObtainPourer() // check that the paths exist as early as possible
	return ret
}

var _ driverbottom.Describable = &DirModel{}
var _ corebottom.FileSource = &DirModel{}
