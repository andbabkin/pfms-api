package http

import (
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/([a-zA-Z0-9]+)(?:/([a-zA-Z0-9]+))?/?$")

// Handle http server requests
func Handle(mux *http.ServeMux) {
	mux.HandleFunc("/auth/", makeHandler(AuthRouter))
	mux.HandleFunc("/app/", makeHandler(AppRouter))
}

// makeHandler creates a function which makes common adjustments and checks then forwards control to specific handler
func makeHandler(fn func(http.ResponseWriter, *http.Request, []string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		// enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token")
			return
		}

		// set default content type
		w.Header().Set("Content-Type", "application/json")

		// block unsupported methods
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", "POST")
			http.Error(w, `{"error":"Not supported method"}`, http.StatusMethodNotAllowed)
			return
		}

		// validate URL path
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.Error(w, `{"error":"Invalid path"}`, http.StatusNotFound)
			return
		}

		fn(w, r, m[1:])
	}
}
