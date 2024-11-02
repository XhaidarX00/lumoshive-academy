package repository

import (
	"database/sql"
	"elearning/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

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

type UserRepositoryDB struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryDB {
	return UserRepositoryDB{DB: db}
}

// login
func (r *UserRepositoryDB) GetUsersLogin(user model.Users) (*model.Users, error) {
	query := `SELECT user_id, username, password, email, role FROM users WHERE username=$1 AND password=$2`
	var Us model.Users

	err := r.DB.QueryRow(query, user.Username, user.Password).Scan(&Us.User_id, &Us.Username, &Us.Password, &Us.Email, &Us.Role)

	if err != nil {
		return nil, err
	}

	return &Us, nil
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
func upsertWithAutoID(params CRUDParams) string {
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

	// var userIDNew string
	// if params.IsUpdate {
	// 	userIDNew = params.KeyUpdate
	// } else {
	// 	userIDNew = newID
	// }
	// msg := fmt.Sprintf("Operasi %s berhasil pada tabel %s dengan %s: %s",
	// 	map[bool]string{true: "update", false: "insert"}[params.IsUpdate], params.TableName, params.IDColumn, userIDNew)
	// utils.SuccesMessage(msg)
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

	return upsertWithAutoID(params)
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

	return upsertWithAutoID(params)
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

	// msg := fmt.Sprintf("Data berhasil dihapus dari tabel %s dengan kondisi: %s", params.TableName, params.Condition)
	// utils.SuccesMessage(msg)
	return ""
}
