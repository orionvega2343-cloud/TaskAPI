package service_test

import (
	"TaskAPI/Internal/models"
	"TaskAPI/Internal/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTaskSuccess(t *testing.T) {
	//Arrange
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	expectedTask := []models.Task{}
	mockRepo.On("GetTasks").Return(expectedTask, nil)

	//Act
	result, err := svc.GetTasks()

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedTask, result)
	mockRepo.AssertExpectations(t)
}

func Test_GetTaskFail(t *testing.T) {
	//Arrange
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	expectedTask := []models.Task{}
	mockRepo.On("GetTasks").Return(expectedTask, errors.New("some error"))

	//Act
	result, err := svc.GetTasks()

	//Assert
	assert.Error(t, err)
	assert.Equal(t, expectedTask, result)
	mockRepo.AssertExpectations(t)
}

func Test_GetTaskById(t *testing.T) {
	//Arrange
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	expectedTask := models.Task{}
	mockRepo.On("GetTaskById", 1).Return(expectedTask, nil)

	//Act
	result, err := svc.GetTaskById(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedTask, result)
	mockRepo.AssertExpectations(t)

}
func Test_GetTaskByIdFail(t *testing.T) {
	//Arrange
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	expectedTask := models.Task{}
	mockRepo.On("GetTaskById", 1).Return(expectedTask, errors.New("some error"))

	//Act
	result, err := svc.GetTaskById(1)
	//Assert
	assert.Error(t, err)
	assert.Equal(t, expectedTask, result)
	mockRepo.AssertExpectations(t)
}

func Test_CreateTask(t *testing.T) {
	//Arrange
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	expectedTask := models.Task{}
	mockRepo.On("CreateTask", expectedTask).Return(expectedTask, nil)

	//Act
	result, err := svc.CreateTask(expectedTask)

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedTask, result)
	mockRepo.AssertExpectations(t)
}

func Test_CreateTaskFail(t *testing.T) {
	//Arrange
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	expectedTask := models.Task{}
	mockRepo.On("CreateTask", expectedTask).Return(expectedTask, errors.New("some error"))

	//Act
	result, err := svc.CreateTask(expectedTask)

	//Assert
	assert.Error(t, err)
	assert.Equal(t, expectedTask, result)
	mockRepo.AssertExpectations(t)
}

func Test_UpdateTask(t *testing.T) {
	//Arrange
	id := 1
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	expectedTask := models.Task{}
	mockRepo.On("UpdateTask", expectedTask, id).Return(expectedTask, nil)

	//Act
	result, err := svc.UpdateTask(expectedTask, id)
	//Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedTask, result)
	mockRepo.AssertExpectations(t)

}

func Test_UpdateTaskFail(t *testing.T) {
	//Arrange
	id := 1
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	expectedTask := models.Task{}
	mockRepo.On("UpdateTask", expectedTask, id).Return(expectedTask, errors.New("some error"))

	//Act
	result, err := svc.UpdateTask(expectedTask, id)

	//Assert
	assert.Error(t, err)
	assert.Equal(t, expectedTask, result)
	mockRepo.AssertExpectations(t)

}

func Test_DeleteTask(t *testing.T) {
	//Arrange
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	mockRepo.On("DeleteTask", 1).Return(nil)

	//Act
	err := svc.DeleteTask(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

}

func Test_DeleteTaskFail(t *testing.T) {
	//Arrange
	mockRepo := new(MockTaskRepository)
	svc := service.NewTaskService(mockRepo)
	mockRepo.On("DeleteTask", 1).Return(errors.New("some error"))

	//Act
	err := svc.DeleteTask(1)

	//Assert
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)

}
