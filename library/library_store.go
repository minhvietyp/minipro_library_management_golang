package library

import (
	"flag"
	"fmt"
	"library-system/models"
	"time"
)

// store data

type Library struct {
	books       map[string]models.Book
	borrowers   map[string]models.Borrower
	transactions map[string]models.Transaction
}

func NewLibrary() *Library {
	return &Library{
		books:        make(map[string]models.Book),
		borrowers:    make(map[string]models.Borrower),
		transactions: make(map[string]models.Transaction),
	}
}

func (lib *Library) AddBooksStore(id, title, author string) error {
	if _, exists := lib.books[id]; exists {
		return fmt.Errorf("Sach voi id %s da  ton tai:", id)
	}
	lib.books[id] = models.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}
	return nil
}

// dau tien tao slice
// Lay du lieu tu map day vo slice
func (lib *Library) ListBooksStore() []models.Book {
	books := make([]models.Book, 0, len(lib.books))
	for _, book := range lib.books {
		books = append(books, book)
	}
	return books
}

func (lib *Library) AddBorrowersStore(id, name, email string) error {
	if _, exists := lib.borrowers[id]; exists {
		return fmt.Errorf("Sach voi id %s da  ton tai:", id)
	}
	lib.borrowers[id] = models.Borrower{
		Id:    id,
		Name:  name,
		Email: email,
	}
	return nil
}

func (lib *Library) ListBorrowersStore() []models.Borrower {
	borrowers := make([]models.Borrower, 0, len(lib.borrowers))
	for _, borrower := range lib.borrowers {
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func (lib *Library) BorrowBooksStore(id, bookId, borrowerId string) error {
	book, bookExists := lib.books[bookId]
	if !bookExists {
		return fmt.Errorf("Sach voi Id %s khong ton tai", bookId)
	}

	if book.IsAvailable {
		return fmt.Errorf("Sach %s da duoc muon", book.Title)
	}

	if _, borrowerExists := lib.borrowers[borrowerId]; !borrowerExists {
		return fmt.Errorf("Nguoi muon sach voi ID %s khong ton tai", borrowerId)
	}

	if _, transactionExists := lib.transactions[id]; transactionExists {
		return fmt.Errorf("Giao dich voi ID %s da ton tai", id)
	}
	// cap nhat trang thai sach va luu nguoc lai vao trong list book
	book.IsAvailable = true
	lib.books[bookId] = book
	lib.transactions[id] = models.Transaction{
		TransactionId: id,
		BookId:        bookId,
		BorrowerId:    borrowerId,
		BorrowDate:    time.Now(),
	}

	return nil
}

func (lib *Library) ListBorrowHistoryByBorrowerstore(borrowerId string) []models.Transaction {
	if _, borrowerExists := lib.borrowers[borrowerId]; !borrowerExists {
		return nil
	}

	var history []models.Transaction
	for _, transaction := range lib.transactions {
		if transaction.BorrowerId == borrowerId {
			history = append(history, transaction)
		}
	}

	return history
}

func (lib * Library) GetBookTitleStore(bookId string) string {
	book, _ := lib.books[bookId]
	return book.Title
}


func (lib *Library) ReturnBookStore(transactionId  string) error {
	transaction, exists := lib.transactions[transactionId]
	if !exists {
		return fmt.Errorf("Giao dich khong ton tai voi ID: %s", transactionId)
	}

	if !transaction.ReturnDate.IsZero() {
		return fmt.Errorf("Sach voi ID %s da duoc tra", transactionId)
	}

	// update status of book
	book, bookExists := lib.books[transaction.BookId]
	if !bookExists {
		return fmt.Errorf("Sach voi Id %s khong ton tai", transaction.BookId)
	}

	book.IsAvailable = false
	lib.books[transaction.BookId] = book

	//update returnDate of transaction
	transaction.ReturnDate = time.Now()
	lib.transactions[transactionId] = transaction
	

	return nil
}