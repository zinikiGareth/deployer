package deployer

type userError struct {
	msg string
}

func (u *userError) Error() string {
	return u.msg
}

func UserError(msg string) error {
	return &userError{msg: msg}
}
