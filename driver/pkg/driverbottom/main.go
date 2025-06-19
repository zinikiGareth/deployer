package driverbottom

type MainHandler interface {
	RunWithArgs(deployer Driver, args []string)
}
