package repository

import (
	"fmt"
	"main/model"
)

func (r *Repository) CreateTaskRepo(task *model.Task) error {
	query := `INSERT INTO tasks (user_id, title, description) VALUES ($1, $2, $3)`
	_, err := r.DB.Exec(query, task.User_id, task.Title, task.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ReadTaskRepo(tasks *[]model.Task) error {
	query := `SELECT task_id, title, description, status FROM tasks`
	rows, err := r.DB.Query(query)
	if err != nil {
		return fmt.Errorf("error executing read query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		task := model.Task{}
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		*tasks = append(*tasks, task)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error reading rows: %v", err)
	}

	return nil
}

func (r *Repository) readTaskStatus(task *model.Task) error {
	query := `SELECT status FROM tasks WHERE task_id = $1`
	err := r.DB.QueryRow(query, task.ID).Scan(&task.Status)
	fmt.Println(task.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateTaskRepo(task *model.Task) error {
	ErrorMessage := r.readTaskStatus(task)
	if ErrorMessage != nil {
		return ErrorMessage
	}
	if task.Status == "pending" {
		task.Status = "completed"
	} else {
		task.Status = "pending"
	}

	fmt.Printf("Status : %v\n", task.Status)
	query := `UPDATE tasks SET user_id = $1, title = $2, description = $3, status = $4 WHERE task_id = $5`
	_, err := r.DB.Exec(query, task.User_id, task.Title, task.Description, task.Status, task.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteTaskRepo(task *model.Task) error {
	query := "DELETE FROM tasks WHERE task_id = $1"
	_, err := r.DB.Exec(query, task.ID)
	if err != nil {
		return err
	}
	return nil
}
