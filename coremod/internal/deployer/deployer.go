package deployer

import (
	"fmt"
	"io"

	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type DeployerImpl struct {
	driver       *impl.DriverImpl
	tools        *external.Tools
	userErrorsTo io.StringWriter
	srcdir       string
	input        []string
}

func (d *DeployerImpl) Deploy(targetNames ...string) error {
	err := driver.DoStuff()
	if err != nil {
		return err
	}

	targets, err := d.findTargets(targetNames...)
	if err != nil {
		return err
	}

	d.tools.Storage.SetMode(pluggable.BUILD_MODEL_MODE)
	for _, t := range targets {
		// fmt.Printf("preparing %s:\n", t.String())
		t.BuildModel()
	}

	if d.tools.Reporter.HasErrors() {
		return fmt.Errorf("errors building model ... NOT UPDATING REALITY")
	}

	if d.tools.Options.TearDown {
		d.tools.Storage.SetMode(pluggable.UPDATE_REALITY_MODE)
		for _, t := range targets {
			// fmt.Printf("tearing down %s:\n", t)
			t.TearDown()
		}
	} else {
		d.tools.Storage.SetMode(pluggable.UPDATE_REALITY_MODE)
		for _, t := range targets {
			// fmt.Printf("executing %s:\n", t)
			t.UpdateReality()
		}
	}

	return nil
}

func (d *DeployerImpl) findTargets(names ...string) ([]pluggable.TargetThing, error) {
	var targets []pluggable.TargetThing
	var ue error
	for _, n := range names {
		t := d.tools.Repository.FindTop(pluggable.SymbolName(n))
		var target pluggable.TargetThing
		var ok bool
		if t != nil {
			target, ok = t.(pluggable.TargetThing)
		}
		if t == nil || !ok {
			msg := fmt.Sprintf("there is no target %s\n", n)
			d.userErrorsTo.WriteString(msg)
			if ue == nil {
				ue = deployer.UserError(msg)
			}
		}
		targets = append(targets, target)
	}
	if ue != nil {
		return nil, ue
	}
	return targets, nil
}
