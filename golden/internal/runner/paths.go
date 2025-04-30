package runner

import "path/filepath"

type RunnerPaths struct {
	root    string
	base    string
	test    string
	out     string
	scripts string
	scopes  string

	repoIn    string
	repoOut   string
	prepIn    string
	prepOut   string
	execIn    string
	execOut   string
	errorsIn  string
	errorsOut string
	errorFile string
}

func ConfigurePaths(root, test string) RunnerPaths {
	base := filepath.Join(root, test)
	errin := filepath.Join(base, "errors")
	errdir := filepath.Join(base, "errors-gen")
	errfile := filepath.Join(errdir, "errors.txt")
	outdir := filepath.Join(base, "out")
	repoin := filepath.Join(base, "repository")
	repoout := filepath.Join(base, "repository-gen")
	prepin := filepath.Join(base, "prepare")
	prepout := filepath.Join(base, "prepare-gen")
	execin := filepath.Join(base, "execute")
	execout := filepath.Join(base, "execute-gen")
	scripts := filepath.Join(base, "scripts")
	scopes := filepath.Join(base, "scopes")

	return RunnerPaths{
		root: root, base: base, out: outdir, test: test, scripts: scripts, scopes: scopes,
		errorsOut: errdir, errorsIn: errin, errorFile: errfile, repoIn: repoin, repoOut: repoout,
		prepIn: prepin, prepOut: prepout, execIn: execin, execOut: execout,
	}
}
