package blob

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type Blobber struct {
	driverbottom.Locatable
	expr driverbottom.Expr
}

func (a Blobber) Loc() *errorsink.Location {
	return a.Locatable.Loc()
}

func (a Blobber) ShortDescription() string {
	return fmt.Sprintf("aBlobber %s\n", a.expr.ShortDescription())
}

func (a Blobber) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("ABlobber")
	iw.AttrsWhere(a)
	iw.NestedAttr("expr", a.expr)
	iw.EndAttrs()
}

type BlobberNameMethod struct {
}

func (b *BlobberNameMethod) Invoke(storage driverbottom.RuntimeStorage, obj driverbottom.Expr, args []driverbottom.Expr) any {
	if len(args) > 0 {
		panic("should not have args")
	}
	asVar, ok := obj.(driverbottom.Var)
	if !ok {
		panic("not a var")
	}
	blobber, ok := storage.Get(asVar).(*Blobber)
	if !ok {
		panic("not a blobber")
	}
	return blobber.expr.Eval(storage)
}

func (a *Blobber) ObtainMethod(name string) driverbottom.Method {
	switch name {
	case "name":
		return &BlobberNameMethod{}
	default:
		panic(fmt.Sprintf("there is no method %s", name))
	}
}

type createBlobAction struct {
	driverbottom.Locatable
	expr    driverbottom.Expr
	blobber *Blobber
}

func (c *createBlobAction) Loc() *errorsink.Location {
	return c.Locatable.Loc()
}

func (c *createBlobAction) ShortDescription() string {
	return "create_blob[" + c.expr.ShortDescription() + "]"
}

func (c *createBlobAction) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("create_blob")
	iw.AttrsWhere(c)
	c.expr.DumpTo(iw)
	iw.EndAttrs()
}

func (c *createBlobAction) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	c.blobber = &Blobber{Locatable: c.Locatable, expr: c.expr}
	return driverbottom.MUST_BE_BOUND
}

func (c *createBlobAction) DetermineInitialState(pres driverbottom.ValuePresenter) {
	pres.NotFound()
}

func (c *createBlobAction) DetermineDesiredState(pres driverbottom.ValuePresenter) {
	pres.Present(c.blobber)
}

var _ corebottom.ModelBuilder = &createBlobAction{}
