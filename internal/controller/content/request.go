package content

import "github.com/andbabkin/pfms-api/internal/storage/users"

// PageRequest contains page name for requested content
type PageRequest struct {
	Page string      `json:"page"`
	User *users.User `json:"-"`
}

// Validate validates user defined data in PageRequest
func (r *PageRequest) Validate() []string {
	if len(r.Page) < 1 {
		return []string{"Page is required"}
	}

	return []string{}
}
