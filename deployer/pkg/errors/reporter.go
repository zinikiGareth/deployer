package errors

import "fmt"

type ErrorReporter struct {
	sink ErrorSink
	line *LineLoc
}

func (r *ErrorReporter) At(line *LineLoc) {
	r.line = line
}

func (r *ErrorReporter) Report(offset int, msg string) {
	r.sink.Report(r.line.Location(offset), msg)
}

func (r *ErrorReporter) Reportf(offset int, format string, args ...any) {
	r.sink.Report(r.line.Location(offset), fmt.Sprintf(format, args...))
}

func (r *ErrorReporter) Sink() ErrorSink {
	return r.sink
}

func NewErrorReporter(sink ErrorSink) *ErrorReporter {
	return &ErrorReporter{sink: sink}
}
