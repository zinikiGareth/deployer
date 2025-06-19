package impl

import (
	"fmt"
	"os"

	"ziniki.org/deployer/driver/pkg/deployer"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func Usage() {
	fmt.Println("Usage: deployer [-m|--module module] <target> ...")
}

func RunDeployer(args []string) int {
	sink := errorsink.NewConsoleSink()
	d := NewDriver(sink, os.Stdout)
	var dirs []string
	var envs []string
	var modules []string
	var others []string

	i := 0
	for i < len(args) {
		switch args[i] {
		case "-d":
			i++
			dir, err := nextArg(args, i, "there is no argument dir")
			if err != nil {
				fmt.Printf("%v\n", err)
				return 1
			}
			dirs = append(dirs, dir)
		case "-e":
			i++
			env, err := nextArg(args, i, "there is no argument env")
			if err != nil {
				fmt.Printf("%v\n", err)
				return 1
			}
			envs = append(envs, env)
		case "-m":
			fallthrough
		case "--module":
			i++
			mod, err := nextArg(args, i, "there is no argument module")
			if err != nil {
				fmt.Printf("%v\n", err)
				return 1
			}
			modules = append(modules, mod)
		default:
			others = append(others, args[i])
		}
		i++
	}

	if len(dirs) == 0 {
		fmt.Printf("no script directories specified")
		return 1
	}
	for _, f := range dirs {
		err := d.ReadScriptsFrom(f)
		if err != nil {
			return 1
		}
	}
	for _, e := range envs {
		want := e + ".envs"
		if !d.FindAndReadEnvs(dirs, want) {
			fmt.Printf("did not find any files called %s in the specified directories\n", want)
			return 1
		}
	}

	for _, mod := range modules {
		err := d.UseModule(mod)
		if err != nil {
			fmt.Printf("Could not open module %s: %v\n", mod, err)
			return 1
		}
	}

	mainArgs := d.ObtainCoreTools().Recall.Find("main-args", "main")
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
