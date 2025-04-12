package repo

import (
	"log"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func (d *SimpleRepository) ResolveAll(sink errors.ErrorSink) {
	for who, what := range d.symbols {
		searcher := &Searcher{}
		log.Printf("resolve %s %s\n", who, what.ShortDescription())
		what.Resolve(searcher)
	}
}

type Searcher struct {
}

func (*Searcher) Resolve(name pluggable.Identifier) pluggable.Definition {
	return nil
}
