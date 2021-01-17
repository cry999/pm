package commands

// SigninInput ...
type SigninInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SigninOutput ...
type SigninOutput struct {
	UserDescriptor
}
