package models

type Book struct {
	Title       string
	Author      string
	ISBN        string
	Year        string
	Publisher   string
	Tags        []string
	Identifier  int `json:"id"`
	MaxLoanDays int `json:"max_loan_days"`
	Location    string
	Format      string
}
