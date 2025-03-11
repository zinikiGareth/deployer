package errors

import (
	"fmt"
	"path/filepath"

	"ziniki.org/deployer/deployer/pkg/deployer"
)

type CaseTracker struct {
	caseName    string
	failures    []string
	errorDir    string
	errhandlers map[string]TestErrorHandler
}

func (tracker *CaseTracker) NewCase(caseName, dir string) {
	tracker.caseName = caseName
	tracker.errorDir = dir
}

func (tracker *CaseTracker) ErrorHandlerFor(what string) deployer.ErrorHandler {
	eh := tracker.errhandlers[what]
	if eh == nil {
		eh = tracker.NewErrorHandler(what)
		tracker.errhandlers[what] = eh
	}
	return eh
}

func (tracker *CaseTracker) NewErrorHandler(purpose string) *FileErrorHandler {
	file := filepath.Join(tracker.errorDir, "errors-"+purpose)
	return &FileErrorHandler{tracker: tracker, tofile: file}
}

func (tracker *CaseTracker) Fail() {
	tracker.failures = append(tracker.failures, tracker.caseName)
}
func (tracker *CaseTracker) Done() {
	for _, eh := range tracker.errhandlers {
		eh.Close()
	}
}

func (tracker *CaseTracker) Report() int {
	if len(tracker.failures) > 0 {
		fmt.Printf("%d failures:\n", len(tracker.failures))
		for _, f := range tracker.failures {
			fmt.Printf("  %s\n", f)
		}
	}
	if len(tracker.failures) > 127 {
		return 127
	} else {
		return len(tracker.failures)
	}
}

func NewCaseTracker() *CaseTracker {
	return &CaseTracker{errhandlers: make(map[string]TestErrorHandler)}
}
