package runner

import (
	"path/filepath"

	"ziniki.org/deployer/deployer/pkg/utils"
	"ziniki.org/deployer/golden/internal/errors"
)

type goldenComparator struct {
	tracker   *errors.CaseTracker
	errorsIn  string
	errorsOut string
	repoIn    string
	repoOut   string
	prepIn    string
	prepOut   string
	execIn    string
	execOut   string
}

func newGoldenComparator(tracker *errors.CaseTracker, errorsIn, errorsOut, repoIn, repoOut, prepIn, prepOut, execIn, execOut string) *goldenComparator {
	return &goldenComparator{tracker: tracker, errorsIn: errorsIn, errorsOut: errorsOut, repoIn: repoIn, repoOut: repoOut, prepIn: prepIn, prepOut: prepOut, execIn: execIn, execOut: execOut}
}

func (gc *goldenComparator) compareAll() {
	gc.compareDirectory(gc.errorsIn, gc.errorsOut, false)
	gc.compareDirectory(gc.repoIn, gc.repoOut, true)
	gc.compareDirectory(gc.prepIn, gc.prepOut, true)
	gc.compareDirectory(gc.execIn, gc.execOut, true)
}

func (gc *goldenComparator) compareDirectory(golden, gen string, copyNewFiles bool) {
	base := filepath.Base(golden)
	eh := gc.tracker.ErrorHandlerFor("golden-" + base)
	eh.Writef("comparing %s to %s\n", golden, gen)
	goldenFiles, err1 := utils.FindFiles(golden, "")
	genFiles, err2 := utils.FindFiles(gen, "")
	if err1 != nil && err2 != nil {
		// it's "safe" to assume it's an empty case
		return
	}
	if err1 != nil {
		// OK, it should be there.  Create it and we'll copy the files in
		utils.EnsureDir(golden)
	}
	if err2 != nil {
		// Presumably if there is a golden dir, there should be a gen dir
		eh.Writef("error collecting generated files from %s\n", gen)
		return
	}

	// Go through the golden files, comparing to the generated ones
	genmap := make(map[string]int)
	for k, g := range genFiles {
		genmap[g] = k + 1
	}
	failed := false
	for _, f := range goldenFiles {
		if genmap[f] != 0 {
			if !utils.CompareFiles(filepath.Join(gen, f), filepath.Join(golden, f)) {
				eh.Writef("generated file %s did not match golden file\n", f)
				eh.Fail()
			}
			delete(genmap, f)
		} else { // if there is no generated file, complain: that's a failure
			eh.Writef("there is no gen file for %s\n", f)
			if !failed {
				eh.Fail()
				failed = true
			}
		}
	}

	// If there are any generated files which don't have golden files, let the user know and copy them
	if len(genmap) > 0 {
		if copyNewFiles {
			eh.Writef("generated files in %s were not present in %s ... copying:\n", gen, golden)
			for f := range genmap {
				eh.Writef("  %s\n", f)
				utils.CopyFile(filepath.Join(gen, f), filepath.Join(golden, f))
			}
		} else {
			for _, f := range genFiles {
				eh.Writef("there is no golden file for generated %s\n", f)
				eh.Fail()
			}
		}
	}
}
