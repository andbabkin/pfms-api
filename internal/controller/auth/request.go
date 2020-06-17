package auth

import "github.com/andbabkin/pfms-api/internal/storage/users"

// LoginRequest contains user's credentials for authentication
type LoginRequest struct {
	Name     string          `json:"name"`
	Password string          `json:"pswd"`
	UserRepo users.IUserRepo `json:"-"`
}

// Validate validates Name and Password fields
func (r *LoginRequest) Validate() []string {
	errors := []string{}
	l := len(r.Name)
	if l < 1 {
		errors = append(errors, "Name is required")
	}
	if l > 100 {
		errors = append(errors, "Name size should be less than or equal to 100")
	}
	l = len(r.Password)
	if l < 1 {
		errors = append(errors, "Password is required")
	}
	if l > 100 {
		errors = append(errors, "Password size should be less than or equal to 100")
	}

	return errors
}
