package files

import (
	"fmt"
	"path/filepath"

	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/coremod/pkg/files"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/pluggable"
)

type dirAction struct {
	tools *external.Tools
	loc   *errorsink.Location
	exprs []pluggable.Expr
	res   *PathHolder
}

func (da *dirAction) Loc() *errorsink.Location {
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

func (da *dirAction) Resolve(r pluggable.Resolver) pluggable.BindingRequirement {
	// da.resolved = make([]pluggable.Expr, len(da.exprs))
	for _, e := range da.exprs {
		/*da.resolved[i] = */ e.Resolve(r)
	}
	da.res = &PathHolder{loc: da.loc}
	// b.MustBind(da.res)
	return pluggable.MUST_BE_BOUND
}

func (da *dirAction) BuildModel(pres pluggable.ValuePresenter) {
	if da.tools.Reporter.HasErrors() {
		return
	}
	var val *DirectoryPourer
	var err error
	for _, e := range da.exprs {
		v := da.tools.Storage.Eval(e)
		if val == nil {
			p, ok := v.(*DirectoryPourer)
			if ok {
				val = p
			} else {
				s, ok := v.(string)
				if ok {
					if filepath.IsAbs(s) {
						val, err = NewDirectoryPourer(s)
						if err != nil {
							panic(err)
						}
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
					val, err = val.Relative(s)
					if err != nil {
						panic(err)
					}
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

func (ea *dirAction) UpdateReality() {

}

func (ea *dirAction) TearDown() {

}

type PathHolder struct {
	loc  *errorsink.Location
	path files.FileSource
}

func (p *PathHolder) Loc() *errorsink.Location {
	return p.loc
}

func (p *PathHolder) ShortDescription() string {
	return fmt.Sprintf("PathHolder[%v]", p.path)
}

func (p *PathHolder) DumpTo(iw pluggable.IndentWriter) {
	iw.Intro("PathHolder")
	iw.AttrsWhere(p)
	// if p.path != nil {
	// 	iw.TextAttr("path", p.path)
	// }
	iw.EndAttrs()
}
