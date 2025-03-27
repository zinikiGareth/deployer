package errors

import (
	"fmt"
	"os"
)

type writerSink struct {
	writer *os.File
}

func (w *writerSink) Report(lineNo int, indent int, lineText string, msg string) {
	fmt.Fprintf(w.writer, "%3d.%-3d %s\n", lineNo, indent, msg)
	fmt.Fprintf(w.writer, "  ==> %s\n", lineText)
	w.writer.Sync()
}

func NewConsoleSink() ErrorSink {
	return &writerSink{writer: os.Stdout}
}

func NewFileSink(path string) (ErrorSink, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return &writerSink{writer: file}, nil
}
