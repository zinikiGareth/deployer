package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"ziniki.org/deployer/coremod/pkg/files"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type copyAction struct {
	tools    *pluggable.Tools
	loc      *errors.Location
	exprs    []pluggable.Expr
	assignTo pluggable.Identifier
}

func (ca *copyAction) Loc() *errors.Location {
	return ca.loc
}

func (ca *copyAction) Where() *errors.Location {
	return ca.loc
}

func (ca *copyAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("CopyAction")
	w.AttrsWhere(ca)
	for _, v := range ca.exprs {
		w.IndPrintf("%s\n", v.String())
	}
	w.EndAttrs()
}

func (ca *copyAction) ShortDescription() string {
	return fmt.Sprintf("Dir[%d]", len(ca.exprs))
}

func (ca *copyAction) Completed() {
}

func (ca *copyAction) Resolve(r pluggable.Resolver) {
	// ea.resolved = r.Resolve(ea.what)
}

func (ca *copyAction) Prepare() {
	// Not quite sure what to do here ...
	// Need to prepare
	// Should check things like permissions
	copyFrom := ca.tools.Storage.Eval(ca.exprs[0])
	copyFS, ok := copyFrom.(*files.Path)
	if !ok {
		panic("not a path")
	}
	path := copyFS.File
	if !filepath.IsAbs(path) {
		panic("not an absolute path")
	}
	dir, err := os.Stat(path)
	if err != nil {
		log.Fatalf("stat file failed: %v", err)
	}
	if !dir.IsDir() {
		log.Fatalf("%s not a directory", path)
	}
	return
}

func (ca *copyAction) Execute() {
	srcVar := ca.tools.Storage.Eval(ca.exprs[0])
	src, ok := srcVar.(*files.Path)
	if !ok {
		panic("not the bucket i was looking for")
	}
	destVar := ca.tools.Storage.Eval(ca.exprs[1])
	dest, ok := destVar.(files.ThingyHolder)
	if !ok {
		panic("not the bucket i was looking for")
	}
	files, err := os.ReadDir(src.File)
	if err != nil {
		panic(err)
	}

	d := dest.ObtainDest()
	for _, f := range files {
		d.Pour(f.Name())
	}
}
