package runner

import (
	"path/filepath"

	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/utils"
	"ziniki.org/deployer/golden/internal/errors"
)

type goldenComparator struct {
	tracker *errors.CaseTracker
	RunnerPaths
}

func newGoldenComparator(tracker *errors.CaseTracker, paths RunnerPaths) *goldenComparator {
	return &goldenComparator{tracker: tracker, RunnerPaths: paths}
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
	goldenFiles, genFiles, err := gc.prepare(eh, golden, gen)
	if err != nil {
		return
	}

	// record the generated files
	genmap := gc.catalogGenned(genFiles)
	// Go through the golden files, comparing to the generated ones
	gc.traverseGolden(eh, genmap, golden, gen, goldenFiles)
	// If there are any generated files which don't have golden files, let the user know and copy them
	gc.copyGenned(eh, genmap, golden, gen, copyNewFiles)
}

func (gc *goldenComparator) prepare(eh deployer.ErrorHandler, golden, gen string) ([]string, []string, error) {
	eh.Writef("comparing %s to %s\n", golden, gen)
	goldenFiles, err1 := utils.FindFiles(golden, "")
	genFiles, err2 := utils.FindFiles(gen, "")
	if err1 != nil && err2 != nil {
		// it's "safe" to assume it's an empty case
		return goldenFiles, genFiles, err1
	}
	if err1 != nil {
		// OK, it should be there.  Create it and we'll copy the files in
		utils.EnsureDir(golden)
	}
	if err2 != nil {
		// Presumably if there is a golden dir, there should be a gen dir
		eh.Writef("error collecting generated files from %s\n", gen)
		eh.Fail()
		return goldenFiles, genFiles, err2
	}
	return goldenFiles, genFiles, nil
}

func (gc *goldenComparator) catalogGenned(genFiles []string) map[string]int {
	genmap := make(map[string]int)
	for k, g := range genFiles {
		genmap[g] = k + 1
	}
	return genmap
}

func (gc *goldenComparator) traverseGolden(eh deployer.ErrorHandler, genmap map[string]int, golden, gen string, goldenFiles []string) {
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
}

func (gc *goldenComparator) copyGenned(eh deployer.ErrorHandler, genmap map[string]int, golden, gen string, copyNewFiles bool) {
	if len(genmap) > 0 {
		if copyNewFiles {
			eh.Writef("generated files in %s were not present in %s ... copying:\n", gen, golden)
			for f := range genmap {
				eh.Writef("  %s\n", f)
				utils.CopyFile(filepath.Join(gen, f), filepath.Join(golden, f))
			}
		} else {
			for f := range genmap {
				eh.Writef("there is no golden file for generated %s\n", f)
				eh.Fail()
			}
		}
	}
}
