package library

import (
	"fmt"
	"library-system/utils"
)

// logic code

func AddBook(lib *Library) error {
	id := utils.GenerateId()
	title := utils.GetNoneEmptyString("Nhap tieu de: ")
	author := utils.GetNoneEmptyString("Nhap ten tach gia: ")

	if err := lib.AddBooksStore(id, title, author); err != nil {
		return err
	}

	fmt.Printf("Them sach thanh cong voi ID \n: %s", id)

	return nil
}

func ListBook(lib *Library) error {
	books := lib.ListBooksStore()
	if len(books) == 0 {
		fmt.Println("Thu vien chua co sach vui long them sach")
		return nil
	}
	for _, book := range books {
		status := "Con"
		if book.IsAvailable {
			status = "Da muon"
		}
		fmt.Printf("id : %s, Tieu de: %s, Tac gia: %s, Trang thai: %s", book.Id, book.Title, book.Author, status)
	}

	return nil
}

func AddBorrower(lib *Library) error {
	id := utils.GenerateId()
	name := utils.GetNoneEmptyString("Nhap ten: ")
	email := utils.GetNoneEmptyString("Nhap email: ")

	if err := lib.AddBorrowersStore(id, name, email); err != nil {
		return err
	}

	fmt.Printf("Them nguoi muon thanh cong voi ID \n: %s", id)

	return nil
}

func ListBorrower(lib *Library) error {
	borrowers := lib.ListBorrowersStore()
	if len(borrowers) == 0 {
		fmt.Println("Thu vien chua co nguoi muon vui long them sach")
		return nil
	}
	fmt.Println("Danh sach nguoi muon")
	for _, borrower := range borrowers {
		fmt.Printf("id : %s, Ten: %s, Email: %s", borrower.Id, borrower.Name, borrower.Email)
	}

	return nil
}

func BorrowBook(lib *Library) error {
	transactionId := utils.GenerateId()
	bookId := utils.GetNoneEmptyString("Nhap ID sach: ")
	borrowerId := utils.GetNoneEmptyString("Nhap ID nguoi muon: ")

	if err := lib.BorrowBooksStore(transactionId , bookId, borrowerId); err != nil {
		return err
	}

	fmt.Printf("Muon sach thanh cong voi ID giao dich: %s", transactionId)


	return nil
}

func ListBorrowHistory(lib *Library) error {
	borrowerId := utils.GetNoneEmptyString("NHap ID nguoi muon: ")
	transaction := lib.ListBorrowHistoryByBorrower(borrowerId)
	if len(transaction) == 0 {
		fmt.Printf("Khong co lich su muon sach voi ID %s", borrowerId)
	}

	fmt.Println("Lich su muon sach: ")
	for _, transaction := range transaction {
		returnDate := "CHua tra"
		if !transaction.ReturnDate.IsZero() {
			returnDate = transaction.ReturnDate.Format("2006-01-01")
		}
		fmt.Printf("Giao dich: %s, Sach: %s, Ngay muon: %s, Ngay tra: %s", transaction.TransactionId, transaction.BookId, transaction.BorrowDate.Format("2006-01-01"), returnDate)
	}
	return nil
}

func ReturnBook(lib *Library) error {
	return nil
}

func SearchBook(lib *Library) error {
	return nil
}
