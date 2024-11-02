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

func ManageStock(db *sql.DB, userService *UserService) bool {
	index := utils.ChoseMenu(view.MenuProduct)
	params := repository.CRUDParams{DB: db}

	switch index {
	case 1:
		params.TableName = "products"
		params.IDColumn = "id"
		params.Prefix = "pd"
		params.KeyUpdate = "pd21"
		params.KeyDelete = "pd22"
		params.Page = 2
		params.Limit = 5
		params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/newProduct.json"
		params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/updateProduct.json"
		params.Columns = []string{"id", "name", "price", "stock", "category_id", "location_id"}

		chosenIndex := utils.ChoseMenuCrud(view.MenuCrud, strings.ToUpper(params.TableName))
		if chosenIndex == 0 {
			return false
		}

		if !CrudChose(chosenIndex, db, userService, params) {
			return false
		}

	case 2:

	default:
		utils.ErrorMessage("menu tidak ditemukan!")
		return false
	}

	return true
}

func CrudChose(index int, db *sql.DB, userService *UserService, params repository.CRUDParams) bool {
	switch index {
	case 1: // Create
		params.IsUpdate = false
		if err := repository.UpsertFromJSON(params); err != "" {
			utils.ErrorMessage(fmt.Sprintf("Error inserting data from JSON: %v", err))
			return false
		}
		utils.DisplayDataJson(model.Response{StatusCode: 201, Message: "Create Success"})

	case 2: // Read
		if params.TableName == "products" {
			results, err := repository.GetProductsPaginated(db, params.Page, params.Limit)
			if err != nil {
				utils.ErrorMessage(err.Error())
				return false
			}
			utils.DisplayDataJson(model.Response{StatusCode: 200, Message: "Read Success", Data: results})
		} else {
			results, err := repository.ReadData(params)
			if err != "" {
				utils.ErrorMessage(err)
				return false
			}
			utils.DisplayDataJson(model.Response{StatusCode: 200, Message: "Read Success", Data: results})
		}

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

// func RoleExcMentor(db *sql.DB, userService *UserService, index int) bool {
// 	params := repository.CRUDParams{DB: db}

// 	// Tampilkan menu dan pilih indeks untuk CRUD
// 	switch index {
// 	case 1:
// 		params.TableName = "Class_Schedule"
// 		params.Columns = []string{"schedule_id", "class_id", "date", "start_time", "end_time"}

// 	case 2:
// 		params.TableName = "materials"
// 		params.Columns = []string{"material_id", "title", "description", "video_url", "mentor_id"}

// 	case 3:
// 		params.TableName = "Attendance"
// 		params.Columns = []string{"attendance_id", "user_id", "schedule_id", "status"}
// 		params.Condition = "user_id LIKE 'mnt%'"

// 	case 4:
// 		params.TableName = "Attendance"
// 		params.Columns = []string{"attendance_id", "user_id", "schedule_id", "status"}
// 		params.Condition = "user_id LIKE 'std%'"

// 	case 5:
// 		params.TableName = "Grades"
// 		params.Columns = []string{"grade_id", "user_id", "assignment_id", "grade"}

// 	case 6:
// 		params.TableName = "Assignments"
// 		params.IDColumn = "assignment_id"
// 		params.Prefix = "asg"
// 		params.KeyUpdate = "asg2"
// 		params.KeyDelete = "asg3"
// 		params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Inventory/file_json/updateProduct.jsonnewAssignment.json"
// 		params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/updateAssignment.json"
// 		params.Columns = []string{"assignment_id", "class_id", "title", "description", "deadline"}

// 	default:
// 		utils.ErrorMessage("menu tidak ditemukan!")
// 		return false
// 	}

// 	if index == 6 {
// 		chosenIndex := utils.ChoseMenuCrud(view.MenuCrud, strings.ToUpper(params.TableName))
// 		if chosenIndex == 0 {
// 			return false
// 		}

// 		if !CrudChose(chosenIndex, db, userService, params) {
// 			return false
// 		}

// 	} else {
// 		results, err := repository.ReadData(params)
// 		if err != "" {
// 			utils.ErrorMessage(err)
// 			return false
// 		}
// 		utils.DisplayDataJson(model.Response{StatusCode: 200, Message: "Read Success", Data: results})
// 	}

// 	return true
// }

// func RoleExcStudent(db *sql.DB, userService *UserService, index int) bool {
// 	params := repository.CRUDParams{DB: db}

// 	// Tampilkan menu dan pilih indeks untuk CRUD
// 	switch index {
// 	case 1:
// 		params.TableName = "Class_Schedule"
// 		params.Columns = []string{"schedule_id", "class_id", "date", "start_time", "end_time"}

// 	case 2:
// 		params.TableName = "materials"
// 		params.Columns = []string{"material_id", "title", "description", "video_url", "mentor_id"}

// 	case 3:
// 		params.TableName = "Attendance"
// 		params.Columns = []string{"attendance_id", "user_id", "schedule_id", "status"}
// 		params.Condition = "user_id LIKE 'std%'"

// 	case 4:
// 		params.TableName = "Grades"
// 		params.Columns = []string{"grade_id", "user_id", "assignment_id", "grade"}

// 	case 5:
// 		params.TableName = "Assignments"
// 		params.Columns = []string{"assignment_id", "class_id", "title", "description", "deadline"}

// 	default:
// 		utils.ErrorMessage("menu tidak ditemukan!")
// 		return false
// 	}

// 	results, err := repository.ReadData(params)
// 	if err != "" {
// 		utils.ErrorMessage(err)
// 		return false
// 	}
// 	utils.DisplayDataJson(model.Response{StatusCode: 200, Message: "Read Success", Data: results})

// 	return true
// }
