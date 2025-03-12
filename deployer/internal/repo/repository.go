package repo

import "ziniki.org/deployer/deployer/pkg/deployer"

type Repository interface {
	ReadingFile(file string)
	IntroduceSymbol(where deployer.Location, what deployer.SymbolType, who deployer.SymbolName)
	AddSymbolListener(lsnr deployer.SymbolListener)
}

type SimpleRepository struct {
	symbolLsnrs []deployer.SymbolListener
}

func (d *SimpleRepository) ReadingFile(file string) {
	for _, lsnr := range d.symbolLsnrs {
		lsnr.ReadingFile(file)
	}
}

func (d *SimpleRepository) IntroduceSymbol(where deployer.Location, what deployer.SymbolType, who deployer.SymbolName) {
	for _, lsnr := range d.symbolLsnrs {
		lsnr.Symbol(where, what, who)
	}
}

func (d *SimpleRepository) AddSymbolListener(lsnr deployer.SymbolListener) {
	d.symbolLsnrs = append(d.symbolLsnrs, lsnr)
}

func NewRepository() Repository {
	return &SimpleRepository{}
}
