package deployer

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/external"
	"ziniki.org/deployer/driver/pkg/deployer"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type DeployerImpl struct {
	driver deployer.Driver
	tools  *external.Tools
}

func (d *DeployerImpl) ObtainTools() *external.Tools {
	return d.tools
}

func (d *DeployerImpl) Deploy(targetNames ...string) error {
	err := d.driver.DoStuff()
	if err != nil {
		return err
	}

	targets, err := d.findTargets(targetNames...)
	if err != nil {
		return err
	}

	d.tools.Storage.SetMode(driverbottom.BUILD_MODEL_MODE)
	for _, t := range targets {
		// fmt.Printf("preparing %s:\n", t.String())
		t.BuildModel()
	}

	if d.tools.Reporter.HasErrors() {
		return fmt.Errorf("errors building model ... NOT UPDATING REALITY")
	}

	if d.tools.Options.TearDown {
		d.tools.Storage.SetMode(driverbottom.UPDATE_REALITY_MODE)
		for _, t := range targets {
			// fmt.Printf("tearing down %s:\n", t)
			t.TearDown()
		}
	} else {
		d.tools.Storage.SetMode(driverbottom.UPDATE_REALITY_MODE)
		for _, t := range targets {
			// fmt.Printf("executing %s:\n", t)
			t.UpdateReality()
		}
	}

	return nil
}

func (d *DeployerImpl) findTargets(names ...string) ([]driverbottom.TargetThing, error) {
	var targets []driverbottom.TargetThing
	var ue error
	for _, n := range names {
		t := d.tools.Repository.FindTop(driverbottom.SymbolName(n))
		var target driverbottom.TargetThing
		var ok bool
		if t != nil {
			target, ok = t.(driverbottom.TargetThing)
		}
		if t == nil || !ok {
			msg := fmt.Sprintf("there is no target %s\n", n)
			d.driver.UserError(msg)
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

func NewDeployer(driver deployer.Driver, tools *external.Tools) external.Deployer {
	return &DeployerImpl{driver: driver, tools: tools}
}
