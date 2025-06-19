package errors

import (
	"fmt"
	"path/filepath"
	"slices"

	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type CaseTracker struct {
	caseName        string
	errorDir        string
	failuresInOrder []string
	failures        map[string][]string
	errhandlers     map[string]TestErrorHandler
}

func (tracker *CaseTracker) NewCase(caseName, dir string) {
	tracker.caseName = caseName
	tracker.errorDir = dir
	tracker.errhandlers = make(map[string]TestErrorHandler)
}

func (tracker *CaseTracker) ErrorHandlerFor(what string) driverbottom.ErrorHandler {
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
	if !slices.Contains(tracker.failuresInOrder, tracker.caseName) {
		tracker.failuresInOrder = append(tracker.failuresInOrder, tracker.caseName)
	}
	areas := tracker.failures[tracker.caseName]
	if !slices.Contains(areas, area) {
		fmt.Printf("  FAIL %s\n", area)
		areas = append(areas, area)
		tracker.failures[tracker.caseName] = areas
	}
}

func (tracker *CaseTracker) Done() {
	for _, eh := range tracker.errhandlers {
		eh.Close()
	}
}

func (tracker *CaseTracker) Report() int {
	if len(tracker.failures) > 0 {
		fmt.Printf("\n%d failures:\n", len(tracker.failures))
		for _, f := range tracker.failuresInOrder {
			fmt.Printf("  %s %s\n", f, tracker.failures[f])
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
