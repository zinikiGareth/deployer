package testDB

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type FieldExpr struct {
	loc   *errorsink.Location
	kind  string
	name  string
	ftype string
}

func (f *FieldExpr) Loc() *errorsink.Location {
	return f.loc
}

func (f *FieldExpr) ShortDescription() string {
	return fmt.Sprintf("Field[%s:%s %s]", f.kind, f.name, f.ftype)
}

func (f *FieldExpr) DumpTo(to driverbottom.IndentWriter) {
	to.Intro("FieldExpr")
	to.AttrsWhere(f)
	to.TextAttr("kind", f.kind)
	to.TextAttr("name", f.name)
	to.TextAttr("type", f.ftype)
	to.EndAttrs()
}

func (f *FieldExpr) Resolve(r driverbottom.Resolver) {
}

func (f *FieldExpr) Eval(s driverbottom.RuntimeStorage) any {
	// I think for all intents and purposes this is evaluated already
	return f
}

func (f *FieldExpr) String() string {
	return f.ShortDescription()
}

var _ driverbottom.Expr = &FieldExpr{}
