package http

import (
	"fmt"
	"log"
	"net/http"
)

// InternalServerError prepares http response with provided message and reports error in log
func InternalServerError(w http.ResponseWriter, err error, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, `{"error":"%s"}`, msg)
	log.Println(err)
}
