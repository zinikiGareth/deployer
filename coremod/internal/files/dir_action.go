package files

import (
	"fmt"
	"log"
	"path/filepath"

	"ziniki.org/deployer/coremod/pkg/files"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type dirAction struct {
	tools    *pluggable.Tools
	loc      *errors.Location
	exprs    []pluggable.Expr
	assignTo pluggable.Identifier
}

func (da *dirAction) Loc() *errors.Location {
	return da.loc
}

func (da *dirAction) Where() *errors.Location {
	return da.loc
}

func (da *dirAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("DirAction")
	w.AttrsWhere(da)
	for _, v := range da.exprs {
		w.IndPrintf("%s\n", v.String())
	}
	w.EndAttrs()
}

func (da *dirAction) ShortDescription() string {
	return fmt.Sprintf("Dir[%d]", len(da.exprs))
}

func (da *dirAction) Completed() {
}

func (da *dirAction) Resolve(r pluggable.Resolver) {
	// ea.resolved = r.Resolve(ea.what)
}

func (da *dirAction) Prepare(runtime pluggable.RuntimeStorage) pluggable.ExecuteAction {
	var val *files.Path
	for _, e := range da.exprs {
		v := runtime.Eval(e)
		if val == nil {
			p, ok := v.(files.Path)
			if ok {
				val = &p
			} else {
				s, ok := v.(string)
				if ok {
					if filepath.IsAbs(s) {
						val = &files.Path{File: s}
					} else {
						log.Fatalf("cannot use non-abs path here: %v\n", v)
					}
				} else {
					log.Fatalf("cannot handle base path %v\n", v)
				}
			}
		} else {
			s, ok := v.(string)
			if ok {
				if !filepath.IsAbs(s) {
					val = &files.Path{File: filepath.Join(val.File, s)}
				} else {
					log.Fatalf("cannot use abs path here: %v\n", v)
				}
			} else {
				log.Fatalf("cannot handle nested path %v\n", v)
			}
		}
	}
	runtime.Bind(pluggable.SymbolName(da.assignTo.Id()), val)
	return nil
}
