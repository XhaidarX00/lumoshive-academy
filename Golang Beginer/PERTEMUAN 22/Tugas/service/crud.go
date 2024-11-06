package service

import "main/model"

func (s *Service) CreateTaskService(task *model.Task) error {
	err := s.Repo.CreateTaskRepo(task)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReadTaskService(task *[]model.Task) error {
	err := s.Repo.ReadTaskRepo(task)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateTaskService(task *model.Task) error {
	err := s.Repo.UpdateTaskRepo(task)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteTaskService(task_id int) error {
	err := s.Repo.DeleteTaskRepo(task_id)
	if err != nil {
		return err
	}
	return nil
}
