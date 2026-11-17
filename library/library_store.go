package library

import "library-system/models"

type Library struct {
	book map[string]models.Book
}