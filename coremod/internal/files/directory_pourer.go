package files

import (
	"fmt"
	"os"
	"path/filepath"

	"ziniki.org/deployer/coremod/pkg/files"
)

type DirectoryPourer struct {
	Path string
}

func (dp *DirectoryPourer) Relative(s string) (*DirectoryPourer, error) {
	return NewDirectoryPourer(filepath.Join(dp.Path, s))
}

func (dp *DirectoryPourer) PourAll(into files.FileDest) {
	// TODO: just file case
	// TODO: recursive dir case
	files, err := os.ReadDir(dp.Path)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		into.PourInto(f.Name())
	}
}

func (dp *DirectoryPourer) PourOut(name string, into files.FileDest) {
	into.PourInto(name)
}

func NewDirectoryPourer(path string) (*DirectoryPourer, error) {
	if !filepath.IsAbs(path) {
		return nil, fmt.Errorf("not an absolute path")
	}
	dir, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !dir.IsDir() {
		return nil, fmt.Errorf("%s not a directory", path)
	}
	return &DirectoryPourer{Path: path}, nil
}
