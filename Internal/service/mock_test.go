package service_test

import (
	"TaskAPI/Internal/models"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetTasks() ([]models.Task, error) {
	args := m.Called()
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTaskById(id int) (models.Task, error) {
	args := m.Called(id)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockTaskRepository) CreateTask(task models.Task) (models.Task, error) {
	args := m.Called(task)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(task models.Task, id int) (models.Task, error) {
	args := m.Called(task)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockTaskRepository) DeleteTask(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
