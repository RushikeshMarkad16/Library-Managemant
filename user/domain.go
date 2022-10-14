package user

import "unicode"

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

type UserToDisplay struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	Last_name string `json:"last_name"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Mob_no    string `json:"mob_no"`
	Role      string `json:"role"`
}

type ChangePassword struct {
	ID          string `json:"id"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type FindByIDResponse struct {
	User User `json:"user"`
}

type FindByDataResponse struct {
	User User `json:"user"`
}

type ListResponse struct {
	Users []User `json:"users"`
	//Count int    `json:"total_count"` //to be updated
}

func (cr User) Validate() (err error) {
	if cr.FirstName == "" {
		return errEmptyFirstName
	}
	for _, r := range cr.FirstName {
		if !unicode.IsLetter(r) {
			return errInvalidFirstName
		}
	}
	if cr.Last_name == "" {
		return errEmptyLastName
	}
	for _, r := range cr.Last_name {
		if !unicode.IsLetter(r) {
			return errInvalidLastName
		}
	}
	if cr.Password == "" {
		return errEmptyPassword
	}
	if len(cr.Password) < 6 {
		return errInvalidPassword
	}
	if cr.Gender == "" || cr.Gender != "Male" && cr.Gender != "male" && cr.Gender != "Female" && cr.Gender != "female" {
		return errInvalidGender
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
	if cr.Role != "user" && cr.Role != "admin" {
		return errRoleType
	}

	// _, b := mail.ParseAddress(cr.Email)
	// if b != nil {
	// 	return errNotValidMail
	// }
	validateEmail := cr.Email
	flag := false
	lastapperance := 0
	for i := 0; i < len(validateEmail); i++ {
		if validateEmail[i] == '@' {
			flag = true
			lastapperance = i
		}
	}
	if !flag {
		return errNotValidMail
	}
	flag = false
	for i := lastapperance; i < len(validateEmail); i++ {
		if validateEmail[i] == '.' {
			flag = true
		}
	}
	if !flag {
		return errNotValidMail
	}
	for _, r := range cr.Email {
		if unicode.IsSpace(r) {
			return errNotValidMail
		}
	}
	// if len(cr.Mob_no) < 10 || len(cr.Mob_no) > 10 {
	if len(cr.Mob_no) != 10 {
		return errInvalidMobNo
	}
	for _, r := range cr.Mob_no {
		if !unicode.IsNumber(r) {
			return errInvalidMobNo
		}
	}

	return
}

func (up User) ValidateUpdate() (err error) {
	if up.FirstName == "" {
		return errInvalidFirstName
	}
	for _, r := range up.FirstName {
		if !unicode.IsLetter(r) {
			return errInvalidFirstName
		}
	}
	if up.Last_name == "" {
		return errInvalidLastName
	}
	for _, r := range up.Last_name {
		if !unicode.IsLetter(r) {
			return errInvalidLastName
		}
	}
	if up.ID == "" {
		return errEmptyID
	}
	return
}
