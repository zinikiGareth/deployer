package impl

import (
	"fmt"
	"os"
	"strings"

	"ziniki.org/deployer/deployer/pkg/errors"
)

func Usage() {
	fmt.Println("Usage: deployer [-m|--module module] <target> ...")
}

func RunDeployer(args []string) int {
	sink := errors.NewConsoleSink()
	deployer := NewDeployer(sink, os.Stdout)
	var targets []string

	i := 0
	for i < len(args) {
		switch args[i] {
		case "-m":
			fallthrough
		case "--module":
			i++
			mod, err := nextArg(args, i, "there is no argument module")
			if err != nil {
				return 1
			}
			err = deployer.UseModule(mod)
			if err != nil {
				fmt.Printf("Could not open module %s: %v\n", mod, err)
				return 1
			}
		// case "--pattern":
		// 	i++
		// 	patt, err := nextArg(args, i, "there is no argument pattern")
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	deployer.MatchPattern(patt)
		default:
			if strings.HasPrefix(args[i], "-") {
				return 1
			}
			targets = append(targets, args[i])
		}
		i++
	}

	err := deployer.ReadScriptsFrom("trials")
	if err != nil {
		return 1
	}
	for _, s := range targets {
		err = deployer.Deploy(s)
		if err != nil {
			return 1
		}
	}
	return 0
}

func nextArg(args []string, i int, err string) (string, error) {
	if i == len(args) {
		return "", fmt.Errorf("%v", err)
	}
	return args[i], nil
}
