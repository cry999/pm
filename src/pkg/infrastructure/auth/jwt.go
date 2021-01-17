package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Tokenizer ...
type Tokenizer struct {
	private *rsa.PrivateKey
	public  *rsa.PublicKey
	issuer  string
}

// NewTokenizer creates a new Tokenizer instance
func NewTokenizer(prvkeyFile, pubkeyFile string) (_ *Tokenizer, err error) {
	rawPrivate, err := ioutil.ReadFile(prvkeyFile)
	if err != nil {
		return
	}
	private, err := jwt.ParseRSAPrivateKeyFromPEM(rawPrivate)
	if err != nil {
		return
	}
	rawPublic, err := ioutil.ReadFile(pubkeyFile)
	if err != nil {
		return
	}
	public, err := jwt.ParseRSAPublicKeyFromPEM(rawPublic)
	if err != nil {
		return
	}
	return &Tokenizer{
		private: private,
		public:  public,
		issuer:  "https://iam.pm-projects.com/",
	}, nil
}

// Tokenize ...
func (t *Tokenizer) Tokenize(userID string) (_ string, err error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		// standard claims
		"aud": "https://api.pm-projects.com/", // TODO: 外部から指定できるようにする
		"eat": now.Add(24 * time.Hour).Unix(),
		"jti": uuid.New().String(),
		"iat": now.Unix(),
		"iss": t.issuer,
		"nbf": now.Unix(),
		"sub": "AccessToken",
		// private claims
		t.privateKey("user-id"): userID,
	})
	return token.SignedString(t.private)
}

// Restore ...
func (t *Tokenizer) Restore(token string) (userID string, err error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return t.public, nil
	})
	if err != nil {
		return
	}
	if !parsedToken.Valid {
		if err = parsedToken.Claims.Valid(); err != nil {
			return
		}
		err = fmt.Errorf("invalid token")
		return
	}
	claim, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("invalid claims")
		return
	}
	maybeUserID, ok := claim[t.privateKey("user-id")]
	if !ok {
		err = fmt.Errorf("user-id not found in claims")
		return
	}
	userID, ok = maybeUserID.(string)
	if !ok {
		err = fmt.Errorf("userID is not string: %v", maybeUserID)
		return
	}
	return
}

func (t *Tokenizer) privateKey(key string) string {
	return path.Join(t.issuer, key)
}
