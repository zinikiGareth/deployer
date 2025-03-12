package errors

import (
	"fmt"
	"path/filepath"

	"ziniki.org/deployer/deployer/pkg/deployer"
)

type CaseTracker struct {
	caseName    string
	errorDir    string
	failures    map[string][]string
	errhandlers map[string]TestErrorHandler
}

func (tracker *CaseTracker) NewCase(caseName, dir string) {
	tracker.caseName = caseName
	tracker.errorDir = dir
	tracker.errhandlers = make(map[string]TestErrorHandler)
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
	return &FileErrorHandler{tracker: tracker, purpose: purpose, tofile: file}
}

func (tracker *CaseTracker) Fail(area string) {
	fmt.Printf("  FAIL %s\n", area)
	areas := tracker.failures[tracker.caseName]
	areas = append(areas, area)
	tracker.failures[tracker.caseName] = areas
}

func (tracker *CaseTracker) Done() {
	for _, eh := range tracker.errhandlers {
		eh.Close()
	}
}

func (tracker *CaseTracker) Report() int {
	if len(tracker.failures) > 0 {
		fmt.Printf("\n%d failures:\n", len(tracker.failures))
		for f, as := range tracker.failures {
			fmt.Printf("  %s %s\n", f, as)
		}
	}
	if len(tracker.failures) > 127 {
		return 127
	} else {
		return len(tracker.failures)
	}
}

func NewCaseTracker() *CaseTracker {
	return &CaseTracker{failures: make(map[string][]string)}
}
