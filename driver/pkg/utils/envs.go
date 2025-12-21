package utils

import (
	"os"
	"strings"
	"syscall"
)

// Is it worth abstracting this return into an actual struct with "Set" and "Reset" methods?
func ReadEnvs(file string) (map[string]string, error) {
	lines, err := FileAsLines(file)

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

func SetEnvs(envs map[string]string) {
	for k, v := range envs {
		// log.Printf("setting %s to %s\n", k, v)
		os.Setenv(k, v)
	}
}

func UnsetEnvs(envs map[string]string) {
	for k := range envs {
		os.Setenv(k, "") // this is as close as go will let you get to unset, but it doesn't matter because get on an unset will return "" anyway
	}
}

func PruneLines(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		if strings.HasPrefix(l, "#") {
			continue
		}
		ret = append(ret, l)
	}
	return ret
}
