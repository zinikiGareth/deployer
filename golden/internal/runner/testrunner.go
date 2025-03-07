package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type TestRunner struct {
	base  string
	input []string
}

func (r *TestRunner) Run() {
	fmt.Printf("%s\n", r.base)
	for _, f := range r.input {
		fmt.Printf("Process input file %s\n", f)
	}
}

func NewTestRunner(root, test string) (*TestRunner, error) {
	base := filepath.Join(root, test)
	scripts := filepath.Join(base, "scripts")
	files, err := os.ReadDir(scripts)
	if err != nil {
		return nil, fmt.Errorf("could not read script directory %s: %v", scripts, err)
	}
	deployFiles := make([]string, 0)
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".dply") {
			deployFiles = append(deployFiles, f.Name())
		}
	}
	return &TestRunner{base: base, input: deployFiles}, nil
}
