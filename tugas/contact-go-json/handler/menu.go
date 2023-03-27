package handler

import (
	"contact-go/helper"
	"fmt"
)

func Menu(handler ContactHandler) {
	helper.ClearTerminal()
	helper.ShowMenuList()

	for {
		var menu int
		fmt.Scanln(&menu)

		if menu == 5 {
			helper.ClearTerminal()
			break
		}

		switch menu {
		case 0:
			helper.ClearTerminal()
			helper.ShowMenuList()
		case 1:
			helper.ClearTerminal()
			handler.List()
		case 2:
			helper.ClearTerminal()
			handler.Add()
		case 3:
			helper.ClearTerminal()
			handler.Update()
		case 4:
			helper.ClearTerminal()
			handler.Delete()
		}
	}
}
