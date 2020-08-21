package models

type Book struct {
	Title       string
	Author      string
	ISBN        string
	Year        string
	Publisher   string
	Tags        []string
	Identifier  int
	MaxLoanDays int
	Location    string
	Format      string
}
