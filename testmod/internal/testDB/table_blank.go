package testDB

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type TableBlank struct{}

func (b *TableBlank) Find(ct *corebottom.Tools, loc *errorsink.Location, id corebottom.CoinId, named string) corebottom.FindCoin {
	// log.Printf("find coin id = %s\n", id)
	return &tableCreator{tools: ct, loc: loc, coin: id, name: named}
}

func (b *TableBlank) Mint(ct *corebottom.Tools, loc *errorsink.Location, id corebottom.CoinId, named string, props map[driverbottom.Identifier]driverbottom.Expr, teardown corebottom.TearDown) corebottom.Ensurable {
	// log.Printf("mint coin id = %s\n", id.VarName().Id())
	return &tableCreator{tools: ct, loc: loc, coin: id, name: named, props: props}
}

func (b *TableBlank) Loc() *errorsink.Location {
	panic("not implemented")
}

func (b *TableBlank) ShortDescription() string {
	return "test.S3.Bucket[]"
}

func (b *TableBlank) DumpTo(iw driverbottom.IndentWriter) {
	panic("not implemented")
}

var _ corebottom.Blank = &TableBlank{}
