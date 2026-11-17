package main

import (
	"fmt"
	"library-system/library"
	"library-system/utils"
)

func main() {
	lib := library
	for {
		utils.ClearScreen()

		fmt.Println("--- Chuong trinh quan ly thu vien ---")
		fmt.Println("1. Them sach")
		fmt.Println("2. Xem danh sach sach")
		fmt.Println("3. Them nguoi muon")
		fmt.Println("4. Xem danh sach nguoi muon")
		fmt.Println("5. Muon sach")
		fmt.Println("6. Xem lich su muon")
		fmt.Println("7. Tra sach")
		fmt.Println("8. Tim kiem sach")
		fmt.Println("9. Thoat")
		choice := utils.GetPositiveInt("=> Chon chuc nang: ")

		utils.ClearScreen()

		switch choice {
		case 1:
			fmt.Println("======Them sach=====")
			if err := library.AddBook(); err != nil {
				fmt.Printf("XXX Loi khi them sach: %v \n", err)
			}
		case 2:
			fmt.Println("======Xem danh sach sach=====")
			if err := library.ListBook(); err != nil {
				fmt.Printf("XXX Loi xem danh sach them sach: %v \n", err)
			}
		case 3:
			fmt.Println("======Them nguoi muon sach=====")
			if err := library.AddBorrower(); err != nil {
				fmt.Printf("XXX Loi khi them nguoi muonm sach: %v \n", err)
			}
		case 4:
			fmt.Println("======Xem danh sach nguoi muon =====")
			if err := library.ListBorrower(); err != nil {
				fmt.Printf("XXX Loi khi xem danh sach nguoi muon: %v \n", err)
			}
		case 5:
			fmt.Println("======Muon sach =====")
			if err := library.BorrowBook(); err != nil {
				fmt.Printf("XXX Loi khi muon sach: %v \n", err)
			}
		case 6:
			fmt.Println("======Xem lich su muon sach =====")
			if err := library.ListBorrowHistory(); err != nil {
				fmt.Printf("XXX Loi khi xem lich su muon sach: %v \n", err)
			}
		case 7:
			fmt.Println("======Tra sach=====")
			if err := library.ReturnBook(); err != nil {
				fmt.Printf("XXX Loi khi tra sach: %v \n", err)
			}
		case 8:
			fmt.Println("======Tim kiem sach=====")
			if err := library.SearchBook(); err != nil {
				fmt.Printf("XXX Loi khi tim kiem sach: %v \n", err)
			}
		case 9:
			fmt.Println("======Thoat=====")
			return
		default:
			fmt.Println("XXX Lua chon khong hop le ")

		}

		utils.ReadInput("\n NHan enter de tiep tuc")
	}
}
