package runner

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type ErrorHandler interface {
	io.Writer
	WriteMsg(msg string)
	Writef(fmt string, args ...any)
	Fail()
	Close()
}

type FileErrorHandler struct {
	tofile string
	file   *os.File
	failed bool
}

func (eh *FileErrorHandler) Write(bs []byte) (int, error) {
	err := eh.ensureOpen()
	if err != nil {
		fmt.Printf("cannot open %s: %v\n", eh.tofile, err)
		return 0, err
	}
	return eh.file.Write(bs)
}

func (eh *FileErrorHandler) WriteMsg(msg string) {
	err := eh.ensureOpen()
	if err != nil {
		fmt.Printf("cannot open %s: %v\n", eh.tofile, err)
		return
	}
	eh.file.WriteString(msg)
}

func (eh *FileErrorHandler) Writef(msg string, args ...any) {
	err := eh.ensureOpen()
	if err != nil {
		fmt.Printf("cannot open %s: %v\n", eh.tofile, err)
		return
	}
	eh.file.WriteString(fmt.Sprintf(msg, args...))
}

func (eh *FileErrorHandler) Fail() {
	eh.failed = true
}

func (eh *FileErrorHandler) Close() {
	eh.file.Close()
}

func (eh *FileErrorHandler) ensureOpen() error {
	if eh.file != nil {
		return nil
	}
	var err error
	eh.file, err = os.Create(eh.tofile)
	return err
}

// TODO: this should be part of something bigger
func NewErrorHandler(outdir, purpose string) ErrorHandler {
	file := filepath.Join(outdir, "errors-"+purpose)
	return &FileErrorHandler{tofile: file}
}
