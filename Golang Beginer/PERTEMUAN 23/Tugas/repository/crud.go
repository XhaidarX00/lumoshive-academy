package repository

import (
	"fmt"
	todosModel "main/model/todos"
)

func (r *Repository) CreateTaskRepo(task *todosModel.Task) error {
	query := `INSERT INTO tasks (user_id, title, description) VALUES ($1, $2, $3)`
	_, err := r.DB.Exec(query, task.User_id, task.Title, task.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ReadTaskRepo(tasks *[]todosModel.Task) error {
	query := `SELECT task_id, title, description, status FROM tasks`
	rows, err := r.DB.Query(query)
	if err != nil {
		return fmt.Errorf("error executing read query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		task := todosModel.Task{}
		if err := rows.Scan(&task.Task_id, &task.User_id, &task.Title, &task.Description, &task.Status); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		*tasks = append(*tasks, task)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error reading rows: %v", err)
	}

	return nil
}

func (r *Repository) readTaskStatus(task *todosModel.Task) error {
	query := `SELECT status FROM tasks WHERE task_id = $1`
	err := r.DB.QueryRow(query, task.Task_id).Scan(&task.Status)
	fmt.Println(task.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateTaskRepo(task *todosModel.Task) error {
	ErrorMessage := r.readTaskStatus(task)
	if ErrorMessage != nil {
		return ErrorMessage
	}
	if task.Status == "pending" {
		task.Status = "completed"
	} else {
		task.Status = "pending"
	}

	query := `UPDATE tasks SET status = $1 WHERE task_id = $2`
	_, err := r.DB.Exec(query, task.Status, task.Task_id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteTaskRepo(task_id int) error {
	query := "DELETE FROM tasks WHERE task_id = $1"
	_, err := r.DB.Exec(query, task_id)
	if err != nil {
		return err
	}
	return nil
}
