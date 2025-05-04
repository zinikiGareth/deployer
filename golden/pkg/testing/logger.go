package testing

import (
	"fmt"
	"os"

	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

type TestStepLoggerFile struct {
	storage  pluggable.RuntimeStorage
	prepFile *os.File
	execFile *os.File
}

func (logger *TestStepLoggerFile) Log(format string, args ...any) {
	var toFile *os.File
	if logger.storage.IsMode(pluggable.PREPARE_MODE) {
		toFile = logger.prepFile
	} else {
		toFile = logger.execFile
	}
	fmt.Fprintf(toFile, format, args...)
}

func NewTestStepLogger(storage pluggable.RuntimeStorage, prepFile string, execFile string) (testhelpers.TestStepLogger, error) {
	prep, err := os.Create(prepFile)
	if err != nil {
		return nil, err
	}
	exec, err := os.Create(execFile)
	if err != nil {
		return nil, err
	}
	return &TestStepLoggerFile{storage: storage, prepFile: prep, execFile: exec}, nil
}
