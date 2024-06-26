package domain

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/n30w/Darkspace/internal/models"
)

// validateCredentials validates credentials using credentials interface
// method. Firstly, the credentials are checked if they are blank.
// Then, in each Valid() method, specific requirements for the credentials are
// checked.
func validateCredentials(c *models.User) error {
	var err error

	err = c.Credentials.Username.Valid()
	if err != nil {
		return err
	}

	err = c.Credentials.Password.Valid()
	if err != nil {
		return err
	}

	err = c.Credentials.Email.Valid()
	if err != nil {
		return err
	}

	err = c.Credentials.Membership.Valid()
	if err != nil {
		return err
	}

	return nil
}

// Password is a hashed string from the frontend.
type Password string

func (p Password) Valid() error {
	if p == "" {
		return errors.New("password field empty")
	}

	var (
		hasMinLen      bool = len(p) >= 8
		hasUpper       bool
		hasLower       bool
		hasNumber      bool
		hasSpecialChar bool
	)

	for _, char := range p {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	if !hasMinLen {
		return errors.New("password must be at least 8 characters long")
	}
	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must contain at least one digit")
	}
	if !hasSpecialChar {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func (p Password) String() string {
	return string(p)
}

// Username is a string defined by the user they can
// use to login.
type Username string

func (u Username) Valid() error {
	if len(u) == 0 {
		return errors.New("username cannot be empty")
	}

	if len(u) < 3 {
		return errors.New("username must be at least 3 characters long")
	}

	return nil
}

func (u Username) String() string {
	return string(u)
}

// Email is a valid NYU email address.
type Email string

func (e Email) Valid() error {
	if len(e) == 0 {
		return errors.New("email cannot be empty")
	}

	if !strings.Contains(string(e), "nyu.edu") {
		return errors.New("email must contain nyu.edu")
	}

	// TODO check if its after the @ symbol.
	if !strings.Contains(string(e), ".") {
		return errors.New("email must have a TLD")
	}

	atIndex := strings.Index(string(e), "@")
	if !(atIndex > 2) { // Checking for the local part to be more than two
		// characters.
		return errors.New("email local part must be more than two characters")
	}

	return nil
}

func (e Email) String() string {
	return string(e)
}

// Membership defines the type of permissions that a user is default
// scoped to. There are only two valid Membership possibilities for
// a POST request can add or change in the database, 0 and 1. Although
// there are integers greater than 1 defined, such as ADMIN,
// this is not supposed to be accessible by the frontend, and therefore,
// not bothered to be checked.
type Membership int

func (m Membership) Valid() error {
	if m < 0 || m > 1 {
		return errors.New("membership must either be 0 or 1")
	}

	return nil
}

func (m Membership) String() string {
	return fmt.Sprintf("%d", m)
}
