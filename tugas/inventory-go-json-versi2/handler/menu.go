package handler

import (
	"fmt"
	"inventory/helper"
)

func Menu(handler InventoryHandler) {
	// helper.CallClear()
	Start()
	for {
		var menu int
		fmt.Scanln(&menu)
		if menu == 9 {
			break
		}
		switch menu {
		case 1:
			handler.ShowProduct()
		case 2:
			handler.InputPurchaseOrder()
		case 3:
			handler.InputSalesOrder()
		case 4:
			handler.ShowPurchaseOrderDetail()
		case 5:
			handler.ShowSalesOrderDetail()
		default:
			Start()
		}
	}
}

func Start() {
	helper.CallClear()
	fmt.Println("MENU PROGRAM INVENTORY")
	fmt.Println("==========================")
	fmt.Println("Tampilan menu")
	fmt.Println("1. Daftar product")
	fmt.Println("2. Tambah purchase order")
	fmt.Println("3. Tambah sales order*")
	fmt.Println("4. Lihat purchase order")
	fmt.Println("5. Lihat sales order*")
	fmt.Println()
	fmt.Println("9. Exit")
	fmt.Println()
	fmt.Print("Pilih menu:")
}
