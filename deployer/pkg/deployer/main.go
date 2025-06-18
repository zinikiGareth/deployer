package deployer

type MainHandler interface {
	RunWithArgs(deployer Driver, args []string)
}
