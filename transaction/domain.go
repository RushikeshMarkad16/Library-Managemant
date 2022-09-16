package transaction

type Transaction struct {
	ID         string `json:"id"`
	IssueDate  int    `json:"issue_date"`
	DueDate    int    `json:"due_date"`
	ReturnDate int    `json:"return_date"`
	BookID     string `json:"book_id"`
	UserID     string `json:"user_id"`
}

type listResponse struct {
	Transactions []Transaction `json:"transactions"`
	Count        int           `json:"total_count"` //to be updated
}

func (cr Transaction) Validate() (err error) {
	if cr.BookID == "" {
		return errEmptyBookID
	}
	if cr.UserID == "" {
		return errEmptyUserID
	}

	return
}
