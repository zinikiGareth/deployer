package impl

import (
	"fmt"
	"os"

	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/errorsink"
)

func Usage() {
	fmt.Println("Usage: deployer [-m|--module module] <target> ...")
}

func RunDeployer(args []string) int {
	sink := errorsink.NewConsoleSink()
	d := NewDeployer(sink, os.Stdout)
	var others []string

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
		default:
			others = append(others, args[i])
		}
		i++
	}

	err := d.ReadScriptsFrom("trials")
	if err != nil {
		return 1
	}
	mainArgs := d.ObtainTools().Recall.Find("main-args", "main")
	runAs, ok := mainArgs.(deployer.MainHandler)
	if !ok {
		panic("main handler was not a MainHandler")
	}
	runAs.RunWithArgs(d, others)
	return 0
}

func nextArg(args []string, i int, err string) (string, error) {
	if i == len(args) {
		return "", fmt.Errorf("%v", err)
	}
	return args[i], nil
}
