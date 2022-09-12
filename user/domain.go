package user

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
		return errEmptyName
	}
	return
}
