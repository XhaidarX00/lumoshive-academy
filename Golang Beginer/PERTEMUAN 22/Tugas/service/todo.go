package service

import "main/model"

func (s *Service) AddTodoService(task *model.Task) (int, error) {
	return s.Repo.AddTodoRepo(task)
}

func (s *Service) GetTodosService(tasks *[]model.Task) {
	s.Repo.GetTodosRepo(tasks)
}

func (s *Service) ChangeTodoStatusService(id int) error {
	return s.Repo.ChangeTodoStatusRepo(id)
}
