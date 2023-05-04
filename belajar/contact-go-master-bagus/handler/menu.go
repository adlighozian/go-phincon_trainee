package handler

import (
	"contact-go/helper"
	"fmt"
)

func Menu(handler ContactHandler) {
	err := helper.ClearTerminal()
	if err != nil {
		fmt.Println(err)
	}

	helper.ShowMenuList()

	for {
		var menu int
		fmt.Scanln(&menu)

		if menu == 6 {
			err := helper.ClearTerminal()
			if err != nil {
				fmt.Println(err)
			}
			break
		}

		switch menu {
		default: // case 0 atau selain 0
			err := helper.ClearTerminal()
			if err != nil {
				fmt.Println(err)
			}
			helper.ShowMenuList()
		case 1:
			handler.List()
		case 2:
			handler.Add()
		case 3:
			handler.Detail()
		case 4:
			handler.Update()
		case 5:
			handler.Delete()
		}
	}
}
