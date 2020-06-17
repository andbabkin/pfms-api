package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andbabkin/pfms-api/internal/controller"
)

// InternalServerError prepares http response with provided message and reports error in log
func InternalServerError(w http.ResponseWriter, err error, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, `{"error":"%s"}`, msg)
	log.Println(err)
}

// ControllerError prepares http response with error returned by controller
func ControllerError(w http.ResponseWriter, err error, s controller.ResponseStatus) {
	w.WriteHeader(int(s))
	fmt.Fprintf(w, `{"error":"%s"}`, err.Error())
}
