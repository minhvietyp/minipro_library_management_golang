package library

import (
	"fmt"
	"library-system/models"
	"library-system/utils"
)




func  AddBook(lib *Library) error {
	id := utils.GenerateId()
	title := utils.GetNoneEmptyString("Nhap tieu de: ")
	author := utils.GetNoneEmptyString("Nhap ten tach gia: ")

	lib.book[id]  = models.Book{
		Id: id,
		Title: title,
		Author: author,

	}
	

	fmt.Println("Them sach thanh cong")
	fmt.Printf("Book: %+v", lib.book)

	return nil
}

func ListBook() error {
	return nil
}

func AddBorrower() error {
	return nil
}

func ListBorrower() error {
	return nil
}

func BorrowBook() error {
	return nil
}

func ListBorrowHistory() error {
	return nil
}

func ReturnBook() error {
	return nil
}

func SearchBook() error {
	return nil
}
