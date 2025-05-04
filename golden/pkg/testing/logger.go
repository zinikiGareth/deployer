package testing

import (
	"fmt"
	"os"

	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type TestStepLoggerFile struct {
	tools    *pluggable.Tools
	prepFile *os.File
	execFile *os.File
}

func (logger *TestStepLoggerFile) Log(format string, args ...any) {
	var toFile *os.File
	if logger.tools.Storage.IsMode(pluggable.PREPARE_MODE) {
		toFile = logger.prepFile
	} else {
		toFile = logger.execFile
	}
	fmt.Fprintf(toFile, format, args...)
}

func NewTestStepLogger(tools *pluggable.Tools, prepFile string, execFile string) (testhelpers.TestStepLogger, error) {
	prep, err := os.Create(prepFile)
	if err != nil {
		return nil, err
	}
	exec, err := os.Create(execFile)
	if err != nil {
		return nil, err
	}
	return &TestStepLoggerFile{tools: tools, prepFile: prep, execFile: exec}, nil
}
