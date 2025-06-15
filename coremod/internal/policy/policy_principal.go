package policy

type principalItem struct {
	key, value string
}

func (pi *principalItem) Key() string {
	return pi.key
}

func (pi *principalItem) Value() string {
	return pi.value
}

func NewPrincipal(key string, value string) *principalItem {
	return &principalItem{key: key, value: value}
}
