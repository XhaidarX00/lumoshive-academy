package view

import (
	"fmt"
	"inventory/utils"
)

func MenuCrud(page string) int {
	fmt.Println(utils.ColorMessage("yellow", fmt.Sprintf("\n=---------- MENU CRUD %s ----------=", page)))
	pageCap := utils.Capitalize(page)
	fmt.Println("1. " + utils.ColorMessage("green", fmt.Sprintf("Create %s", pageCap)))
	fmt.Println("2. " + utils.ColorMessage("green", fmt.Sprintf("Read %s", pageCap)))
	fmt.Println("3. " + utils.ColorMessage("green", fmt.Sprintf("Update %s", pageCap)))
	fmt.Println("4. " + utils.ColorMessage("green", fmt.Sprintf("Delete %s", pageCap)))
	fmt.Println("5. " + utils.ColorMessage("red", "Kembali"))
	fmt.Print(utils.ColorMessage("yellow", "Pilih opsi: "))

	return 5
}

func MenuAdmin() int {
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- MENU ADMIN ----------="))
	fmt.Println("1. " + utils.ColorMessage("green", "Manage Product"))
	fmt.Println("2. " + utils.ColorMessage("green", "Manage Category"))
	fmt.Println("3. " + utils.ColorMessage("green", "Manage Location"))
	fmt.Println("4. " + utils.ColorMessage("green", "Search By (Category, Name, Item Code)"))
	fmt.Println("5. " + utils.ColorMessage("green", "Transaction History"))
	fmt.Println("6. " + utils.ColorMessage("red", "Keluar"))
	fmt.Print(utils.ColorMessage("yellow", "Pilih opsi: "))

	return 6
}

func MenuProduct() int {
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- PAGE PRODUCT ----------="))
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- MENU PRODUCT ----------="))
	fmt.Println("1. " + utils.ColorMessage("green", "CRUD Product"))
	fmt.Println("2. " + utils.ColorMessage("green", "Record Item Outflow"))
	fmt.Println("3. " + utils.ColorMessage("green", "Record Item Inflow"))
	fmt.Println("4. " + utils.ColorMessage("red", "Back"))
	fmt.Print(utils.ColorMessage("yellow", "Choose an option: "))

	return 4
}

func MenuCategory() int {
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- PAGE CATEGORY ----------="))
	rangeMenu := MenuCrud("CATEGORY")
	return rangeMenu
}

func MenuLocations() int {
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- PAGE LOCATION ----------="))
	rangeMenu := MenuCrud("LOCATION")
	return rangeMenu
}

func MenuSearch() int {
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- PAGE SEARCH ----------="))
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- MENU SEARCH ----------="))
	fmt.Println("1. " + utils.ColorMessage("green", "By Item Name"))
	fmt.Println("2. " + utils.ColorMessage("green", "By Category"))
	fmt.Println("3. " + utils.ColorMessage("green", "By Item Code"))
	fmt.Println("4. " + utils.ColorMessage("red", "Back"))
	fmt.Print(utils.ColorMessage("yellow", "Choose an option: "))

	return 4
}

func MenuTransaction() int {
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- PAGE TRANSACTION HISTORY ----------="))
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- MENU TRANSACTION HISTORY ----------="))
	fmt.Println("1. " + utils.ColorMessage("green", "History Item Inflow"))
	fmt.Println("2. " + utils.ColorMessage("green", "History Item Outflow"))
	fmt.Println("3. " + utils.ColorMessage("red", "Back"))
	fmt.Print(utils.ColorMessage("yellow", "Choose an option: "))

	return 3
}

func LoginMenu() int {
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- MENU ENDPOINT ----------="))
	fmt.Println("1. " + utils.ColorMessage("green", "login"))
	fmt.Println("2. " + utils.ColorMessage("red", "stop program"))
	fmt.Print(utils.ColorMessage("yellow", "Masukan Endpoint : "))

	return 2
}

func ChoiceMenu(role string) func() int {
	switch role {
	case "admin":
		return MenuAdmin
	default:
		return nil
	}
}
