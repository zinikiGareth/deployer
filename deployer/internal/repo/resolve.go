package repo

import (
	"log"
	"reflect"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func (repo *SimpleRepository) ResolveAll(tools *pluggable.Tools) {
	for _, what := range repo.tops {
		searcher := &Searcher{repo: repo, recall: tools.Recall, reporter: tools.Reporter}
		what.Resolve(searcher)
	}
}

func (d *SimpleRepository) GetDefinition(name pluggable.SymbolName) pluggable.Locatable {
	return d.symbols[name]
}

type Searcher struct {
	repo     *SimpleRepository
	recall   pluggable.Recall
	reporter errors.ErrorRepI
}

func (s *Searcher) Resolve(name pluggable.Identifier) pluggable.Blank {
	defn := s.repo.GetDefinition(pluggable.SymbolName(name.Id()))
	ret, ok := defn.(pluggable.Blank)
	if ret != nil && ok {
		return ret
	}
	ret = s.recall.Find(reflect.TypeFor[pluggable.Blank](), name.Id()).(pluggable.Blank)
	if ret != nil {
		return ret
	}
	log.Printf("failed to resolve %s\n", name)
	s.reporter.At(name.Loc().Line)
	s.reporter.Reportf(name.Loc().Offset, "could not resolve symbol %s", name.Id())
	return nil
}
