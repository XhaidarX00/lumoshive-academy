package service

import (
	"database/sql"
	"fmt"
	"inventory/model"
	"inventory/repository"
	"inventory/utils"
	"inventory/view"
	"strings"
)

type UserService struct {
	RepoUser repository.UserRepositoryDB
}

func NewUserService(repo repository.UserRepositoryDB) *UserService {
	return &UserService{RepoUser: repo}
}

func (usr *UserService) LoginService(user model.Users) (*model.Response, string, error) {

	userReslut, err := usr.RepoUser.GetUsersLogin(user)

	if err != nil {
		return nil, "", err
	}
	response := model.Response{
		StatusCode: 200,
		Message:    "login success",
		Data:       userReslut,
	}
	return &response, userReslut.Username, nil
}

func ManageProduct(db *sql.DB) bool {
	index := utils.ChoseMenu(view.MenuProduct)
	if index == 0 {
		return false
	}

	params := repository.CRUDParams{DB: db}

	switch index {
	case 1:
		params.TableName = "products"
		params.IDColumn = "id"
		params.Prefix = "pd"
		params.KeyUpdate = "pd21"
		params.KeyDelete = "pd22"
		params.Page = 1
		params.Limit = 5
		params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/newProduct.json"
		params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/updateProduct.json"

		chosenIndex := utils.ChoseMenuCrud(view.MenuCrud, strings.ToUpper(params.TableName))
		if chosenIndex == 0 {
			return false
		}

		if !CrudChose(chosenIndex, db, params) {
			return false
		}

	case 2:
		params.TableName = "products"
		params.Page = 1
		params.Limit = 2
		params.Under10 = true

		results, err := repository.GetProductsPaginated(db, params)
		if err != "" {
			utils.ErrorMessage(err)
			return false
		}
		utils.DisplayDataJson(model.Response{StatusCode: 200, Message: "Read Success", Data: results})

	case 3:
		params.TableName = "transactions"
		params.IDColumn = "id"
		params.Prefix = "trx"
		params.KeyUpdate = "trx21"
		params.KeyDelete = "trx22"
		params.Page = 5
		params.Limit = 5
		params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/newTransactions.json"
		params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/updateTransactions.json"

		chosenIndex := utils.ChoseMenuCrud(view.MenuCrud, strings.ToUpper(params.TableName))
		if chosenIndex == 0 {
			return false
		}

		if !CrudChose(chosenIndex, db, params) {
			return false
		}

	default:
		utils.ErrorMessage("menu tidak ditemukan!")
		return false
	}

	return true
}

func ManageCategory(db *sql.DB) bool {
	params := repository.CRUDParams{DB: db}
	params.TableName = "categories"
	params.IDColumn = "id"
	params.Prefix = "ctg"
	params.KeyUpdate = "ctg6"
	params.KeyDelete = "ctg6"
	params.Page = 3
	params.Limit = 2
	params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/newCategory.json"
	params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/updateCategory.json"

	fmt.Println(utils.ColorMessage("yellow", "\n=---------- PAGE CATEGORY ----------="))
	chosenIndex := utils.ChoseMenuCrud(view.MenuCrud, strings.ToUpper(params.TableName))
	if chosenIndex == 0 {
		return false
	}

	if !CrudChose(chosenIndex, db, params) {
		return false
	}

	return true
}

func ManageLocations(db *sql.DB) bool {
	params := repository.CRUDParams{DB: db}
	params.TableName = "locations"
	params.IDColumn = "id"
	params.Prefix = "loc"
	params.KeyUpdate = "loc5"
	params.KeyDelete = "loc6"
	params.Page = 1
	params.Limit = 2
	params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/newLocation.json"
	params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/updateLocation.json"

	fmt.Println(utils.ColorMessage("yellow", "\n=---------- PAGE LOCATIONS ----------="))
	chosenIndex := utils.ChoseMenuCrud(view.MenuCrud, strings.ToUpper(params.TableName))
	if chosenIndex == 0 {
		return false
	}

	if !CrudChose(chosenIndex, db, params) {
		return false
	}

	return true
}

func SearchBy(db *sql.DB) bool {
	index := utils.ChoseMenu(view.MenuSearch)
	if index == 0 {
		return false
	}

	params := repository.CRUDParams{DB: db}
	params.TableName = "products"
	params.Page = 1
	params.Limit = 5

	switch index {
	case 1: // name
		params.BY = "p.name"
	case 2: // category
		params.BY = "c.name"
	case 3: // item code
		params.BY = "p.id"
	}

	var searchBy string
	fmt.Print("Enter search keyword(s): ")
	fmt.Scan(&searchBy)

	keywords := parseKeywords(searchBy)
	params.Filter = buildSearchQuery(params.BY, keywords)
	results, err := repository.GetProductsPaginated(db, params)
	if err != "" {
		utils.ErrorMessage(err)
		return false
	}

	utils.DisplayDataJson(model.Response{StatusCode: 200, Message: "Read Success", Data: results})

	return true
}

func parseKeywords(input string) []string {
	words := strings.Fields(input) // Pisahkan input berdasarkan spasi
	return words
}

func buildSearchQuery(column string, keywords []string) string {
	var conditions []string
	for _, keyword := range keywords {
		if !strings.HasPrefix(keyword, "pd") {
			keyword = utils.Capitalize(keyword)
		}
		conditions = append(conditions, fmt.Sprintf("%s LIKE '%%%s%%'", column, keyword))
	}
	return strings.Join(conditions, " OR ") // Menggabungkan kondisi dengan OR untuk pencarian fleksibel
}

func ManageTrxHistory(db *sql.DB) bool {
	index := utils.ChoseMenu(view.MenuTransaction)
	if index == 0 {
		return false
	}

	params := repository.CRUDParams{DB: db}

	params.TableName = "transactions"
	params.Page = 1
	params.Limit = 5
	params.BY = "transaction_type"

	switch index {
	case 1:
		params.Filter = fmt.Sprintf("%s = '%s'", params.BY, "in")
	case 2:
		params.Filter = fmt.Sprintf("%s = '%s'", params.BY, "out")
	}

	results, err := repository.GetProductsPaginated(db, params)
	if err != "" {
		utils.ErrorMessage(err)
		return false
	}
	utils.DisplayDataJson(model.Response{StatusCode: 200, Message: "Read Success", Data: results})

	return true
}

func CrudChose(index int, db *sql.DB, params repository.CRUDParams) bool {
	switch index {
	case 1: // Create
		params.IsUpdate = false
		if err := repository.UpsertFromJSON(params); err != "" {
			utils.ErrorMessage(fmt.Sprintf("Error inserting data from JSON: %v", err))
			return false
		}
		utils.DisplayDataJson(model.Response{StatusCode: 201, Message: "Create Success"})

	case 2: // Read
		results, err := repository.GetProductsPaginated(db, params)
		if err != "" {
			utils.ErrorMessage(err)
			return false
		}
		utils.DisplayDataJson(model.Response{StatusCode: 200, Message: "Read Success", Data: results})

	case 3: // Update
		params.IsUpdate = true
		if err := repository.UpsertFromJSON(params); err != "" {
			utils.ErrorMessage(fmt.Sprintf("Error updating data from JSON: %v", err))
			return false
		}
		utils.DisplayDataJson(model.Response{StatusCode: 201, Message: "Update Success"})

	case 4: // Delete
		params.Condition = fmt.Sprintf("%s = '%s'", params.IDColumn, params.KeyDelete)
		if err := repository.DeleteData(params); err != "" {
			utils.ErrorMessage(err)
			return false
		}
		utils.DisplayDataJson(model.Response{StatusCode: 200, Message: "Delete Success"})

	case 5: // Exit
		return false
	}

	return true
}
