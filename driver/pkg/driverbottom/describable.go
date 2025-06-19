package driverbottom

import "ziniki.org/deployer/driver/pkg/errorsink"

type Locatable interface {
	Loc() *errorsink.Location
}

type Describable interface {
	Locatable

	// ShortDescription enables clients to describe what they are pointing to in a unique way
	ShortDescription() string

	// DumpTo renders the whole of the text of the definition in a reproducible and unique, but not necessarily parseable form
	DumpTo(to IndentWriter)
}
