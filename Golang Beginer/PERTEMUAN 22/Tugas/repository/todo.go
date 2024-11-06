package repository

import (
	todosModel "main/model/todos"
)

func (r *Repository) AddTodoRepo(task *todosModel.Task) (int, error) {
	query := `INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING task_id`
	var id int

	err := r.DB.QueryRow(query, task.Title, task.Description).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) readTaskStatusChange(id int) string {
	var status string
	query := `SELECT status FROM tasks WHERE task_id = $1`
	err := r.DB.QueryRow(query, id).Scan(&status)
	if err != nil {
		panic(err)
	}
	return status
}

func (r *Repository) ChangeTodoStatusRepo(id int) error {
	Status := r.readTaskStatusChange(id)

	if Status == "pending" {
		Status = "completed"
	} else {
		Status = "pending"
	}

	query := `UPDATE tasks SET status = $1 WHERE task_id = $2`
	_, err := r.DB.Exec(query, Status, id)
	if err != nil {
		return err
	}

	return nil
}
