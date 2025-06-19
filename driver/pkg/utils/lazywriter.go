package utils

import (
	"fmt"
	"io"
	"os"
)

type lazyFileCreator struct {
	name string
	open *os.File
}

func (lfc *lazyFileCreator) WriteString(s string) (n int, err error) {
	if lfc.open == nil {
		lfc.create()
	}
	return lfc.open.WriteString(s)
}

func NewLazyFileCreator(file string) io.StringWriter {
	return &lazyFileCreator{name: file}
}

func (lfc *lazyFileCreator) create() {
	var err error
	lfc.open, err = os.Create(lfc.name)
	if err != nil {
		panic(fmt.Sprintf("error creating error file %s: %v", lfc.name, err))
	}
}
