package user

import "errors"

var (
	errEmptyID          = errors.New("User ID must be present")
	errInvalidFirstName = errors.New("Invalid First name")
	errInvalidLastName  = errors.New("Invalid Last name")
	errNoUsers          = errors.New("No users present")
	errNoUserId         = errors.New("User is not present")
	errWrongPassword    = errors.New("Wrong Password")
	errEmptyPassword    = errors.New("Password cannot be empty")
	errInvalidPassword  = errors.New("Password length must be 6 or more")
	errInvalidGender    = errors.New("Enter valid gender")
	errEmptyAddress     = errors.New("Address must be present")
	errEmptyEmail       = errors.New("Email must be present")
	errEmptyMobNo       = errors.New("Mob no must be present")
	errEmptyRole        = errors.New("Role must be present")
	errRoleType         = errors.New("Enter a valid Role type from user and admin")
	errNotValidMail     = errors.New("Email is not valid")
	errInvalidMobNo     = errors.New("Mob Number is not valid")
)
