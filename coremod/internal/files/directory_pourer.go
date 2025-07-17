package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"ziniki.org/deployer/coremod/pkg/corebottom"
)

type DirectoryPourer struct {
	Path string
}

func (dp *DirectoryPourer) Relative(s string) (*DirectoryPourer, error) {
	return NewDirectoryPourer(filepath.Join(dp.Path, s))
}

func (dp *DirectoryPourer) PourAll(into corebottom.FileDest) {
	// TODO: just file case
	// TODO: recursive dir case
	files, err := os.ReadDir(dp.Path)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.IsDir() {
			fromdir, err := dp.Relative(f.Name())
			if err != nil {
				panic(err)
			}
			intodir, err := into.Relative(f.Name())
			if err != nil {
				panic(err)
			}
			fromdir.PourAll(intodir)
		} else {
			dp.PourOut(f.Name(), into)
		}
	}
}

func (dp *DirectoryPourer) PourOut(name string, into corebottom.FileDest) {
	full := filepath.Join(dp.Path, name)
	fd, err := os.Open(full)
	if err != nil {
		log.Printf("could not open %s\n", full)
		return
	}
	defer fd.Close()
	into.PourInto(name, fd)
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
