package files

import (
	"fmt"
	"path/filepath"

	"ziniki.org/deployer/coremod/pkg/files"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type dirAction struct {
	tools *pluggable.Tools
	loc   *errors.Location
	exprs []pluggable.Expr
	res   *PathHolder
}

func (da *dirAction) Loc() *errors.Location {
	return da.loc
}

func (da *dirAction) DumpTo(w pluggable.IndentWriter) {
	w.Intro("DirAction")
	w.AttrsWhere(da)
	for _, v := range da.exprs {
		w.IndPrintf("%s\n", v.ShortDescription())
	}
	w.EndAttrs()
}

func (da *dirAction) ShortDescription() string {
	return fmt.Sprintf("Dir[%d]", len(da.exprs))
}

func (da *dirAction) Completed() {
}

func (da *dirAction) Resolve(r pluggable.Resolver, b pluggable.Binder) {
	// da.resolved = make([]pluggable.Expr, len(da.exprs))
	for _, e := range da.exprs {
		/*da.resolved[i] = */ e.Resolve(r)
	}
	da.res = &PathHolder{loc: da.loc}
	b.MustBind(da.res)
}

func (da *dirAction) Prepare(pres pluggable.ValuePresenter) {
	var val *files.Path
	for _, e := range da.exprs {
		v := da.tools.Storage.Eval(e)
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
						panic(fmt.Sprintf("cannot use non-abs path here: %v\n", v))
					}
				} else {
					panic(fmt.Sprintf("cannot handle base path %T\n", v))
				}
			}
		} else {
			s, ok := v.(string)
			if ok {
				if !filepath.IsAbs(s) {
					val = &files.Path{File: filepath.Join(val.File, s)}
				} else {
					panic(fmt.Sprintf("cannot use abs path here: %v\n", v))
				}
			} else {
				panic(fmt.Sprintf("cannot handle nested path %T\n", v))
			}
		}
	}
	pres.Present(val)
}

func (ea *dirAction) Execute() {

}

type PathHolder struct {
	loc  *errors.Location
	path *files.Path
}

func (p *PathHolder) Loc() *errors.Location {
	return p.loc
}

func (p *PathHolder) ShortDescription() string {
	dir := "<nil>"
	if p.path != nil {
		dir = p.path.File
	}
	return fmt.Sprintf("PathHolder[%s]", dir)
}

func (p *PathHolder) DumpTo(iw pluggable.IndentWriter) {
	iw.Intro("PathHolder")
	iw.AttrsWhere(p)
	if p.path != nil {
		iw.TextAttr("path", p.path.File)
	}
	iw.EndAttrs()
}
