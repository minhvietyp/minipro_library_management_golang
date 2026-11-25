package models

type Book struct {
	Id string
	Title string
	Author string
	IsAvailable bool
	borrowerId string
}

type Borrower struct {
	Id string
	Name string 
	Email string
}
