package errors

import (
	"fmt"
	"log"
	"os"
)

type writerSink struct {
	path   string
	writer *os.File
}

func (w *writerSink) Report(lineNo int, indent int, lineText string, msg string) {
	if w.writer == nil {
		var err error
		w.writer, err = os.Create(w.path)
		if err != nil {
			log.Printf("cannot create file at %s: %v\n", w.path, err)
			return
		}
	}
	fmt.Fprintf(w.writer, "%3d.%-3d %s\n", lineNo, indent, msg)
	fmt.Fprintf(w.writer, "  ==> %s\n", lineText)
	w.writer.Sync()
}

func (w *writerSink) Reportf(lineNo int, indent int, lineText string, format string, args ...any) {
	w.Report(lineNo, indent, lineText, fmt.Sprintf(format, args...))
}

func NewConsoleSink() ErrorSink {
	return &writerSink{writer: os.Stdout}
}

func NewFileSink(path string) ErrorSink {
	return &writerSink{path: path}
}
