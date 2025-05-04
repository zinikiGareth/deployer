package repo

import (
	"log"
	"reflect"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func (repo *SimpleRepository) ResolveAll(sink errors.ErrorSink, register pluggable.Recall) {
	for _, what := range repo.symbols {
		searcher := &Searcher{repo: repo, recall: register, sink: sink}
		what.Resolve(searcher)
	}
}

func (d *SimpleRepository) GetDefinition(name pluggable.SymbolName) pluggable.Definition {
	return d.symbols[name]
}

type Searcher struct {
	repo   *SimpleRepository
	recall pluggable.Recall
	sink   errors.ErrorSink
}

func (s *Searcher) Resolve(name pluggable.Identifier) pluggable.Noun {
	defn := s.repo.GetDefinition(pluggable.SymbolName(name.Id()))
	ret, ok := defn.(pluggable.Noun)
	if ret != nil && ok {
		return ret
	}
	ret = s.recall.Find(reflect.TypeFor[pluggable.Noun](), name.Id()).(pluggable.Noun)
	if ret != nil {
		return ret
	}
	log.Printf("failed to resolve %s\n", name)
	s.sink.Reportf(name.Loc(), "could not resolve symbol %s", name.Id())
	return nil
}
