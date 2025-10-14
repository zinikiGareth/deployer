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

type PathType int

const (
	DIR_TYPE PathType = iota
	FILE_TYPE
	ANY_TYPE
)

type DirModel struct {
	loc       *errorsink.Location
	paths     []any
	pourer    *DirectoryPourer
	mustExist bool
	mustBe    PathType
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

func (d *DirModel) ObtainPourer() (*DirectoryPourer, error) {
	if d.pourer != nil {
		return d.pourer, nil
	}
	for _, v := range d.paths {
		var err error
		if d.pourer == nil {
			switch p := v.(type) {
			case *DirectoryPourer:
				d.pourer = p
			case fmt.Stringer:
				d.pourer, err = d.handleBaseString(p.String())
			case string:
				d.pourer, err = d.handleBaseString(p)
			default:
				err = fmt.Errorf("cannot handle base path %T", v)
			}
		} else {
			switch p := v.(type) {
			case fmt.Stringer:
				d.pourer, err = d.handleNestedString(d.pourer, p.String())
			case string:
				d.pourer, err = d.handleNestedString(d.pourer, p)
			default:
				err = fmt.Errorf("cannot handle nested path %T", v)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	err := d.pourer.AssertConstraints()
	if err != nil {
		return nil, err
	} else {
		return d.pourer, nil
	}
}

func (d *DirModel) handleBaseString(s string) (*DirectoryPourer, error) {
	if filepath.IsAbs(s) {
		pourer, err := NewDirectoryPourer(s, d.mustExist, d.mustBe)
		if err != nil {
			return nil, err
		}
		return pourer, nil
	} else {
		return nil, fmt.Errorf("cannot use non-abs path here: %s", s)
	}
}

func (d *DirModel) handleNestedString(dp *DirectoryPourer, s string) (*DirectoryPourer, error) {
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
	pourer, err := dp.ObtainPourer()
	if err != nil {
		return err
	}
	return pourer.PourAll(into)
}

func (dp *DirModel) PourOut(name string, into corebottom.FileDest) error {
	pourer, err := dp.ObtainPourer()
	if err != nil {
		return err
	}
	pourer, err = pourer.Relative(name)
	if err != nil {
		return err
	}
	return pourer.PourOut(into)
}

func NewDirModel(loc *errorsink.Location, paths []any) (*DirModel, error) {
	ret := &DirModel{loc: loc, paths: paths, mustExist: true, mustBe: DIR_TYPE}
	_, err := ret.ObtainPourer() // check that the paths exist as early as possible
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func NewFileModel(loc *errorsink.Location, paths []any) (*DirModel, error) {
	ret := &DirModel{loc: loc, paths: paths, mustExist: true, mustBe: FILE_TYPE}
	_, err := ret.ObtainPourer() // check that the paths exist as early as possible
	if err != nil {
		return nil, err
	}
	return ret, nil
}

var _ driverbottom.Describable = &DirModel{}
var _ corebottom.FileSource = &DirModel{}
