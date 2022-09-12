package user

import "github.com/RushikeshMarkad16/Library-Managemant/db"

type updateRequest struct {
	ID         string `json:"id"`
	First_Name string `json:"name"`
}

type createRequest struct {
	First_Name string `json:"first_name"`
}

type findByIDResponse struct {
	User db.User `json:"user"`
}

type listResponse struct {
	Users []db.User `json:"users"`
}

func (cr createRequest) Validate() (err error) {
	if cr.First_Name == "" {
		return errEmptyName
	}
	return
}

func (ur updateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.First_Name == "" {
		return errEmptyName
	}
	return
}
