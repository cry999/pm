package iam

// Authenticator ...
type Authenticator interface {
	Hash(raw string) (hashed string, err error)
	Auth(raw, hashed string) error
}
