package deployer

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type DeployerImpl struct {
	driver driverbottom.Driver
	tools  *corebottom.Tools
}

func (d *DeployerImpl) ObtainTools() *corebottom.Tools {
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

	d.tools.Storage.SetMode(corebottom.DETERMINE_INITIAL_MODE)
	for _, t := range targets {
		t.DetermineInitialState()
	}

	d.tools.Storage.SetMode(corebottom.DETERMINE_DESIRED_MODE)
	for _, t := range targets {
		t.DetermineDesiredState()
	}

	if d.tools.Reporter.HasErrors() {
		return fmt.Errorf("errors building model ... NOT UPDATING REALITY")
	}

	if d.tools.Options.TearDown {
		d.tools.Storage.SetMode(corebottom.TEARDOWN_MODE)
		for _, t := range targets {
			// fmt.Printf("tearing down %s:\n", t)
			t.TearDown()
		}
	} else {
		d.tools.Storage.SetMode(corebottom.UPDATE_REALITY_MODE)
		for _, t := range targets {
			// fmt.Printf("executing %s:\n", t)
			t.UpdateReality()
		}
	}

	return nil
}

func (d *DeployerImpl) findTargets(names ...string) ([]corebottom.Target, error) {
	var targets []corebottom.Target
	var ue error
	for _, n := range names {
		t := d.tools.Repository.FindTop(driverbottom.SymbolName(n))
		var target corebottom.Target
		var ok bool
		if t != nil {
			target, ok = t.(corebottom.Target)
		}
		if t == nil || !ok {
			msg := fmt.Sprintf("there is no target %s\n", n)
			d.driver.UserError(msg)
			if ue == nil {
				ue = driverbottom.UserError(msg)
			}
		}
		targets = append(targets, target)
	}
	if ue != nil {
		return nil, ue
	}
	return targets, nil
}

func NewDeployer(driver driverbottom.Driver, tools *corebottom.Tools) corebottom.Deployer {
	return &DeployerImpl{driver: driver, tools: tools}
}
