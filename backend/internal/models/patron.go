package models

type Patron struct {
	Name        string
	Identifier  string
	CardNumber  string
	DateOfBirth string
	Address     string
	Active      bool
	MaxLoanDays int
}
