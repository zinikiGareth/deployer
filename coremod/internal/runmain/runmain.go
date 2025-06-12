package runmain

import (
	"fmt"
	"strings"

	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type mainHandler struct {
	tools *pluggable.Tools
}

func (m *mainHandler) RunWithArgs(d deployer.Deployer, args []string) {
	options := m.tools.Options
	targets := []string{}

	for _, x := range args {
		switch x {
		case "--teardown":
			options.TearDown = true
		default:
			if strings.HasPrefix(x, "-") {
				fmt.Printf("unknown option: %s\n", x)
				return
			}
			targets = append(targets, x)
		}
	}

	for _, s := range targets {
		err := d.Deploy(s)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
	}
}

func MakeMainHandler(tools *pluggable.Tools) deployer.MainHandler {
	return &mainHandler{tools: tools}
}
