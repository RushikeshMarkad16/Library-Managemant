package book

import (
	"strconv"

	valid "github.com/asaskevich/govalidator"
)

type Book struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Author          string `json:"author"`
	Price           int    `json:"price"`
	TotalCopies     int    `json:"total_copies"`
	Status          string `json:"status"`
	AvailableCopies int    `json:"available_copies"`
}

type BookToDisplay struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Status string `json:"status"`
}

type FindByIDResponse struct {
	Book Book `json:"book"`
}

type ListResponse struct {
	Books []Book `json:"books"`
}

func (cr Book) Validate() (err error) {
	if cr.Name == "" {
		return errEmptyName
	}
	if cr.Author == "" {
		return errEmptyAuthor
	}

	if cr.TotalCopies == 0 {
		return errZeroCopies
	}

	t := strconv.Itoa(cr.TotalCopies)
	if !valid.IsInt(t) {
		return errInvalidTotalCopies
	}
	if cr.Price < 1 {
		return errInvalidPrice
	}
	t2 := strconv.Itoa(cr.Price)
	if !valid.IsInt(t2) {
		return errInvalidPrice
	}

	if cr.Status != "available" {
		return errInvalidStatus
	}
	if cr.AvailableCopies > cr.TotalCopies {
		return errInvalidAvailableCopies
	}
	t1 := strconv.Itoa(cr.AvailableCopies)
	if !valid.IsInt(t1) {
		return err1InvalidAvailableCopies
	}

	return
}
