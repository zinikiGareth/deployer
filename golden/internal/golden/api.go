package golden

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"ziniki.org/deployer/golden/internal/runner"
)

type GoldenRunner struct {
	modules  []string
	testdirs []string
}

func NewGoldenRunner(args []string) (*GoldenRunner, error) {
	ret := &GoldenRunner{}
	i := 0
	for i < len(args) {
		switch args[i] {
		case "-m":
			fallthrough
		case "--module":
			i++
			mod, err := nextArg(args, i, "there is no argument module")
			if err != nil {
				return nil, err
			}
			ret.UseModule(mod)
		default:
			ret.RunTestsUnder(args[i])
		}
		i++
	}

	return ret, nil
}

func nextArg(args []string, i int, err string) (string, error) {
	if i == len(args) {
		return "", fmt.Errorf("%v", err)
	}
	return args[i], nil
}

func (r *GoldenRunner) UseModule(path string) {
	r.modules = append(r.modules, path)
}

func (r *GoldenRunner) RunTestsUnder(root string) {
	r.testdirs = append(r.testdirs, root)
}

func (r *GoldenRunner) RunAll() {
	for _, p := range r.testdirs {
		r.runOne(p)
	}
}

func (r *GoldenRunner) runOne(root string) {
	merged := gatherTestsInOrder(root)
	for _, s := range merged {
		r.runCase(root, s)
	}
}

func (r *GoldenRunner) runCase(root, dir string) {
	run, err := runner.NewTestRunner(root, dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	run.Run(r.modules)
}

func gatherTestsInOrder(root string) []string {
	coll := make([]string, 0)
	coll = collectFiles(coll, root, "")
	saveTo := filepath.Join(root, "testorder")
	curr := loadTestOrder(saveTo)
	merged := mergeOrders(curr, coll)
	saveTestOrder(saveTo, merged)
	return merged
}

func collectFiles(coll []string, root string, sub string) []string {
	dir := filepath.Join(root, sub)
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("could not read directory %s: %v\n", dir, err)
		return coll
	}
	for _, f := range files {
		if f.IsDir() {
			if f.Name() == "scripts" {
				coll = append(coll, sub)
			} else {
				coll = collectFiles(coll, root, filepath.Join(sub, f.Name()))
			}
		}
	}
	return coll
}

func saveTestOrder(saveTo string, coll []string) {
	file, err := os.Create(saveTo)
	if err != nil {
		fmt.Printf("could not save to %s\n", saveTo)
		return
	}
	defer file.Close()

	for _, f := range coll {
		file.WriteString(f + "\n")
	}
	file.Sync()
}

func loadTestOrder(loadFrom string) []string {
	file, err := os.Open(loadFrom)
	if err != nil {
		fmt.Printf("could not read from %s\n", loadFrom)
		return make([]string, 0)
	}
	defer file.Close()

	ret := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}

func mergeOrders(curr, coll []string) []string {
	for _, n := range coll {
		if !slices.Contains(curr, n) {
			curr = append(curr, n)
		}
	}
	return curr
}
