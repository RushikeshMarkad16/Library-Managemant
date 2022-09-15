package user

import "net/mail"

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	Last_name string `json:"last_name"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Mob_no    string `json:"mob_no"`
	Role      string `json:"role"`
}

type findByIDResponse struct {
	User User `json:"user"`
}

type listResponse struct {
	Users []User `json:"users"`
	Count int    `json:"total_count"` //to be updated
}

func (cr User) Validate() (err error) {
	if cr.FirstName == "" {
		return errEmptyFirstName
	}
	if cr.Last_name == "" {
		return errEmptyLastName
	}
	if cr.Password == "" {
		return errEmptyPassword
	}
	if cr.Gender == "" {
		return errEmptyGender
	}
	if cr.Address == "" {
		return errEmptyAddress
	}
	if cr.Email == "" {
		return errEmptyEmail
	}
	if cr.Mob_no == "" {
		return errEmptyMobNo
	}
	if cr.Role == "" {
		return errEmptyRole
	}
	if cr.Role != "user" && cr.Role != "admin" && cr.Role != "superadmin" {
		return errRoleType
	}
	_, b := mail.ParseAddress(cr.Email)
	if b != nil {
		return errNotValidMail
	}
	if len(cr.Mob_no) < 10 || len(cr.Mob_no) > 10 {
		return errInvalidMobNo
	}

	return
}
