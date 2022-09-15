package user

import "errors"

var (
	errEmptyID        = errors.New("User ID must be present")
	errEmptyFirstName = errors.New("User First name must be present")
	errEmptyLastName  = errors.New("User Last name must be present")
	errNoUsers        = errors.New("No users present")
	errNoUserId       = errors.New("User is not present")
	errEmptyPassword  = errors.New("Password cannot be empty")
	errEmptyGender    = errors.New("User gender must be present")
	errEmptyAddress   = errors.New("Address must be present")
	errEmptyEmail     = errors.New("Email must be present")
	errEmptyMobNo     = errors.New("Mob no must be present")
	errEmptyRole      = errors.New("Role must be present")
	errRoleType       = errors.New("Enter a valid Role type")
	errNotValidMail   = errors.New("Email is not valid")
	errInvalidMobNo   = errors.New("Mob Number is not valid")
)
