package library

import (
	"fmt"
	"library-system/models"
	"time"
)

// store data

type Library struct {
	book       map[string]models.Book
	borrower    map[string]models.Borrower
	transaction map[string]models.Transaction
}

func NewLibrary() *Library {
	return &Library{
		book:        make(map[string]models.Book),
		borrower:    make(map[string]models.Borrower),
		transaction: make(map[string]models.Transaction),
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

func (lib *Library) AddBorrowersStore(id, name, email string) error {
	if _, exists := lib.borrower[id]; exists {
		return fmt.Errorf("Sach voi id %s da  ton tai:", id)
	}
	lib.borrower[id] = models.Borrower{
		Id:    id,
		Name:  name,
		Email: email,
	}
	return nil
}

func (lib *Library) ListBorrowersStore() []models.Borrower {
	borrowers := make([]models.Borrower, 0, len(lib.borrower))
	for _, borrower := range lib.borrower {
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func (lib *Library) BorrowBooksStore(id, bookId, borrowerId string) error {
	book, bookExists := lib.book[bookId]
	if !bookExists {
		return fmt.Errorf("Sach voi Id %s khong ton tai", bookId)
	}

	if book.IsAvailable {
		return fmt.Errorf("Sach %s da duoc muon", book.Title)
	}

	
	if _, borrowerExists := lib.borrower[borrowerId]; !borrowerExists {
		return fmt.Errorf("Nguoi muon sach voi ID %s khong ton tai", borrowerId)
	}

	if _, transactionExists := lib.transaction[id]; transactionExists {
		return fmt.Errorf("Giao dich voi ID %s da ton tai", id)
	}
	// cap nhat trang thai sach va luu nguoc lai vao trong list book
	book.IsAvailable = true
	lib.book[bookId] = book
	lib.transaction[id] = models.Transaction{
		TransactionId:    id,
		BookId: bookId,
		BorrowerId: borrowerId,
		BorrowDate: time.Now(),
	}
	
	return nil
}

func (lib *Library) ListBorrowHistoryByBorrower(borrowerId string) []models.Transaction {
	if _, borrowerExists := lib.borrower[borrowerId]; !borrowerExists {
		return nil
	}

	var history []models.Transaction
	for _, transaction := range lib.transaction {
		if transaction.BorrowerId == borrowerId {
			history = append(history, transaction)
		}
	}

	return history
}
