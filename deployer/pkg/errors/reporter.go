package errors

type ErrorReporter struct {
	sink ErrorSink
}

func (r *ErrorReporter) At(lineNo int, lineText string) {

}

func (r *ErrorReporter) Report(indent int, msg string) {
}

func NewErrorReporter(sink ErrorSink) *ErrorReporter {
	return &ErrorReporter{sink: sink}
}
