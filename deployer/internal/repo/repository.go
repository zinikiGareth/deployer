package repo

import (
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type SimpleRepository struct {
	symbolLsnrs []pluggable.SymbolListener
	symbols     map[pluggable.SymbolName]pluggable.Describable
	tops        []pluggable.TargetThing
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

func (d *SimpleRepository) TopLevel(defn pluggable.TargetThing) {
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

func (d *SimpleRepository) FindTarget(name pluggable.SymbolName) pluggable.TargetThing {
	defn := d.symbols[name]
	if defn == nil {
		return nil
	}
	target, ok := defn.(pluggable.TargetThing)
	if !ok {
		return nil
	}
	return target
}

func NewRepository() pluggable.Repository {
	return &SimpleRepository{symbols: make(map[pluggable.SymbolName]pluggable.Describable)}
}
