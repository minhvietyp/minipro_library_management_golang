package models

import "time"

type Book struct {
	Id          string
	Title       string
	Author      string
	IsAvailable bool
	borrowerId  string
}

type Borrower struct {
	Id    string
	Name  string
	Email string
}

type Transaction struct {
	TransactionId         string
	BookId     string
	BorrowerId string
	BorrowDate time.Time
	ReturnDate time.Time
}
