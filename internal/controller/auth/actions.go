package auth

import (
	"errors"
	"log"

	"github.com/andbabkin/pfms-api/internal/controller"
	"github.com/andbabkin/pfms-api/internal/domain/auth"
	"github.com/andbabkin/pfms-api/internal/storage/users"
)

// LoginAction processes a login request and returns
// a token in response if user is successfully authenticated
func LoginAction(r *LoginRequest) (*LoginResponse, controller.ResponseStatus, error) {
	// validate request
	err := controller.Validate(r)
	if err != nil {
		return nil, controller.StatusBadRequest, err
	}

	// authenticate
	user, err := auth.Authenticate(r.Name, r.Password, r.UserRepo)
	if errors.Is(err, users.ErrUserNotFound) || errors.Is(err, auth.ErrAuthenticationFailed) {
		return nil, controller.StatusUnauthorized, errors.New(controller.StatusUnauthorized.Text())
	}
	if err != nil {
		log.Println(err)
		return nil, controller.StatusInternalServerError, errors.New(controller.StatusInternalServerError.Text())
	}

	return &LoginResponse{Token: user.Token.String}, controller.StatusOK, nil
}
