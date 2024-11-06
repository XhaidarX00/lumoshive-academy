package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"main/model"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"
)

func ChoseMenu(displayData func() int) int {
	var input string
	for {
		rangeMenu := displayData()
		fmt.Scan(&input)
		ClearScreen()

		IntInput, err := strconv.Atoi(input)
		if err != nil || IntInput < 1 || IntInput > rangeMenu {
			ErrorMessage("input tidak valid!")
			continue
		}

		if IntInput == rangeMenu {
			return 0
		}

		return IntInput
	}
}

func ChoseMenuCrud(displayData func(page string) int, page string) int {
	var input string
	for {
		rangeMenu := displayData(page)
		fmt.Scan(&input)
		ClearScreen()

		IntInput, err := strconv.Atoi(input)
		if err != nil || IntInput < 1 || IntInput > rangeMenu {
			ErrorMessage("input tidak valid!")
			continue
		}

		if IntInput == rangeMenu {
			return 0
		}

		return IntInput
	}
}

func ExitProgram() {
	defer os.Exit(0)
	ClearScreen()
	SuccesMessage("Keluar dari Program\n")
}

func ResetSessionTimeout(duration int) (context.Context, context.CancelFunc) {
	ctx := context.Background()
	deadline := time.Now().Add(time.Duration(duration) * time.Second)
	return context.WithDeadline(ctx, deadline)
}

// func applyOperation(a int, b int, operation func(int, int) int) int {
// 	return operation(a, b)
// }

func InputIndex(display func() int) int {
	var choice string
	for {
		rangeMenu := display()
		fmt.Scan(&choice)
		ClearScreen()

		intInput, err := strconv.Atoi(choice)
		if err != nil || intInput < 1 || intInput > rangeMenu {
			ErrorMessage(("Input harus berupa angka yang valid dan tidak boleh lebih dari 1 angka"))
			continue
		}

		return intInput
	}
}

// func ErrorMessage(text string) {
// 	red := color.New(color.FgRed).SprintFunc()
// 	err := errors.New(text)
// 	fmt.Printf("⚠️   %s%s\n", red("Error : "), err)
// }

func ErrorMessage(text string) {
	red := color.New(color.FgRed).SprintFunc()
	response := model.ResponseError{
		StatusCode: 400,
		Message:    text,
	}
	// fmt.Printf("\n⚠️  %s {StatusCode: %v, Message: %s}\n", red("Error : "), response.StatusCode, response.Message)
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		ErrorMessage("Failed to marshal JSON: " + err.Error())
		return
	}
	fmt.Printf("⚠️   %s%s\n", red("Error : "), string(jsonData))
}

func SuccesMessage(text string) {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()
	fmt.Printf("\n%s\n✔️  %s%s\n%s\n", yellow("=--------------------------------="), green("Succes : "), text, yellow("=--------------------------------="))
}

func ColorMessage(color_ string, text string) string {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()

	switch color_ {
	case "red":
		return red(text)
	case "green":
		return green(text)
	case "blue":
		return blue(text)
	case "yellow":
		return yellow(text)
	}

	return "0"
}

func PrintColorMsg(color string, text string) {
	fmt.Println(ColorMessage(color, text))
}

func Capitalize(s string) string {
	s = strings.ToLower(s)

	capitalizeNext := true
	result := []rune(s)

	for i, r := range result {
		if capitalizeNext && unicode.IsLetter(r) {
			result[i] = unicode.ToUpper(r)
			capitalizeNext = false
		} else if unicode.IsSpace(r) {
			capitalizeNext = true
		}
	}

	return string(result)
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		module := exec.Command("cmd", "/c", "cls")
		module.Stdout = os.Stdout
		module.Run()
	} else {
		module := exec.Command("clear")
		module.Stdout = os.Stdout
		module.Run()
	}
}

func IsLenVar(input any, len_ int) bool {
	lenVal := reflect.ValueOf(input)

	switch lenVal.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		if lenVal.Len() >= len_ {
			return true
		}
	default:
		ErrorMessage(fmt.Sprintf("tipe data %s tidak mendukung operasi Len()\n", lenVal.Kind()))
	}

	return false
}

func ConvertSliceToMap(slice interface{}) []map[string]interface{} {
	var data []map[string]interface{}

	val := reflect.ValueOf(slice)

	if val.Kind() != reflect.Slice {
		return nil
	}

	for i := 0; i < val.Len(); i++ {
		item := val.Index(i).Interface()
		itemMap := make(map[string]interface{})
		itemVal := reflect.ValueOf(item)

		for j := 0; j < itemVal.NumField(); j++ {
			field := itemVal.Type().Field(j)
			itemMap[field.Name] = itemVal.Field(j).Interface()
		}
		data = append(data, itemMap)
	}

	return data
}

func GetStructKeys(p interface{}, dk []string) []string {
	var keys []string
	val := reflect.ValueOf(p)

	if val.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		excluded := false
		if len(dk) != 0 {
			for _, v := range dk {
				if fieldName == v {
					excluded = true
					break
				}
			}
		}
		if !excluded {
			keys = append(keys, fieldName)
		}
	}

	return keys
}

func DisplayData(title string, data []map[string]interface{}, keys []string) {
	msg := fmt.Sprintf("\n============== ◉  %s ◉  ==============", title)
	PrintColorMsg("yellow", msg)
	fmt.Println(strings.Repeat("-", 50))
	for i, result := range data {
		fmt.Printf("%d. ", i+1)
		for _, key := range keys {
			value, ok := result[key]
			if !ok {
				value = "N/A"
			}

			val := reflect.ValueOf(value)
			switch val.Kind() {
			case reflect.String:
				fmt.Printf("%s | ", val.String())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fmt.Printf("%d | ", val.Int())
			case reflect.Float32, reflect.Float64:
				fmt.Printf("%f | ", val.Float())
			case reflect.Bool:
				fmt.Printf("%t | ", val.Bool())
			default:
				fmt.Printf("%v | ", val.Interface())
			}
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("-", 50))
}

func DisplayDataJson(result interface{}) {
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		ErrorMessage("Failed to marshal JSON: " + err.Error())
		return
	}
	SuccesMessage(string(jsonData))
}

func ConvertJsonToStruct(data interface{}, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(data); err != nil && err != io.EOF {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func WriteJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteErrorResponse(w http.ResponseWriter, message string, status int) {
	WriteJSONResponse(w, map[string]string{"error": message}, status)
}
