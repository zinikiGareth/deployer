package testing

import (
	"fmt"
	"os"

	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/pluggable"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

type TestStepLoggerFile struct {
	tools    *external.Tools
	prepFile *os.File
	execFile *os.File
}

func (logger *TestStepLoggerFile) Log(format string, args ...any) {
	var toFile *os.File
	if logger.tools.Storage.IsMode(pluggable.BUILD_MODEL_MODE) {
		toFile = logger.prepFile
	} else {
		toFile = logger.execFile
	}
	fmt.Fprintf(toFile, format, args...)
}

func NewTestStepLogger(tools *external.Tools, prepFile string, execFile string) (testhelpers.TestStepLogger, error) {
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
