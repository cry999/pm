package commands

// SignupInput ...
type SignupInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignupOutput ...
type SignupOutput struct {
	UserDescriptor
}
