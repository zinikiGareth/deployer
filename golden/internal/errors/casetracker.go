package errors

import "ziniki.org/deployer/deployer/pkg/deployer"

type CaseTracker struct {
	errorDir    string
	errhandlers map[string]TestErrorHandler
}

func (tracker *CaseTracker) UseDirectory(dir string) {
	tracker.errorDir = dir
}

func (tracker *CaseTracker) ErrorHandlerFor(what string) deployer.ErrorHandler {
	eh := tracker.errhandlers[what]
	if eh == nil {
		eh = NewErrorHandler(tracker.errorDir, what)
		tracker.errhandlers[what] = eh
	}
	return eh
}

func (tracker *CaseTracker) Done() {
	for _, eh := range tracker.errhandlers {
		eh.Close()
	}
}

func NewCaseTracker() *CaseTracker {
	return &CaseTracker{errhandlers: make(map[string]TestErrorHandler)}
}
