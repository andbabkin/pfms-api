package content

import (
	"fmt"

	"github.com/andbabkin/pfms-api/internal/controller"
	"github.com/andbabkin/pfms-api/internal/domain/content"
)

// PageAction is responsible for textual content displayed on the page
func PageAction(r *PageRequest) (map[string]map[string]string, controller.ResponseStatus, error) {
	// validate request
	err := controller.Validate(r)
	if err != nil {
		return nil, controller.StatusBadRequest, err
	}

	// prepare content
	var c map[string]map[string]string
	s := controller.StatusOK
	l := r.User.Lang
	switch r.Page {
	case "home":
		c, err = content.GetHomePage(l)
	default:
		s = controller.StatusNotFound
		err = fmt.Errorf("No content for page (%s)", r.Page)
	}

	return c, s, err
}
