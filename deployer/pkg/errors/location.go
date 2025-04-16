package errors

import (
	"fmt"
	"path/filepath"
)

type FileLoc struct {
	File string
}

type LineLoc struct {
	File   *FileLoc
	Line   int
	Indent int
	Text   string
}

type Location struct {
	Line   *LineLoc
	Offset int
}

func (loc Location) InFile() string {
	return fmt.Sprintf("%d.%d", loc.Line.Line, loc.Offset)
}

func (loc Location) String() string {
	return fmt.Sprintf("%s:%d.%d", filepath.Base(loc.Line.File.File), loc.Line.Line, loc.Offset)
}

func InFile(name string) *FileLoc {
	return &FileLoc{File: name}
}

func (f *FileLoc) AtLine(line, indent int, text string) *LineLoc {
	return &LineLoc{File: f, Line: line, Indent: indent, Text: text}
}

func (ll *LineLoc) Location(offset int) *Location {
	return &Location{Line: ll, Offset: offset}
}

/*
func NewLocation(file string, line, offset int) Location {
	return Location{File: file, Line: line, Offset: offset}
}
*/
