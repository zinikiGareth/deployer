package errors

import "fmt"

type ErrorReporter struct {
	sink     ErrorSink
	lineNo   int
	lineText string
}

func (r *ErrorReporter) At(lineNo int, lineText string) {
	r.lineNo = lineNo
	r.lineText = lineText
}

func (r *ErrorReporter) Report(indent int, msg string) {
	r.sink.Report(r.lineNo, indent, r.lineText, msg)
}

func (r *ErrorReporter) Reportf(indent int, format string, args ...any) {
	r.sink.Report(r.lineNo, indent, r.lineText, fmt.Sprintf(format, args...))
}

func (r *ErrorReporter) Sink() ErrorSink {
	return r.sink
}

func NewErrorReporter(sink ErrorSink) *ErrorReporter {
	return &ErrorReporter{sink: sink}
}
