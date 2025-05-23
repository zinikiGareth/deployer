package errors

import (
	"fmt"
	"log"
	"os"
)

type writerSink struct {
	path      string
	writer    *os.File
	hasErrors bool
}

func (w *writerSink) Report(loc *Location, msg string) {
	w.hasErrors = true
	if w.writer == nil {
		var err error
		w.writer, err = os.Create(w.path)
		if err != nil {
			log.Printf("cannot create file at %s: %v\n", w.path, err)
			return
		}
	}
	fmt.Fprintf(w.writer, "%3d.%-3d %s\n", loc.Line.Line, loc.Offset, msg)
	fmt.Fprintf(w.writer, "  ==> %s\n", loc.Line.Text)
	w.writer.Sync()
}

func (w *writerSink) Reportf(loc *Location, format string, args ...any) {
	w.Report(loc, fmt.Sprintf(format, args...))
}

func (s *writerSink) HasErrors() bool {
	return s.hasErrors
}

func NewConsoleSink() ErrorSink {
	return &writerSink{writer: os.Stdout}
}

func NewFileSink(path string) ErrorSink {
	return &writerSink{path: path}
}
