package lsnrs

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/utils"
)

type RepositoryStorer interface {
	driverbottom.RepositoryTraverser
	DumpNamesTo(outdir string)
	DumpDefnsTo(outdir string)
}

type goldenRepoStorer struct {
	defns map[driverbottom.SymbolName][]driverbottom.Describable
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
		for _, m := range s.defns[key] {
			fmt.Fprintf(writeTo, "%s => %s\n", key, m.ShortDescription())
		}
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
		for _, d := range s.defns[key] {
			iw.IndPrintf("symbol %s is bound to:\n", key)
			d.DumpTo(iw)
		}
	}
	writeTo.Close()
}

func (s *goldenRepoStorer) Visit(who driverbottom.SymbolName, what driverbottom.Describable) {
	s.defns[who] = append(s.defns[who], what)
}

func NewGoldenRepoStorer() RepositoryStorer {
	return &goldenRepoStorer{defns: make(map[driverbottom.SymbolName][]driverbottom.Describable)}
}
