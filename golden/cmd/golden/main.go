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


	harness, err := golden.NewGoldenRunner(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}
	harness.RunAll()
}
