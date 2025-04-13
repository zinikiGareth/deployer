package errors

type ErrorSink interface {
	Report(lineNo int, indent int, lineText string, msg string)
	Reportf(lineNo int, indent int, lineText string, format string, args ...any)
	HasErrors() bool
}

type ErrorRepI interface {
	At(lineNo int, lineText string)
	Report(indent int, msg string)
	Reportf(indent int, fmt string, args ...any)
}
