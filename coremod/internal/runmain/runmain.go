package runmain

import (
	"fmt"
	"strings"

	"ziniki.org/deployer/coremod/internal/deployer"
	"ziniki.org/deployer/coremod/pkg/external"
	deployer2 "ziniki.org/deployer/deployer/pkg/deployer"
)

type mainHandler struct {
	tools *external.Tools
}

func (m *mainHandler) RunWithArgs(d deployer2.Driver, args []string) {
	options := m.tools.Options
	targets := []string{}

	failed := false

	for _, x := range args {
		switch x {
		case "--teardown":
			options.TearDown = true
		default:
			if strings.HasPrefix(x, "-") {
				fmt.Printf("unknown option: %s\n", x)
				failed = true
			}
			targets = append(targets, x)
		}
	}

	if failed {
		return
	}

	d2 := deployer.NewDeployer(m.tools)
	for _, s := range targets {
		err := d2.Deploy(s)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
	}
}

func MakeMainHandler(tools *external.Tools) deployer2.MainHandler {
	return &mainHandler{tools: tools}
}
