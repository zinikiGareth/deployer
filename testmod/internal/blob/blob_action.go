package blob

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Blobber struct {
	pluggable.Locatable
	expr pluggable.Expr
}

func (a Blobber) Loc() *errorsink.Location {
	return a.Locatable.Loc()
}

func (a Blobber) ShortDescription() string {
	return fmt.Sprintf("aBlobber %s\n", a.expr.ShortDescription())
}

func (a Blobber) DumpTo(iw pluggable.IndentWriter) {
	iw.Intro("ABlobber")
	iw.AttrsWhere(a)
	iw.NestedAttr("expr", a.expr)
	iw.EndAttrs()
}

type BlobberNameMethod struct {
}

func (b *BlobberNameMethod) Invoke(storage pluggable.RuntimeStorage, obj pluggable.Expr, args []pluggable.Expr) any {
	if len(args) > 0 {
		panic("should not have args")
	}
	asVar, ok := obj.(pluggable.Var)
	if !ok {
		panic("not a var")
	}
	blobber, ok := storage.Get(asVar).(*Blobber)
	if !ok {
		panic("not a blobber")
	}
	return blobber.expr.Eval(storage)
}

func (a *Blobber) ObtainMethod(name string) pluggable.Method {
	switch name {
	case "name":
		return &BlobberNameMethod{}
	default:
		panic(fmt.Sprintf("there is no method %s", name))
	}
}

type createBlobAction struct {
	pluggable.Locatable
	expr    pluggable.Expr
	blobber *Blobber
}

func (c *createBlobAction) Loc() *errorsink.Location {
	return c.Locatable.Loc()
}

func (c *createBlobAction) ShortDescription() string {
	return "create_blob[" + c.expr.ShortDescription() + "]"
}

func (c *createBlobAction) DumpTo(iw pluggable.IndentWriter) {
	iw.Intro("create_blob")
	iw.AttrsWhere(c)
	c.expr.DumpTo(iw)
	iw.EndAttrs()
}

func (c *createBlobAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	c.blobber = &Blobber{Locatable: c.Locatable, expr: c.expr}
	// b.MustBind(c.blobber)
	return pluggable.MUST_BE_BOUND
}

func (c *createBlobAction) Prepare(pres pluggable.ValuePresenter) {
	pres.Present(c.blobber)
}
