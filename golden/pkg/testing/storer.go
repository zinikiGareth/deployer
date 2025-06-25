package testing

import (
	"os"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/testhelpers"
	"ziniki.org/deployer/driver/pkg/utils"
)

type TestStorerFile struct {
	tools   *corebottom.Tools
	files   map[int]*os.File
	writers map[int]driverbottom.IndentWriter
}

func (logger *TestStorerFile) GetWriter(mode int) driverbottom.IndentWriter {
	if logger.writers[mode] != nil {
		return logger.writers[mode]
	}
	writer := utils.NewIndentWriter(logger.files[mode])
	logger.writers[mode] = writer
	return writer
}

func NewTestStorer(tools *corebottom.Tools, resolveFile string, findFile string, prepFile string, execFile string) (testhelpers.TestStorer, error) {
	storeAs := make(map[int]*os.File)
	resolve, err := os.Create(resolveFile)
	if err != nil {
		return nil, err
	}
	storeAs[corebottom.RESOLVE_MODE] = resolve
	find, err := os.Create(findFile)
	if err != nil {
		return nil, err
	}
	storeAs[corebottom.DETERMINE_INITIAL_MODE] = find
	prep, err := os.Create(prepFile)
	if err != nil {
		return nil, err
	}
	storeAs[corebottom.DETERMINE_DESIRED_MODE] = prep
	exec, err := os.Create(execFile)
	if err != nil {
		return nil, err
	}
	storeAs[corebottom.UPDATE_REALITY_MODE] = exec
	storeAs[corebottom.TEARDOWN_MODE] = exec
	return &TestStorerFile{tools: tools, files: storeAs, writers: make(map[int]driverbottom.IndentWriter)}, nil
}

var _ testhelpers.TestStorer = &TestStorerFile{}
