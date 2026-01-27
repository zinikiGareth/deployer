package files

import (
	"fmt"
	"path/filepath"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/utils"
)

type ProjectDir struct {
	driverbottom.Locatable
	absprojdir string
}

func (t *ProjectDir) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	return driverbottom.MUST_BE_BOUND
}

func (t *ProjectDir) Eval(s driverbottom.RuntimeStorage) any {
	return t
}

func (t *ProjectDir) ShortDescription() string {
	return fmt.Sprintf("ProjectDir[%s]", t.absprojdir)
}

func (t *ProjectDir) DumpTo(iw driverbottom.IndentWriter) {
	iw.Intro("ProjectDir")
	iw.AttrsWhere(t)
	iw.TextAttr("absdir", t.absprojdir)
	iw.EndAttrs()
}

func (t ProjectDir) String() string {
	return t.absprojdir
}

type ProjectDirFunc struct {
	tools *corebottom.Tools
}

func (i *ProjectDirFunc) Fixity() driverbottom.Fixity {
	return driverbottom.OP_PREFIX
}

func (h *ProjectDirFunc) Associativity() bool {
	return false
}

func (h *ProjectDirFunc) Precedence() int {
	return 10 // it doesn't really matter, since we don't take any arguments
}

func (h *ProjectDirFunc) ReduceExpr(me driverbottom.Token, before []driverbottom.Expr, after []driverbottom.Expr) (driverbottom.Expr, bool) {
	rep := h.tools.Reporter
	if len(before) != 0 || len(after) != 0 {
		rep.Report(me.Loc().Offset, "files.project_dir")
		return nil, false
	}
	absdir, err := utils.MakeAbs(me.Loc().Line.File.File)
	if err != nil {
		rep.Reportf(me.Loc().Offset, "files.project_dir failed to get cwd: %v", err)
		return nil, false
	}
	return &ProjectDir{Locatable: me, absprojdir: filepath.Dir(absdir)}, true
}

func MakeProjectDirFunc(tools *corebottom.Tools) *ProjectDirFunc {
	return &ProjectDirFunc{tools: tools}
}

var _ driverbottom.Function = &ProjectDirFunc{}
