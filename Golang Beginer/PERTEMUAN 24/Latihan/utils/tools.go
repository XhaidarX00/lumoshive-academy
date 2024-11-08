package utils

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/fatih/color"
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

func ErrorMessage(text string) {
	red := color.New(color.FgRed).SprintFunc()
	err := errors.New(text)
	fmt.Printf("⚠️   %s%s\n", red("Error : "), err)
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

type CRUDParams struct {
	DB             *sql.DB
	TableName      string
	IDColumn       string
	Prefix         string
	Columns        []string
	Values         []interface{}
	IsUpdate       bool
	KeyUpdate      string
	KeyDelete      string
	JSONPath       string
	JSONUPDATEPath string
	DataStruct     interface{}
	Condition      string // Kondisi untuk operasi READ atau DELETE
}

// Fungsi untuk mendapatkan ID terkecil yang hilang atau ID berikutnya
func getNextID(params CRUDParams) (string, string) {
	var missingID sql.NullString
	queryMissingID := fmt.Sprintf(`
		SELECT '%s' || MIN((regexp_replace(%s, '[^0-9]', '', 'g')::INT + 1)::TEXT)
		FROM %s
		WHERE '%s' || (regexp_replace(%s, '[^0-9]', '', 'g')::INT + 1)::TEXT NOT IN (SELECT %s FROM %s);
	`, params.Prefix, params.IDColumn, params.TableName, params.Prefix, params.IDColumn, params.IDColumn, params.TableName)

	err := params.DB.QueryRow(queryMissingID).Scan(&missingID)
	if err != nil {
		return "", fmt.Sprintf("error retrieving missing ID: %v", err)
	}

	if !missingID.Valid {
		var maxID int
		queryMaxID := fmt.Sprintf(`SELECT COALESCE(MAX(regexp_replace(%s, '[^0-9]', '', 'g')::INT), 0) + 1 FROM %s`, params.IDColumn, params.TableName)
		err = params.DB.QueryRow(queryMaxID).Scan(&maxID)
		if err != nil {
			return "", fmt.Sprintf("error retrieving max ID: %v", err)
		}
		missingID.String = fmt.Sprintf("%s%d", params.Prefix, maxID)
	}

	return missingID.String, ""
}

// Fungsi untuk menjalankan operasi insert atau update dengan ID otomatis
func UpsertWithAutoID(params CRUDParams) string {
	newID, errG := getNextID(params)
	if errG != "" {
		return errG
	}

	if len(params.Values) != len(params.Columns) {
		return fmt.Sprintf("jumlah values (%d) tidak sesuai dengan jumlah kolom (%d)", len(params.Values), len(params.Columns))
	}

	// Tambahkan ID baru ke kolom dan value jika insert
	if !params.IsUpdate {
		params.Columns = append([]string{params.IDColumn}, params.Columns...)
		params.Values = append([]interface{}{newID}, params.Values...)
	}

	// Buat query untuk insert atau update
	query := ""
	if params.IsUpdate {
		setClauses := make([]string, len(params.Columns))
		for i, column := range params.Columns {
			setClauses[i] = fmt.Sprintf("%s=$%d", column, i+1)
		}
		query = fmt.Sprintf("UPDATE %s SET %s WHERE %s=$%d", params.TableName, strings.Join(setClauses, ", "), params.IDColumn, len(params.Columns)+1)
		params.Values = append(params.Values, params.KeyUpdate)
	} else {
		placeholders := make([]string, len(params.Columns))
		for i := range placeholders {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
		}
		query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", params.TableName, strings.Join(params.Columns, ", "), strings.Join(placeholders, ", "))
	}

	// fmt.Printf("%v\n", params.Values)
	// fmt.Println(query)
	_, err := params.DB.Exec(query, params.Values...)
	if err != nil {
		return fmt.Sprintf("error executing upsert query: %v", err)
	}

	return ""
}

// Fungsi untuk menambahkan data dari JSON
func UpsertFromJSON(params CRUDParams) string {
	var fileData []byte
	var err error

	if params.IsUpdate {
		fileData, err = ioutil.ReadFile(params.JSONUPDATEPath)
	} else {
		fileData, err = ioutil.ReadFile(params.JSONPath)
	}

	if err != nil {
		return fmt.Sprintf("error reading JSON file: %v", err)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(fileData, &jsonData); err != nil {
		return fmt.Sprintf("error unmarshalling JSON data: %v", err)
	}

	params.Columns = make([]string, 0, len(jsonData))
	params.Values = make([]interface{}, 0, len(jsonData))
	for key, value := range jsonData {
		params.Columns = append(params.Columns, key)
		params.Values = append(params.Values, value)
	}

	return UpsertWithAutoID(params)
}

// Fungsi untuk menambahkan data dari struct
func UpsertFromStruct(params CRUDParams) string {
	val := reflect.ValueOf(params.DataStruct)
	typ := reflect.TypeOf(params.DataStruct)

	if val.Kind() != reflect.Struct {
		return fmt.Sprintf("data harus berupa struct")
	}

	params.Columns = make([]string, 0, val.NumField())
	params.Values = make([]interface{}, 0, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		params.Columns = append(params.Columns, field.Name)
		params.Values = append(params.Values, val.Field(i).Interface())
	}

	return UpsertWithAutoID(params)
}

func ReadData(params CRUDParams) ([]map[string]interface{}, string) {
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(params.Columns, ", "), params.TableName)
	if params.Condition != "" {
		query += " WHERE " + params.Condition
	}

	rows, err := params.DB.Query(query)
	if err != nil {
		return nil, fmt.Sprintf("error executing read query: %v", err)
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Sprintf("error getting columns: %v", err)
	}

	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(sql.NullString)
	}

	var results []map[string]interface{}

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			return nil, fmt.Sprintf("error scanning row: %v", err)
		}

		rowData := make(map[string]interface{})
		for i, col := range columns {
			val := values[i].(*sql.NullString)
			if val.Valid {
				rowData[col] = val.String
			} else {
				rowData[col] = nil
			}
		}
		results = append(results, rowData)
	}

	return results, ""
}

// Fungsi untuk menghapus data berdasarkan ID atau kondisi tertentu
func DeleteData(params CRUDParams) string {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", params.TableName, params.Condition)

	_, err := params.DB.Exec(query)
	if err != nil {
		return fmt.Sprintf("error executing delete query: %v", err)
	}

	return ""
}
