package utils

import "ziniki.org/deployer/driver/pkg/driverbottom"

func PropsMap(props map[driverbottom.Identifier]driverbottom.Expr) map[string]driverbottom.Identifier {
	ret := make(map[string]driverbottom.Identifier)
	for k := range props {
		ret[k.Id()] = k
	}
	return ret
}

func UseProps(props map[driverbottom.Identifier]driverbottom.Expr, notused map[string]driverbottom.Identifier, which ...string) map[driverbottom.Identifier]driverbottom.Expr {
	ret := make(map[driverbottom.Identifier]driverbottom.Expr)
	for _, s := range which {
		for k, v := range props {
			if k.Id() == s {
				ret[k] = v
				notused[s] = nil
				break
			}
		}
	}
	return ret
}

func FindProp(props map[driverbottom.Identifier]driverbottom.Expr, notused map[string]driverbottom.Identifier, which string) driverbottom.Expr {
	for k, v := range props {
		if k.Id() == which {
			notused[which] = nil
			return v
		}
	}
	panic("could not find " + which)
}

func HasProp(props map[driverbottom.Identifier]driverbottom.Expr, which ...string) bool {
	for k := range props {
		for _, s := range which {
			if k.Id() == s {
				return true
			}
		}
	}
	return false
}
