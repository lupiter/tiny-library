package models

type Loan struct {
	Identifier int    `json:"id"`
	Book       Book   `json:"book"`
	Patron     Patron `json:"patron"`
	Lent       string `json:"lent"`
	DueBack    string `json:"due_back"` // Lent + Patron's borrowing period or book's borrowing period whichever is shorter, here for convenience
	Returned   string `jons:"returned"`
}
