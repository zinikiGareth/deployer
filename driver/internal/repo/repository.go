package repo

import (
	"ziniki.org/deployer/driver/pkg/pluggable"
)

type SimpleRepository struct {
	symbolLsnrs []pluggable.SymbolListener
	symbols     map[pluggable.SymbolName]pluggable.Describable
	tops        []pluggable.TopLevelForm
}

func (d *SimpleRepository) ReadingFile(file string) {
	for _, lsnr := range d.symbolLsnrs {
		lsnr.ReadingFile(file)
	}
}

func (d *SimpleRepository) IntroduceSymbol(who pluggable.SymbolName, is pluggable.Describable) {
	if d.symbols[who] != nil {
		panic("duplicate definition of " + who)
	}
	d.symbols[who] = is
	for _, lsnr := range d.symbolLsnrs {
		lsnr.Symbol(who, is)
	}
}

func (d *SimpleRepository) TopLevel(defn pluggable.TopLevelForm) {
	d.tops = append(d.tops, defn)
}

func (d *SimpleRepository) AddSymbolListener(lsnr pluggable.SymbolListener) {
	d.symbolLsnrs = append(d.symbolLsnrs, lsnr)
}

func (d *SimpleRepository) Traverse(lsnr pluggable.RepositoryTraverser) {
	for who, what := range d.symbols {
		lsnr.Visit(who, what)
	}
}

func (d *SimpleRepository) FindTop(name pluggable.SymbolName) pluggable.TopLevelForm {
	for _, top := range d.tops {
		if top.Name() == name {
			return top
		}
	}
	return nil
}

func NewRepository() pluggable.Repository {
	return &SimpleRepository{symbols: make(map[pluggable.SymbolName]pluggable.Describable)}
}
