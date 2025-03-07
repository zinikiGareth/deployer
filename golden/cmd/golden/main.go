package main

import (
	"fmt"
	"os"

	"ziniki.org/deployer/golden/internal/golden"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: golden <test-dir-root> ...")
		return
	}

	for _, d := range os.Args[1:] {
		golden.RunTestsUnder(d)
	}
}
