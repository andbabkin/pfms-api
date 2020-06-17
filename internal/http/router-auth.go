package http

import (
	"net/http"

	"github.com/andbabkin/pfms-api/internal/controller/auth"
	"github.com/andbabkin/pfms-api/internal/storage/users"
)

// AuthRouter works with authentication/authorization routes (login, password reset, etc)
func AuthRouter(w http.ResponseWriter, r *http.Request, path []string) {
	switch path[1] {
	case "login":
		lr := &auth.LoginRequest{}
		if DecodeJSONBody(w, r, lr) {
			lr.UserRepo = &users.UserRepo{}
			data, s, err := auth.LoginAction(lr)
			if err != nil {
				ControllerError(w, err, s)
				return
			}
			SendJSONResponse(w, data)
		}
	default:
		http.Error(w, `{"error":"Route does not exist"}`, http.StatusNotFound)
		return
	}
}
