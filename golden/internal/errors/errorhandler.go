package errors

import (
	"fmt"
	"os"

	"ziniki.org/deployer/deployer/pkg/deployer"
)

type TestErrorHandler interface {
	deployer.ErrorHandler
	Fail()
}

type FileErrorHandler struct {
	tracker *CaseTracker
	purpose string
	tofile  string
	file    *os.File
	failed  bool
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
	eh.tracker.Fail(eh.purpose)
}

func (eh *FileErrorHandler) Close() {
	if eh.file != nil {
		eh.file.Close()
	}
}

func (eh *FileErrorHandler) ensureOpen() error {
	if eh.file != nil {
		return nil
	}
	var err error
	eh.file, err = os.Create(eh.tofile)
	return err
}
