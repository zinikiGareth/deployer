package runner

import (
	"ziniki.org/deployer/driver/pkg/utils"
)

func (r *TestRunner) FileExists(file string) bool {
	_, err := utils.FileAsLines(file)

	return err == nil
}

func (r *TestRunner) SetTearDown(b bool) {
	r.deployer.ObtainTools().Options.TearDown = b
}

func (r *TestRunner) SetDestroy(b bool) {
	tools := r.deployer.ObtainTools()
	tools.Options.Destroy = b
}
