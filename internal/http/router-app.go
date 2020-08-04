package http

import (
	"log"
	"net/http"

	"github.com/andbabkin/pfms-api/internal/controller"
	"github.com/andbabkin/pfms-api/internal/controller/content"
	"github.com/andbabkin/pfms-api/internal/domain/auth"
	"github.com/andbabkin/pfms-api/internal/storage/users"
)

// AppRouter works with app routes which require authorization
func AppRouter(w http.ResponseWriter, r *http.Request, path []string) {
	// Authorize
	u, err := authorize(r)
	if err != nil {
		log.Println(err.Error())
		w.Header().Set("WWW-Authenticate", "Bearer realm=\"app\"")
		http.Error(w, `{"error":"Authorization failed"}`, http.StatusUnauthorized)
		return
	}

	// Execute controller action
	switch path[1] {
	case "content":
		processContent(w, r, u)
	default:
		http.Error(w, `{"error":"App route does not exist"}`, http.StatusNotFound)
		return
	}
}

func authorize(r *http.Request) (*users.User, error) {
	bearer := r.Header.Get("Authorization")
	if len(bearer) < 8 || bearer[0:7] != "Bearer " {
		return nil, auth.ErrAccessDenied
	}

	token := bearer[7:]
	user, err := auth.Authorize(token, "app", &users.UserRepo{})
	if err != nil {
		return nil, auth.ErrAccessDenied
	}

	return user, nil
}

func processContent(w http.ResponseWriter, r *http.Request, u *users.User) {
	pageRequest := &content.PageRequest{}
	if DecodeJSONBody(w, r, pageRequest) {
		pageRequest.User = u
		data, s, err := content.PageAction(pageRequest)
		sendResponse(w, data, s, err)
	}
}

func sendResponse(w http.ResponseWriter, data interface{}, s controller.ResponseStatus, err error) {
	if err != nil {
		ControllerError(w, err, s)
		return
	}
	SendJSONResponse(w, data)
}
