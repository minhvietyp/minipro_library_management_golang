package library

import (
	"fmt"
	"library-system/utils"
)


// logic code 

func  AddBook(lib *Library) error {
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
	return nil
}

func ListBorrowHistory(lib *Library) error {
	return nil
}

func ReturnBook(lib *Library) error {
	return nil
}

func SearchBook(lib *Library) error {
	return nil
}
