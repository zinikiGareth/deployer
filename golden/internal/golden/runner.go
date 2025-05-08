package golden

import (
	"fmt"
	"strings"

	"ziniki.org/deployer/golden/internal/errors"
	"ziniki.org/deployer/golden/internal/runner"
)

type GoldenRunner struct {
	tracker  *errors.CaseTracker
	modules  []string
	patterns []string
	testdirs []string
}

func (r *GoldenRunner) UseModule(path string) {
	r.modules = append(r.modules, path)
}

func (r *GoldenRunner) MatchPattern(patt string) {
	r.patterns = append(r.modules, patt)
}

func (r *GoldenRunner) RunTestsUnder(root string) {
	r.testdirs = append(r.testdirs, root)
}

func (r *GoldenRunner) RunAll() {
	for _, p := range r.testdirs {
		r.runOne(p)
	}
}

func (r *GoldenRunner) Report() int {
	return r.tracker.Report()
}

func (r *GoldenRunner) runOne(root string) {
	merged := gatherTestsInOrder(root)
	for _, s := range merged {
		if r.matchesPatterns(s) {
			r.runCase(root, s)
		}
	}
}

func (r *GoldenRunner) matchesPatterns(cs string) bool {
	if len(r.patterns) == 0 {
		return true
	}
	for _, s := range r.patterns {
		if strings.Contains(cs, s) {
			return true
		}
	}
	return false
}

func (r *GoldenRunner) runCase(root, dir string) {
	run, err := runner.NewTestRunner(r.tracker, root, dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	run.Run(r.modules)
}
