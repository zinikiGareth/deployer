package files

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"ziniki.org/deployer/coremod/pkg/corebottom"
)

type DirectoryPourer struct {
	Path      string
	mustExist bool
	mustBe    PathType
}

func (dp *DirectoryPourer) Relative(s string) (*DirectoryPourer, error) {
	return NewDirectoryPourer(filepath.Join(dp.Path, s), dp.mustExist, dp.mustBe)
}

func (dp *DirectoryPourer) AssertConstraints() error {
	if dp.mustExist {
		dir, err := os.Stat(dp.Path)
		if err != nil {
			return err
		}
		if dp.mustBe == DIR_TYPE && !dir.IsDir() {
			return fmt.Errorf("%s not a directory", dp.Path)
		} else if dp.mustBe == FILE_TYPE && dir.IsDir() {
			return fmt.Errorf("%s was a directory", dp.Path)
		}
	}
	return nil
}

func (dp *DirectoryPourer) PourAll(into corebottom.FileDest) error {
	switch dp.mustBe {
	case FILE_TYPE:
		return dp.PourOut(into)
	case DIR_TYPE:
		return dp.pourDir(into)
	case ANY_TYPE:
		dir, err := os.Stat(dp.Path)
		if err != nil {
			return err
		}
		if dir.IsDir() {
			return dp.pourDir(into)
		} else {
			return dp.PourOut(into)
		}
	default:
		panic("unknown file type")
	}
}

func (dp *DirectoryPourer) pourDir(into corebottom.FileDest) error {
	files, err := os.ReadDir(dp.Path)
	if err != nil {
		return err
	}

	for _, f := range files {
		from, err := dp.Relative(f.Name())
		if err != nil {
			return err
		}
		if f.IsDir() {
			intodir, err := into.Relative(f.Name())
			if err != nil {
				return err
			}
			from.PourAll(intodir)
		} else {
			err := from.PourOut(into)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (dp *DirectoryPourer) PourOut(into corebottom.FileDest) error {
	fd, err := os.Open(dp.Path)
	if err != nil {
		return fmt.Errorf("could not open %s", dp.Path)
	}
	defer fd.Close()
	return into.PourInto(path.Base(dp.Path), fd)
}

func NewDirectoryPourer(path string, mustExist bool, mustBe PathType) (*DirectoryPourer, error) {
	if !filepath.IsAbs(path) {
		return nil, fmt.Errorf("not an absolute path")
	}
	return &DirectoryPourer{Path: path, mustExist: mustExist, mustBe: mustBe}, nil
}
