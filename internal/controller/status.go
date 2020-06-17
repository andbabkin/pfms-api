package controller

// ResponseStatus is a status of controller action response
type ResponseStatus int

// Available statuses for controller responses
const (
	StatusOK                  ResponseStatus = 200
	StatusBadRequest          ResponseStatus = 400
	StatusUnauthorized        ResponseStatus = 401
	StatusForbidden           ResponseStatus = 403
	StatusNotFound            ResponseStatus = 404
	StatusInternalServerError ResponseStatus = 500
)

var textMap = map[ResponseStatus]string{
	StatusOK:                  "OK",
	StatusBadRequest:          "Bad Request",
	StatusUnauthorized:        "Unauthorized",
	StatusForbidden:           "Forbidden",
	StatusNotFound:            "Not Found",
	StatusInternalServerError: "Internal Server Error",
}

// Text provides textual representaion of the status
func (s ResponseStatus) Text() string {
	return textMap[s]
}
