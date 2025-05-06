package files

import (
	"fmt"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type copyAction struct {
	tools    *pluggable.Tools
	loc      *errors.Location
	exprs    []pluggable.Expr
	assignTo pluggable.Identifier
}

func (ca *copyAction) Loc() *errors.Location {
	return ca.loc
}

func (ca *copyAction) Where() *errors.Location {
	return ca.loc
}

func (ca *copyAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("CopyCommand")
	w.AttrsWhere(ca)
	for _, v := range ca.exprs {
		w.IndPrintf("%s\n", v.String())
	}
	w.EndAttrs()
}

func (ca *copyAction) ShortDescription() string {
	return fmt.Sprintf("Dir[%d]", len(ca.exprs))
}

func (ca *copyAction) Completed() {
}

func (ca *copyAction) Resolve(r pluggable.Resolver) {
	// ea.resolved = r.Resolve(ea.what)
}

func (ca *copyAction) Prepare(runtime pluggable.RuntimeStorage) (pluggable.ExecuteAction, any) {
	// Not quite sure what to do here ...
	// Need to prepare
	// Should check things like permissions
	// Deffo need to return an ExecuteAction
	return nil, nil
}
