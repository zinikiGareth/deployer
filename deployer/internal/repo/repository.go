package repo

import "ziniki.org/deployer/deployer/pkg/pluggable"

type SimpleRepository struct {
	symbolLsnrs []pluggable.SymbolListener
}

func (d *SimpleRepository) ReadingFile(file string) {
	for _, lsnr := range d.symbolLsnrs {
		lsnr.ReadingFile(file)
	}
}

func (d *SimpleRepository) IntroduceSymbol(where pluggable.Location, what pluggable.SymbolType, who pluggable.SymbolName) {
	for _, lsnr := range d.symbolLsnrs {
		lsnr.Symbol(where, what, who)
	}
}

func (d *SimpleRepository) AddSymbolListener(lsnr pluggable.SymbolListener) {
	d.symbolLsnrs = append(d.symbolLsnrs, lsnr)
}

func NewRepository() pluggable.Repository {
	return &SimpleRepository{}
}
