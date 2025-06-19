package runmain

import (
	"fmt"
	"strings"

	"ziniki.org/deployer/coremod/internal/deployer"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type mainHandler struct {
	tools *corebottom.Tools
}

func (m *mainHandler) RunWithArgs(driver driverbottom.Driver, args []string) {
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

	depl := deployer.NewDeployer(driver, m.tools)
	for _, s := range targets {
		err := depl.Deploy(s)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
	}
}

func MakeMainHandler(tools *corebottom.Tools) driverbottom.MainHandler {
	return &mainHandler{tools: tools}
}
