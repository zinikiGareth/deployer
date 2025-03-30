package lsnrs

import (
	"fmt"
	"os"
	"path/filepath"

	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type RepoListener struct {
	writeTo *os.File
}

func (r *RepoListener) ReadingFile(file string) {
	r.writeTo.WriteString(file)
	r.writeTo.WriteString(":\n")
}

func (r *RepoListener) Symbol(who pluggable.SymbolName, is pluggable.Definition) {
	r.writeTo.WriteString(fmt.Sprintf("%s %s %s\n", is.Where().InFile(), is.What(), who))
}

func (r *RepoListener) Close() {
	r.writeTo.Sync()
	r.writeTo.Close()
}

func NewRepoListener(outdir string) (*RepoListener, error) {
	path := filepath.Join(outdir, "symbols.repo")
	writeTo, err := os.Create(path)
	if err != nil {
		fmt.Printf("could not save to %s: %v\n", path, err)
		return nil, err
	}

	return &RepoListener{writeTo: writeTo}, nil
}
