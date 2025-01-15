package command

type LoginWithSMSCommand struct {
	PhoneNumber string
	Code        string
}

type LoginWithSMSCommandResult struct {
	AccessToken string
}
