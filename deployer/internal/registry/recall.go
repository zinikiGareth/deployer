package registry

import "ziniki.org/deployer/deployer/pkg/pluggable"

type Recall interface {
	FindVerb(verb string) pluggable.Action
}
