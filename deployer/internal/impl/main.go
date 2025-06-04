package impl

import (
	"fmt"
	"os"
	"strings"

	"ziniki.org/deployer/deployer/pkg/errorsink"
)

func Usage() {
	fmt.Println("Usage: deployer [-m|--module module] <target> ...")
}

func RunDeployer(args []string) int {
	sink := errorsink.NewConsoleSink()
	d := NewDeployer(sink, os.Stdout)
	var targets []string
	options := d.ObtainTools().Options

	i := 0
	for i < len(args) {
		switch args[i] {
		case "-m":
			fallthrough
		case "--module":
			i++
			mod, err := nextArg(args, i, "there is no argument module")
			if err != nil {
				fmt.Printf("%v\n", err)
				return 1
			}
			err = d.UseModule(mod)
			if err != nil {
				fmt.Printf("Could not open module %s: %v\n", mod, err)
				return 1
			}
		case "--teardown":
			options.TearDown = true
		default:
			if strings.HasPrefix(args[i], "-") {
				fmt.Printf("unknown option: %s\n", args[i])
				return 1
			}
			targets = append(targets, args[i])
		}
		i++
	}

	err := d.ReadScriptsFrom("trials")
	if err != nil {
		return 1
	}
	for _, s := range targets {
		err = d.Deploy(s)
		if err != nil {
			fmt.Printf("%v\n", err)
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
