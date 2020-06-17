package auth

import (
	"errors"

	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/andbabkin/pfms-api/internal/storage/users"
)

// ErrAuthenticationFailed notifies about stored and received passwords mismatch
var ErrAuthenticationFailed error = errors.New("Authentication failed")

// Authenticate finds an user by credentials and generates a token for him
func Authenticate(name, password string, ur users.IUserRepo) (*users.User, error) {
	// get user
	u, err := ur.FindLastByField("name", name)
	if err != nil {
		return nil, err
	}

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(u.Pswd), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return u, ErrAuthenticationFailed
	}
	if err != nil {
		return u, err
	}

	// generate and save token
	u.Token.String = ksuid.New().String()
	u.Token.Valid = true
	err = ur.UpdateField("token", u)

	return u, err
}
