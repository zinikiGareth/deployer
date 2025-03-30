package utils

import (
	"fmt"
	"io"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type indentingWriter struct {
	w      io.Writer
	levels []string
}

func (iw *indentingWriter) Intro(format string, args ...any) {
	iw.Indent()
	iw.Printf(format, args...)
}

func (iw *indentingWriter) AttrsWhere(at pluggable.Locatable) {
	iw.Printf(" {\n")
	iw.levels = append(iw.levels, "A")
	iw.Indent()
	iw.Printf("_where_: %s\n", at.Loc().String())
}

func (iw *indentingWriter) TextAttr(field string, value string) {
	iw.Indent()
	iw.Printf("%s: %s\n", field, value)
}

func (iw *indentingWriter) ListAttr(field string) {
	iw.Indent()
	iw.Printf("%s: [\n", field)
	iw.levels = append(iw.levels, "L")
}

func (iw *indentingWriter) EndList() {
	if iw.levels[len(iw.levels)-1] != "L" {
		panic("EndAttrs but top of stack was not L")
	}
	iw.levels = iw.levels[0 : len(iw.levels)-1]
	iw.Indent()
	iw.Printf("]\n")
}

func (iw *indentingWriter) EndAttrs() {
	if iw.levels[len(iw.levels)-1] != "A" {
		panic("EndAttrs but top of stack was not A")
	}
	iw.levels = iw.levels[0 : len(iw.levels)-1]
	iw.Indent()
	iw.Printf("}\n")
}

func (iw *indentingWriter) Indent() {
	for range iw.levels {
		iw.Printf("\t")
	}
}

func (iw *indentingWriter) Printf(format string, args ...any) {
	fmt.Fprintf(iw.w, format, args...)
}

func NewIndentWriter(w io.Writer) pluggable.IndentWriter {
	return &indentingWriter{w: w}
}
