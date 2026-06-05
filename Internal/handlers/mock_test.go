package handlers_test

import (
	"TaskAPI/Internal/models"

	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) GetTasks() ([]models.Task, error) {
	args := m.Called()
	return args.Get(0).([]models.Task), args.Error(1)

}

func (m *MockTaskService) GetTaskById(id int) (models.Task, error) {
	args := m.Called(id)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockTaskService) CreateTask(task models.Task) (models.Task, error) {
	args := m.Called(task)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockTaskService) UpdateTask(task models.Task, id int) (models.Task, error) {
	args := m.Called(task, id)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockTaskService) DeleteTask(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
