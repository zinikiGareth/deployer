package errors

type ErrorSink interface {
	Report(lineNo int, indent int, lineText string, msg string)
}
