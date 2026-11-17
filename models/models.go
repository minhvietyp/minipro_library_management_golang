package models

type Book struct {
	Id string
	Title string
	Author string
	isAvailable bool
	borrowerId string
}
