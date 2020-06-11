package users

import (
	"github.com/andbabkin/pfms-api/internal/storage"
)

// IUserRepo is an interface to users data repository
type IUserRepo interface {
	Create(*User) (*User, error)
}

// UserRepo is an implementation of IUserRepo which connects to database
type UserRepo struct{}

// Create inserts a new record into users table and sets ID in User model
func (r *UserRepo) Create(u *User) error {
	sqlstr := `INSERT INTO ` + UserTable + ` (name, pswd, role, email, active, created_at, updated_at)
	VALUES (:name, :pswd, :role, :email, :active, :created_at, :updated_at)`

	result, err := storage.GetRDBx().NamedExec(sqlstr, u)
	if err == nil {
		id, _ := result.LastInsertId()
		if id > 0 {
			u.ID = uint64(id)
		}
	}

	return err
}

// UserRepoMock is an implementation of IUserRepo which provides mocked users data
type UserRepoMock struct{}

// Create returns unchanged User struct given in argument
func (r *UserRepoMock) Create(u *User) (*User, error) {
	return u, nil
}
