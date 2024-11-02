package view

import (
	"elearning/utils"
	"fmt"
)

func MenuCrud(page string) int {
	fmt.Println(utils.ColorMessage("yellow", fmt.Sprintf("\n=---------- %s PAGE ----------=", page)))
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- MENU CRUD ----------="))
	fmt.Println("1. " + utils.ColorMessage("green", "Create"))
	fmt.Println("2. " + utils.ColorMessage("green", "Read"))
	fmt.Println("3. " + utils.ColorMessage("green", "Update"))
	fmt.Println("4. " + utils.ColorMessage("green", "Delete"))
	fmt.Println("5. " + utils.ColorMessage("red", "Kembali"))
	fmt.Print(utils.ColorMessage("yellow", "Pilih opsi: "))

	return 5
}

func MenuAdmin() int {
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- MENU ADMIN ----------="))
	fmt.Println("1. " + utils.ColorMessage("green", "Mentor"))
	fmt.Println("2. " + utils.ColorMessage("green", "Student"))
	fmt.Println("3. " + utils.ColorMessage("green", "Material"))
	fmt.Println("4. " + utils.ColorMessage("green", "Schedule"))
	fmt.Println("5. " + utils.ColorMessage("green", "Announcement"))
	fmt.Println("6. " + utils.ColorMessage("red", "Keluar"))
	fmt.Print(utils.ColorMessage("yellow", "Pilih opsi: "))

	return 6
}

func MenuStudent() int {
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- MENU STUDENT ----------="))
	fmt.Println("1. " + utils.ColorMessage("green", "Class Schedule"))
	fmt.Println("2. " + utils.ColorMessage("green", "Material"))
	fmt.Println("3. " + utils.ColorMessage("green", "Absence"))
	fmt.Println("4. " + utils.ColorMessage("green", "Grades"))
	fmt.Println("5. " + utils.ColorMessage("green", "weekly assignment"))
	fmt.Println("6. " + utils.ColorMessage("red", "Kembali"))
	fmt.Print(utils.ColorMessage("yellow", "Pilih opsi: "))

	return 6
}

func MenuMentor() int {
	fmt.Println(utils.ColorMessage("yellow", "\n=---------- MENU MENTOR ----------="))
	fmt.Println("1. " + utils.ColorMessage("green", "Class Schedule"))
	fmt.Println("2. " + utils.ColorMessage("green", "Material"))
	fmt.Println("3. " + utils.ColorMessage("green", "Absence Mentor"))
	fmt.Println("4. " + utils.ColorMessage("green", "Absence Students"))
	fmt.Println("5. " + utils.ColorMessage("green", "Grades"))
	fmt.Println("6. " + utils.ColorMessage("green", "weekly assignment"))
	fmt.Println("7. " + utils.ColorMessage("red", "Kembali"))
	fmt.Print(utils.ColorMessage("yellow", "Pilih opsi: "))

	return 7
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

	case "mentor":
		return MenuMentor

	case "student":
		return MenuStudent
	default:
		return nil
	}
}
