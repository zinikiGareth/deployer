package lsnrs

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"

	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/utils"
)

type RepositoryStorer interface {
	pluggable.RepositoryTraverser
	DumpNamesTo(outdir string)
	DumpDefnsTo(outdir string)
}

type goldenRepoStorer struct {
	defns map[pluggable.SymbolName]pluggable.Action
}

func (s *goldenRepoStorer) DumpNamesTo(outdir string) {
	path := filepath.Join(outdir, "names.repo")
	writeTo, err := os.Create(path)
	if err != nil {
		fmt.Printf("could not save to %s: %v\n", path, err)
		return
	}
	keys := slices.Collect(maps.Keys(s.defns))
	slices.Sort(keys)
	for _, key := range keys {
		fmt.Fprintf(writeTo, "%s => %s\n", key, s.defns[key].ShortDescription())
	}
	writeTo.Close()
}

func (s *goldenRepoStorer) DumpDefnsTo(outdir string) {
	path := filepath.Join(outdir, "defns.repo")
	writeTo, err := os.Create(path)
	if err != nil {
		fmt.Printf("could not save to %s: %v\n", path, err)
		return
	}
	iw := utils.NewIndentWriter(writeTo)
	keys := slices.Collect(maps.Keys(s.defns))
	slices.Sort(keys)
	for _, key := range keys {
		d := s.defns[key]
		iw.IndPrintf("symbol %s is bound to:\n", key)
		d.DumpTo(iw)
	}
	writeTo.Close()
}

func (s *goldenRepoStorer) Visit(who pluggable.SymbolName, what pluggable.Action) {
	s.defns[who] = what
}

func NewGoldenRepoStorer() RepositoryStorer {
	return &goldenRepoStorer{defns: make(map[pluggable.SymbolName]pluggable.Action)}
}
