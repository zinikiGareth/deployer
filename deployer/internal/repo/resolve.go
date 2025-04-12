package repo

import (
	"log"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func (repo *SimpleRepository) ResolveAll(sink errors.ErrorSink, register pluggable.Recall) {
	for _, what := range repo.symbols {
		searcher := &Searcher{repo: repo, recall: register, sink: sink}
		what.Resolve(searcher)
	}
}

func (d *SimpleRepository) GetDefinition(name pluggable.SymbolName) pluggable.Noun {
	return d.symbols[name]
}

type Searcher struct {
	repo   *SimpleRepository
	recall pluggable.Recall
	sink   errors.ErrorSink
}

func (s *Searcher) Resolve(name pluggable.Identifier) pluggable.Noun {
	ret := s.repo.GetDefinition(pluggable.SymbolName(name.Id()))
	if ret != nil {
		return ret
	}
	ret = s.recall.FindNoun(name.Id())
	if ret != nil {
		return ret
	}
	log.Printf("failed to resolve %s\n", name)
	s.sink.Reportf(name.Loc().Line, name.Loc().Offset, name.Loc().String(), "could not resolve symbol %s", name.Id())
	return nil
}
