package testing

import (
	"fmt"
	"os"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

type TestStepLoggerFile struct {
	tools          *corebottom.Tools
	resolveLogFile *os.File
	findLogFile    *os.File
	buildLogFile   *os.File
	execLogFile    *os.File
}

func (logger *TestStepLoggerFile) Log(format string, args ...any) {
	var toFile *os.File
	if logger.tools.Storage.IsMode(corebottom.RESOLVE_MODE) {
		toFile = logger.resolveLogFile
	} else if logger.tools.Storage.IsMode(corebottom.DETERMINE_INITIAL_MODE) {
		toFile = logger.findLogFile
	} else if logger.tools.Storage.IsMode(corebottom.DETERMINE_DESIRED_MODE) {
		toFile = logger.buildLogFile
	} else {
		toFile = logger.execLogFile
	}
	fmt.Fprintf(toFile, format, args...)
}

func NewTestStepLogger(tools *corebottom.Tools, resolveFile string, findFile string, prepFile string, execFile string) (testhelpers.TestStepLogger, error) {
	resolve, err := os.Create(resolveFile)
	if err != nil {
		return nil, err
	}
	find, err := os.Create(findFile)
	if err != nil {
		return nil, err
	}
	prep, err := os.Create(prepFile)
	if err != nil {
		return nil, err
	}
	exec, err := os.Create(execFile)
	if err != nil {
		return nil, err
	}
	return &TestStepLoggerFile{tools: tools, resolveLogFile: resolve, findLogFile: find, buildLogFile: prep, execLogFile: exec}, nil
}
