package service

import (
	"TaskAPI/Internal/models"
)

type TaskRepository interface {
	GetTasks() ([]models.Task, error)
	GetTaskById(id int) (models.Task, error)
	CreateTask(task models.Task) (models.Task, error)
	UpdateTask(task models.Task, id int) (models.Task, error)
	DeleteTask(id int) error
}
type TaskService struct {
	Tr TaskRepository
}

func NewTaskService(tr TaskRepository) *TaskService {
	return &TaskService{Tr: tr}
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	res, err := s.Tr.GetTasks()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TaskService) GetTaskById(id int) (models.Task, error) {
	task, err := s.Tr.GetTaskById(id)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	res, err := s.Tr.CreateTask(task)
	if err != nil {
		return models.Task{}, err
	}
	return res, nil
}

func (s *TaskService) UpdateTask(task models.Task, id int) (models.Task, error) {
	res, err := s.Tr.UpdateTask(task, id)
	if err != nil {
		return models.Task{}, err
	}
	return res, nil
}

func (s *TaskService) DeleteTask(id int) error {
	err := s.Tr.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
