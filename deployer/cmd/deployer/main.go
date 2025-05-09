package main

import (
	"os"

	"ziniki.org/deployer/deployer/internal/impl"
)

func main() {
	if len(os.Args) == 1 {
		impl.Usage()
		os.Exit(1)
	}

	stat := impl.RunDeployer(os.Args[1:])
	os.Exit(stat)
}
