package deployer

type MainHandler interface {
	RunWithArgs(deployer Deployer, args []string)
}
