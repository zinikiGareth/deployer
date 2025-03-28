package errors

type ErrorSink interface {
	Report(lineNo int, indent int, lineText string, msg string)
}

type ErrorRepI interface {
	At(lineNo int, lineText string)
	Report(indent int, msg string)
}
