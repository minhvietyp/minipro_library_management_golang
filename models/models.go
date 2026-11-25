package models

type Book struct {
	Id string
	Title string
	Author string
	IsAvailable bool
	borrowerId string
}
