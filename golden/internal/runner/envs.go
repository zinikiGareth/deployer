package runner

import (
	"os"
	"strings"
	"syscall"

	"ziniki.org/deployer/deployer/pkg/utils"
)

func (r *TestRunner) ReadEnvs(file string) (map[string]string, error) {
	lines, err := utils.FileAsLines(file)

	if err != nil {
		pe, ok := err.(*os.PathError)
		if !ok {
			return nil, err
		}
		if pe.Op == "open" && pe.Err == syscall.ENOENT {
			return nil, nil
		}
		return nil, err
	}

	ret := make(map[string]string)
	lines = PruneLines(lines)
	for _, l := range lines {
		q := strings.Index(l, "=")
		if q == -1 {
			panic("env var did not have =: " + l)
		}
		ret[l[0:q]] = l[q+1:]
	}
	return ret, nil
}

func (r *TestRunner) SetEnvs(envs map[string]string) {
	for k, v := range envs {
		os.Setenv(k, v)
	}
}

func (r *TestRunner) UnsetEnvs(envs map[string]string) {
	for k := range envs {
		os.Setenv(k, "") // this is as close as go will let you get to unset, but it doesn't matter because get on an unset will return "" anyway
	}
}
