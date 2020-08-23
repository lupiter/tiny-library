package models

type Book struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	ISBN        string   `json:"isbn"`
	Year        string   `json:"year"`
	Publisher   string   `json:"publisher"`
	Tags        []string `json:"tags"`
	Identifier  int      `json:"id"`
	MaxLoanDays int      `json:"max_loan_days"`
	Location    string   `json:"location"`
	Format      string   `json:"format"`
}
