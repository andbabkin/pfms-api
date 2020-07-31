package users

import (
	"database/sql"
	"time"

	"github.com/andbabkin/pfms-api/internal/storage/base"
)

const (
	// UserTable is a name of table in the database for User model
	UserTable = "users"
)

// User model maps all fields from a record in "users" table
type User struct {
	base.AutoIncr
	Name   string
	Pswd   string
	Token  sql.NullString
	Role   int8
	Email  sql.NullString
	Active bool
	Lang   string
	base.Created
	base.Updated
	base.Deleted
}

// NewUser creates a new instance of active User
func NewUser(name, pswd string, role int8) *User {
	now := time.Now()
	return &User{
		Name:    name,
		Pswd:    pswd,
		Role:    role,
		Active:  true,
		Lang:    "en",
		Created: base.Created{CreatedAt: now},
		Updated: base.Updated{UpdatedAt: now},
	}
}
