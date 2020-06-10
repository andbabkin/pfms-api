package users

import (
	"time"

	"github.com/andbabkin/pfms-api/internal/storage"
)

// IUserRepo is an interface to users data repository
type IUserRepo interface {
	Create(*User) (*User, error)
}

// UserRepo is an implementation of IUserRepo which connects to database
type UserRepo struct{}

// Create inserts a new record into users table and returns its ID if created
func (r *UserRepo) Create(name, pswd, email string, role int8, active bool) (int64, error) {
	sqlstr := `INSERT INTO ` + UserTable + ` (name, pswd, role, email, active, created_at, updated_at)
	VALUES (:name, :pswd, :role, :email, :active, :created_at, :updated_at)`

	now := time.Now()
	m := map[string]interface{}{
		"name":       name,
		"pswd":       pswd,
		"role":       role,
		"active":     active,
		"created_at": now,
		"updated_at": now,
	}
	if len(email) > 0 {
		m["email"] = email
	} else {
		m["email"] = nil
	}

	result, err := storage.GetRDBx().NamedExec(sqlstr, m)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// UserRepoMock is an implementation of IUserRepo which provides mocked users data
type UserRepoMock struct{}

// Create returns unchanged User struct given in argument
func (r *UserRepoMock) Create(u *User) (*User, error) {
	return u, nil
}
