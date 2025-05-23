package golden

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func gatherTestsInOrder(root string) []string {
	coll := make([]string, 0)
	coll = collectFiles(coll, root, "")
	saveTo := filepath.Join(root, "testorder")
	curr := loadTestOrder(root, saveTo)
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

func loadTestOrder(root, loadFrom string) []string {
	file, err := os.Open(loadFrom)
	if err != nil {
		fmt.Printf("could not read from %s\n", loadFrom)
		return make([]string, 0)
	}
	defer file.Close()

	ret := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f := scanner.Text()
		_, err := os.Stat(filepath.Join(root, f))
		if err != nil {
			fmt.Printf("cannot find test %s in %s, ignoring\n", f, loadFrom)
			continue
		}
		ret = append(ret, f)
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
