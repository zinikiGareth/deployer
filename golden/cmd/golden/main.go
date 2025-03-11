package main

import (
	"fmt"
	"os"

	"ziniki.org/deployer/golden/internal/golden"
)

func main() {
	if len(os.Args) == 1 {
		golden.Usage()
		return
	}

	harness, err := golden.NewGoldenRunner(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}
	harness.RunAll()
	errs := harness.Report()
	os.Exit(errs)
}
