package golden

import (
	"fmt"
	"strings"

	"ziniki.org/deployer/golden/internal/errors"
)

func Usage() {
	fmt.Println("Usage: golden [-m|--module module] <test-dir-root> ...")
}

func NewGoldenRunner(args []string) (*GoldenRunner, error) {
	ret := &GoldenRunner{tracker: errors.NewCaseTracker()}
	i := 0
	for i < len(args) {
		switch args[i] {
		case "-m":
			fallthrough
		case "--module":
			i++
			mod, err := nextArg(args, i, "there is no argument module")
			if err != nil {
				return nil, err
			}
			ret.UseModule(mod)
		case "--pattern":
			i++
			patt, err := nextArg(args, i, "there is no argument pattern")
			if err != nil {
				return nil, err
			}
			ret.MatchPattern(patt)
		default:
			if strings.HasPrefix(args[i], "-") {
				return nil, fmt.Errorf("invalid option: %s", args[i])
			}
			ret.RunTestsUnder(args[i])
		}
		i++
	}

	return ret, nil
}

func nextArg(args []string, i int, err string) (string, error) {
	if i == len(args) {
		return "", fmt.Errorf("%v", err)
	}
	return args[i], nil
}
