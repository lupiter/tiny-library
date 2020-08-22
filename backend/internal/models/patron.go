package models

type Patron struct {
	Name        string `json:"name"`
	Identifier  int    `json:"id"`
	CardNumber  string `json:"card_number"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
	Active      bool   `json:"active"`
	MaxLoanDays int    `json:"max_loan_days"`
}
