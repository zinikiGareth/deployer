package errors

type ErrorSink interface {
	Report(loc *Location, msg string)
	Reportf(loc *Location, format string, args ...any)
	HasErrors() bool
}

type ErrorRepI interface {
	At(line *LineLoc)
	Report(offset int, msg string)
	Reportf(offset int, fmt string, args ...any)
}
