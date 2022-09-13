package book

type Book struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Author          string `json:"author"`
	Price           int    `json:"price"`
	TotalCopies     int    `json:"total_copies"`
	Status          string `json:"status"`
	AvailableCopies int    `json:"available_copies"`
}

type findByIDResponse struct {
	Book Book `json:"book"`
}

type listResponse struct {
	Books []Book `json:"books"`
	Count int    `json:"total_count"` //to be updated
}

func (cr Book) Validate() (err error) {
	if cr.Name == "" {
		return errEmptyName
	}
	return
}
