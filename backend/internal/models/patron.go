package models

type Patron struct {
	Name        string
	Identifier  int    `json:"id"`
	CardNumber  string `json:"card_number"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string
	Active      bool
	MaxLoanDays int `json:"max_loan_days"`
}
