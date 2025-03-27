package errors

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

func NewErrorReporter(sink ErrorSink) *ErrorReporter {
	return &ErrorReporter{sink: sink}
}
