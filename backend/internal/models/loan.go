package models

type Loan struct {
  Identifier string
  Book Book
  Patron Patron
  Lent string
  DueBack string // Lent + Patron's borrowing period or book's borrowing period whichever is shorter, here for convenience
  Returned string
}
