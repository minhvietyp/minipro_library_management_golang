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
