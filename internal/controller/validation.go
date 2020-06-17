package controller

import (
	"fmt"
	"strings"
)

// Validator validates request parameters and returns a list of validation errors
type Validator interface {
	Validate() []string
}

// Validate executes validator and returns an error which combines
// all found validation errors
func Validate(v Validator) error {
	errors := v.Validate()
	if len(errors) > 0 {
		msg := strings.Join(errors, "),(")
		msg = "(" + msg + ")"
		return fmt.Errorf(msg)
	}

	return nil
}
