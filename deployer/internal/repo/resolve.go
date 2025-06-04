package repo

import (
	"log"
	"reflect"

	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func (repo *SimpleRepository) ResolveAll(tools *pluggable.Tools) {
	for _, what := range repo.tops {
		searcher := &Searcher{repo: repo, recall: tools.Recall, reporter: tools.Reporter}
		what.Resolve(searcher)
	}
}

func (d *SimpleRepository) GetDefinition(name pluggable.SymbolName) pluggable.Describable {
	return d.symbols[name]
}

type Searcher struct {
	repo     *SimpleRepository
	recall   pluggable.Recall
	reporter errorsink.ErrorRepI
}

func (s *Searcher) Resolve(name pluggable.Identifier) pluggable.Describable {
	// log.Printf("attempting to resolve %s\n", name)
	defn := s.repo.GetDefinition(pluggable.SymbolName(name.Id()))
	if defn != nil {
		return defn
	}
	ret, ok := s.recall.Find(reflect.TypeFor[pluggable.Blank](), name.Id()).(pluggable.Describable)
	if ret != nil && ok {
		return ret
	}
	log.Printf("failed to resolve %s\n", name)
	s.reporter.At(name.Loc().Line)
	s.reporter.Reportf(name.Loc().Offset, "could not resolve symbol %s", name.Id())
	return nil
}
