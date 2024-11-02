package service

import (
	"database/sql"
	"elearning/model"
	"elearning/repository"
	"elearning/utils"
	"elearning/view"
	"fmt"
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
	return &response, userReslut.Role, nil
}

func RoleExcAdmin(db *sql.DB, userService *UserService, index int) bool {
	params := repository.CRUDParams{DB: db}

	// Tampilkan menu dan pilih indeks untuk CRUD
	switch index {
	case 1:
		params.TableName = "mentor"
		params.IDColumn = "user_id"
		params.Prefix = "mnt"
		params.KeyUpdate = "mnt1"
		params.KeyDelete = "mnt3"
		params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/newMentor.json"
		params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/updateMentor.json"
		params.Columns = []string{"user_id", "name", "degree", "experience"}

	case 2:
		params.TableName = "student"
		params.IDColumn = "user_id"
		params.Prefix = "std"
		params.KeyUpdate = "std1"
		params.KeyDelete = "std2"
		params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/newStudent.json"
		params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/updateStudent.json"
		params.Columns = []string{"user_id", "name", "phone_number", "address"}

	case 3:
		params.TableName = "materials"
		params.IDColumn = "material_id"
		params.Prefix = "mat"
		params.KeyUpdate = "mat2"
		params.KeyDelete = "mat3"
		params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/newMaterial.json"
		params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/updateMaterial.json"
		params.Columns = []string{"material_id", "title", "description", "video_url", "mentor_id"}

	case 4:
		params.TableName = "Class_Schedule"
		params.IDColumn = "schedule_id"
		params.Prefix = "sch"
		params.KeyUpdate = "sch2"
		params.KeyDelete = "sch3"
		params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/newSchedule.json"
		params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/updateSchedule.json"
		params.Columns = []string{"schedule_id", "class_id", "date", "start_time", "end_time"}

	case 5:
		params.TableName = "Announcements"
		params.IDColumn = "announcement_id"
		params.Prefix = "anc"
		params.KeyUpdate = "anc2"
		params.KeyDelete = "anc3"
		params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/newAnnouncement.json"
		params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/updateAnnouncement.json"
		params.Columns = []string{"announcement_id", "title", "content", "date"}

	default:
		utils.ErrorMessage("menu tidak ditemukan!")
		return false
	}

	chosenIndex := utils.ChoseMenuCrud(view.MenuCrud, strings.ToUpper(params.TableName))
	if chosenIndex == 0 {
		return false
	}

	if !CrudChose(chosenIndex, db, userService, params) {
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
		results, err := repository.ReadData(params)
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

func RoleExcMentor(db *sql.DB, userService *UserService, index int) bool {
	params := repository.CRUDParams{DB: db}

	// Tampilkan menu dan pilih indeks untuk CRUD
	switch index {
	case 1:
		params.TableName = "Class_Schedule"
		params.Columns = []string{"schedule_id", "class_id", "date", "start_time", "end_time"}

	case 2:
		params.TableName = "materials"
		params.Columns = []string{"material_id", "title", "description", "video_url", "mentor_id"}

	case 3:
		params.TableName = "Attendance"
		params.Columns = []string{"attendance_id", "user_id", "schedule_id", "status"}
		params.Condition = "user_id LIKE 'mnt%'"

	case 4:
		params.TableName = "Attendance"
		params.Columns = []string{"attendance_id", "user_id", "schedule_id", "status"}
		params.Condition = "user_id LIKE 'std%'"

	case 5:
		params.TableName = "Grades"
		params.Columns = []string{"grade_id", "user_id", "assignment_id", "grade"}

	case 6:
		params.TableName = "Assignments"
		params.IDColumn = "assignment_id"
		params.Prefix = "asg"
		params.KeyUpdate = "asg2"
		params.KeyDelete = "asg3"
		params.JSONPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/newAssignment.json"
		params.JSONUPDATEPath = "C:/Belajar Coding/lumoshive-academy/Golang Beginer/Testing/Elearning/file_json/updateAssignment.json"
		params.Columns = []string{"assignment_id", "class_id", "title", "description", "deadline"}

	default:
		utils.ErrorMessage("menu tidak ditemukan!")
		return false
	}

	if index == 6 {
		chosenIndex := utils.ChoseMenuCrud(view.MenuCrud, strings.ToUpper(params.TableName))
		if chosenIndex == 0 {
			return false
		}

		if !CrudChose(chosenIndex, db, userService, params) {
			return false
		}

	} else {
		results, err := repository.ReadData(params)
		if err != "" {
			utils.ErrorMessage(err)
			return false
		}
		utils.DisplayDataJson(model.Response{StatusCode: 200, Message: "Read Success", Data: results})
	}

	return true
}

func RoleExcStudent(db *sql.DB, userService *UserService, index int) bool {
	params := repository.CRUDParams{DB: db}

	// Tampilkan menu dan pilih indeks untuk CRUD
	switch index {
	case 1:
		params.TableName = "Class_Schedule"
		params.Columns = []string{"schedule_id", "class_id", "date", "start_time", "end_time"}

	case 2:
		params.TableName = "materials"
		params.Columns = []string{"material_id", "title", "description", "video_url", "mentor_id"}

	case 3:
		params.TableName = "Attendance"
		params.Columns = []string{"attendance_id", "user_id", "schedule_id", "status"}
		params.Condition = "user_id LIKE 'std%'"

	case 4:
		params.TableName = "Grades"
		params.Columns = []string{"grade_id", "user_id", "assignment_id", "grade"}

	case 5:
		params.TableName = "Assignments"
		params.Columns = []string{"assignment_id", "class_id", "title", "description", "deadline"}

	default:
		utils.ErrorMessage("menu tidak ditemukan!")
		return false
	}

	results, err := repository.ReadData(params)
	if err != "" {
		utils.ErrorMessage(err)
		return false
	}
	utils.DisplayDataJson(model.Response{StatusCode: 200, Message: "Read Success", Data: results})

	return true
}
