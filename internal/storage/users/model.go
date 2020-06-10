package users

import (
	"database/sql"
	"time"
)

const (
	// UserTable is a name of table in the database for User model
	UserTable = "users"
)

// User model maps all fields from a record in "users" table
type User struct {
	ID        uint64
	Name      string
	Pswd      string
	Token     sql.NullString
	Role      int8
	Email     sql.NullString
	Active    bool
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
