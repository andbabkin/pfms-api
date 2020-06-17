package users

import (
	"database/sql"
	"errors"
	"time"

	"github.com/andbabkin/pfms-api/internal/storage"
)

// ErrUserNotFound should be returned if sql.ErrNoRows is received from user record query
var ErrUserNotFound error = errors.New("User not found")

// IUserRepo is an interface to users data repository
type IUserRepo interface {
	Create(*User) error
	Find(uint64) (*User, error)
	FindLastByField(string, interface{}) (*User, error)
	UpdateField(field string, u *User) error
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

// Find fetches the user record by provided id.
func (r *UserRepo) Find(id uint64) (*User, error) {
	sqlstr := storage.FindByIDSQL(id, UserTable)
	u := &User{}
	err := storage.GetRDBx().Get(u, sqlstr)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}

	return u, err
}

// FindLastByField finds last record where field equals to value.
func (r *UserRepo) FindLastByField(field string, value interface{}) (*User, error) {
	sqlstr := storage.FindByFieldSQL(field, UserTable) + " ORDER BY id DESC LIMIT 1"
	u := &User{}
	stmt, err := storage.GetRDBx().PrepareNamed(sqlstr)
	if err != nil {
		return nil, err
	}
	err = stmt.Get(u, map[string]interface{}{field: value})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}

	return u, err
}

// UpdateField takes field's value from User and updates record in database.
// Field UpdatedAt is set to current time.
func (r *UserRepo) UpdateField(field string, u *User) error {
	sqlstr := storage.UpdateFieldSQL(field, UserTable)
	u.UpdatedAt = time.Now()
	_, err := storage.GetRDBx().NamedExec(sqlstr, u)
	return err
}

// UserRepoMock is an implementation of IUserRepo which provides mocked users data
type UserRepoMock struct {
	User  *User
	Users []User
}

// Create does nothing and always returns nil
func (rm *UserRepoMock) Create(u *User) error {
	return nil
}

// Find returns predefined User
func (rm *UserRepoMock) Find(id uint64) (*User, error) {
	return rm.User, nil
}

// FindLastByField returns predefined User
func (rm *UserRepoMock) FindLastByField(field string, value interface{}) (*User, error) {
	return rm.User, nil
}

// UpdateField does nothing and always returns nil
func (rm *UserRepoMock) UpdateField(field string, u *User) error {
	return nil
}
