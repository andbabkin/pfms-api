package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// ReadLimit - max allowed size of request body in bytes
const ReadLimit int64 = 1048576

// DecodeJSONBody fills up a struct or map variable pointed to by out
// with JSON request body data. Prepares error response if convertion failed.
// Returns TRUE if body is successfully converted
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, out interface{}) bool {
	// A request body larger than provided limit will result in
	// Decode() returning a "http: request body too large" error.
	r.Body = http.MaxBytesReader(w, r.Body, ReadLimit)

	err := json.NewDecoder(r.Body).Decode(out)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			http.Error(w, msg, http.StatusBadRequest)

		// Catch any type errors, like trying to assign a string in the
		// JSON request body to an int field in a struct.
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			http.Error(w, msg, http.StatusBadRequest)

		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: request body too large":
			msg := `Request body must not be larger than ` + string(ReadLimit) + ` bytes`
			http.Error(w, msg, http.StatusRequestEntityTooLarge)

		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			InternalServerError(w, err, "Failed to process JSON request body")
		}
		return false
	}

	return true
}

// SendJSONResponse prepares JSON response from provided data
func SendJSONResponse(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		InternalServerError(w, err, "Failed to create JSON response")
		return
	}
	w.Write(b)
}
