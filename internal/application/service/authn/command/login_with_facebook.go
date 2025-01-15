package command

type LoginWithFacebookCommand struct {
	Token string
}

type LoginWithFacebookCommandResult struct {
	AccessToken string
}
