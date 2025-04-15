package testing

import (
	"fmt"
	"os"
)

type TestStepLogger interface {
	Log(fmt string, args ...any)
}

type TestStepLoggerFile struct {
	toFile *os.File
}

func (logger *TestStepLoggerFile) Log(format string, args ...any) {
	fmt.Fprintf(logger.toFile, format, args...)
}

func NewTestStepLogger(toFile string) (TestStepLogger, error) {
	out, err := os.Create(toFile)
	if err != nil {
		return nil, err
	}
	return &TestStepLoggerFile{toFile: out}, nil
}
