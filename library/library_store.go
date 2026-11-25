package library

import (
	"fmt"
	"library-system/models"
)

// store data

type Library struct {
	book map[string]models.Book
}

func NewLibrary() *Library {
	return &Library{
		book: make(map[string]models.Book),
	}
}

func (lib *Library) AddBooksStore(id, title, author string) error {
	if _, exists := lib.book[id]; exists {
		return fmt.Errorf("Sach voi id %s da  ton tai:", id)
	}
	lib.book[id] = models.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}
	return nil
}

// dau tien tao slice
// Lay du lieu tu map day vo slice
func (lib *Library) ListBooksStore() []models.Book {
	books := make([]models.Book, 0, len(lib.book))
	for _, book := range lib.book {
		books = append(books, book)
	}
	return books
}


