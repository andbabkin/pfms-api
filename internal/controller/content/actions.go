package content

import (
	"github.com/andbabkin/pfms-api/internal/controller"
	"github.com/andbabkin/pfms-api/internal/domain/content"
)

// PageAction is responsible for textual content displayed on the page
func PageAction(r *PageRequest) (map[string]string, controller.ResponseStatus, error) {
	// validate request
	err := controller.Validate(r)
	if err != nil {
		return nil, controller.StatusBadRequest, err
	}

	// prepare content
	var c map[string]string
	s := controller.StatusOK
	c, err = content.GetPage(r.Page, r.User.Lang)
	if err != nil {
		s = controller.StatusNotFound
	}

	return c, s, err
}
