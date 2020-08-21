package models

type Patron struct {
	Name        string
	Identifier  int
	CardNumber  string
	DateOfBirth string
	Address     string
	Active      bool
	MaxLoanDays int
}
