package ensurable

type Ensurable interface {
	Prepare()
	Execute()
}
