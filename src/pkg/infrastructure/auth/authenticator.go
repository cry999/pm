package auth

import (
	"github.com/cry999/pm-projects/pkg/domain/model/iam"
	"golang.org/x/crypto/bcrypt"
)

type authenticator struct {
	cost int
}

// NewAuthenticator creates a new iam.Authenticator implements
func NewAuthenticator() iam.Authenticator {
	return &authenticator{
		cost: 10,
	}
}

func (a *authenticator) Hash(raw string) (hashed string, err error) {
	b, err := bcrypt.GenerateFromPassword([]byte(raw), a.cost)
	if err != nil {
		return
	}
	return string(b), nil
}

func (a *authenticator) Auth(raw string, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(raw))
}
