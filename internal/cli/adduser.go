package cli

import (
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/andbabkin/pfms-api/internal/storage/users"
)

// AddUser is a command to add a new user to the database
type AddUser struct{}

// Execute command
func (a *AddUser) Execute(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("not enough arguments, at least 3 should be given")
	}

	// get input
	name := args[0]
	pswd := []byte(args[1])
	utype, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return err
	}
	var email string
	if len(args) > 3 {
		email = args[3]
	}

	// check type
	if err := checkUserType(utype); err != nil {
		return err
	}

	// encrypt a password
	hash, err := bcrypt.GenerateFromPassword(pswd, 10)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, pswd)
	if err != nil {
		return err
	}

	// insert a new user record
	repo := users.UserRepo{}
	userID, err := repo.Create(name, string(hash), email, int8(utype), true)
	if err != nil {
		return err
	}

	fmt.Printf("Created a new user with ID: %d\n", userID)

	return nil
}

func checkUserType(t int64) error {
	switch t {
	default:
		return fmt.Errorf("wrong user type")
	case 0, 1, 2:
	}

	return nil
}
