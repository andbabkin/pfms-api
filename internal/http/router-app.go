package http

import (
	"net/http"

	"github.com/andbabkin/pfms-api/internal/domain/auth"
	"github.com/andbabkin/pfms-api/internal/storage/users"
)

// AppRouter works with app routes which require authorization
func AppRouter(w http.ResponseWriter, r *http.Request, path []string) {
	_, err := authorize(r)
	if err != nil {
		w.Header().Set("WWW-Authenticate", "Bearer realm=\"app\"")
		http.Error(w, `{"error":"Authorization failed"}`, http.StatusUnauthorized)
		return
	}

	switch path[1] {
	case "content":
		data := map[string]interface{}{
			"a1": 1,
			"a2": "hello",
			"a3": map[string]string{
				"b1": "bar",
				"b2": "foo",
			},
		}
		SendJSONResponse(w, data)
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
