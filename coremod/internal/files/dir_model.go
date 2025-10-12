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

func (d *DirModel) ObtainPourer() error {
	if d.pourer != nil {
		return nil
	}
	for _, v := range d.paths {
		var err error
		if d.pourer == nil {
			switch p := v.(type) {
			case *DirectoryPourer:
				d.pourer = p
			case fmt.Stringer:
				d.pourer, err = handleBaseString(p.String())
			case string:
				d.pourer, err = handleBaseString(p)
			default:
				err = fmt.Errorf("cannot handle base path %T", v)
			}
		} else {
			switch p := v.(type) {
			case fmt.Stringer:
				d.pourer, err = handleNestedString(d.pourer, p.String())
			case string:
				d.pourer, err = handleNestedString(d.pourer, p)
			default:
				err = fmt.Errorf("cannot handle nested path %T", v)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func handleBaseString(s string) (*DirectoryPourer, error) {
	if filepath.IsAbs(s) {
		pourer, err := NewDirectoryPourer(s)
		if err != nil {
			return nil, err
		}
		return pourer, nil
	} else {
		return nil, fmt.Errorf("cannot use non-abs path here: %s", s)
	}
}

func handleNestedString(dp *DirectoryPourer, s string) (*DirectoryPourer, error) {
	if !filepath.IsAbs(s) {
		pourer, err := dp.Relative(s)
		if err != nil {
			return nil, err
		}
		return pourer, nil
	} else {
		return nil, fmt.Errorf("cannot use abs path here: %s", s)
	}

}
func (dp *DirModel) PourAll(into corebottom.FileDest) error {
	err := dp.ObtainPourer()
	if err != nil {
		return err
	}
	return dp.pourer.PourAll(into)
}

func (dp *DirModel) PourOut(name string, into corebottom.FileDest) error {
	err := dp.ObtainPourer()
	if err != nil {
		return err
	}
	return dp.pourer.PourOut(name, into)
}

func NewDirModel(loc *errorsink.Location, paths []any) (*DirModel, error) {
	ret := &DirModel{loc: loc, paths: paths}
	err := ret.ObtainPourer() // check that the paths exist as early as possible
	if err != nil {
		return nil, err
	}
	return ret, nil
}

var _ driverbottom.Describable = &DirModel{}
var _ corebottom.FileSource = &DirModel{}
