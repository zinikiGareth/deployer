package utils

import (
	"fmt"
	"io"

	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type indentingWriter struct {
	w      io.Writer
	levels []string
}

func (iw *indentingWriter) Intro(format string, args ...any) {
	iw.showIndent()
	iw.Printf(format, args...)
}

func (iw *indentingWriter) AttrsWhere(at driverbottom.Locatable) {
	iw.Printf(" {\n")
	iw.levels = append(iw.levels, "A")
	iw.showIndent()
	iw.Printf("_where_: %s\n", at.Loc().String())
}

func (iw *indentingWriter) TextAttr(field string, value string) {
	iw.showIndent()
	iw.Printf("%s: %s\n", field, value)
}

func (iw *indentingWriter) ListAttr(field string) {
	iw.showIndent()
	iw.Printf("%s: [\n", field)
	iw.levels = append(iw.levels, "L")
}

func (iw *indentingWriter) EndList() {
	if iw.levels[len(iw.levels)-1] != "L" {
		panic("EndList but top of stack was not L")
	}
	iw.levels = iw.levels[0 : len(iw.levels)-1]
	iw.showIndent()
	iw.Printf("]\n")
}

func (iw *indentingWriter) NestedAttr(field string, obj driverbottom.Describable) {
	iw.showIndent()
	iw.Printf("%s:\n", field)
	iw.levels = append(iw.levels, "N")
	obj.DumpTo(iw)
	if iw.levels[len(iw.levels)-1] != "N" {
		panic("EndNested but top of stack was not L")
	}
	iw.levels = iw.levels[0 : len(iw.levels)-1]
}

func (iw *indentingWriter) EndAttrs() {
	if iw.levels[len(iw.levels)-1] != "A" {
		panic("EndAttrs but top of stack was not A")
	}
	iw.levels = iw.levels[0 : len(iw.levels)-1]
	iw.showIndent()
	iw.Printf("}\n")
}

func (iw *indentingWriter) Indent() {
	iw.levels = append(iw.levels, "I")
}

func (iw *indentingWriter) UnIndent() {
	if iw.levels[len(iw.levels)-1] != "I" {
		panic("UnIndent but top of stack was not I")
	}
	iw.levels = iw.levels[0 : len(iw.levels)-1]
}

func (iw *indentingWriter) showIndent() {
	for range iw.levels {
		iw.Printf("\t")
	}
}

func (iw *indentingWriter) IndPrintf(format string, args ...any) {
	iw.showIndent()
	iw.Printf(format, args...)
}

func (iw *indentingWriter) Printf(format string, args ...any) {
	fmt.Fprintf(iw.w, format, args...)
}

func NewIndentWriter(w io.Writer) driverbottom.IndentWriter {
	return &indentingWriter{w: w}
}
