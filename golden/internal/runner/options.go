package runner

import (
	"ziniki.org/deployer/deployer/pkg/utils"
)

func (r *TestRunner) ReadTeardown(file string) bool {
	_, err := utils.FileAsLines(file)

	if err != nil {
		return false
	}
	return true
}

func (r *TestRunner) SetTearDown(b bool) {
	r.deployer.ObtainTools().Options.TearDown = b
}
