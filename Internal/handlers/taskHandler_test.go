package handlers_test

import (
	"TaskAPI/Internal/handlers"
	"TaskAPI/Internal/models"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTaskHandler_CreateTask(t *testing.T) {
	//Arrange
	mockTaskService := new(MockTaskService)
	hndlr := handlers.NewTaskHandler(mockTaskService)
	mockTaskService.On("CreateTask", mock.AnythingOfType("models.Task")).Return(models.Task{}, nil)
	w := httptest.NewRecorder()
	body := strings.NewReader(`{}`)
	r := httptest.NewRequest("POST", "/tasks", body)

	//Act
	hndlr.CreateTask(w, r)

	//Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	mockTaskService.AssertCalled(t, "CreateTask", mock.AnythingOfType("models.Task"))
}

func TestTaskHandler_CreateTaskFail(t *testing.T) {
	//Arrange
	mockTaskService := new(MockTaskService)
	hndlr := handlers.NewTaskHandler(mockTaskService)
	mockTaskService.On("CreateTask", mock.AnythingOfType("models.Task")).Return(models.Task{}, errors.New(""))
	w := httptest.NewRecorder()
	body := strings.NewReader(`{}`)
	r := httptest.NewRequest("POST", "/tasks", body)

	//Act
	hndlr.CreateTask(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockTaskService.AssertCalled(t, "CreateTask", mock.AnythingOfType("models.Task"))
}

func Test_GetTask(t *testing.T) {
	//Arrange
	mockTaskService := new(MockTaskService)
	hndlr := handlers.NewTaskHandler(mockTaskService)
	mockTaskService.On("GetTasks").Return([]models.Task{}, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/tasks/1", nil)

	//Act
	hndlr.GetTasks(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockTaskService.AssertCalled(t, "GetTasks")
}

func Test_GetTaskFail(t *testing.T) {
	//Arrange
	mockTaskService := new(MockTaskService)
	hndlr := handlers.NewTaskHandler(mockTaskService)
	mockTaskService.On("GetTasks").Return([]models.Task{}, errors.New("some error"))
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/tasks/1", nil)

	//Act
	hndlr.GetTasks(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockTaskService.AssertCalled(t, "GetTasks")
}

func Test_GetTasksById(t *testing.T) {
	//Arrange
	id := 1
	mockTaskService := new(MockTaskService)
	hndlr := handlers.NewTaskHandler(mockTaskService)
	mockTaskService.On("GetTaskById", id).Return(models.Task{}, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/tasks/1", nil)
	r.SetPathValue("id", "1")

	//Act
	hndlr.GetTaskById(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockTaskService.AssertCalled(t, "GetTaskById", id)

}

func Test_GetTasksByIdFail(t *testing.T) {
	//Arrange
	id := 1
	mockTaskService := new(MockTaskService)
	hndlr := handlers.NewTaskHandler(mockTaskService)
	mockTaskService.On("GetTaskById", id).Return(models.Task{}, errors.New("some error"))
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/tasks/1", nil)
	r.SetPathValue("id", "1")

	//Act
	hndlr.GetTaskById(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockTaskService.AssertCalled(t, "GetTaskById", id)
}

func Test_UpdateTasks(t *testing.T) {
	//Arrange
	id := 1
	mockTaskService := new(MockTaskService)
	hndlr := handlers.NewTaskHandler(mockTaskService)
	mockTaskService.On("UpdateTask", mock.AnythingOfType("models.Task"), id).Return(models.Task{}, nil)
	w := httptest.NewRecorder()
	body := strings.NewReader(`{}`)
	r := httptest.NewRequest("PUT", "/tasks/1", body)
	r.SetPathValue("id", "1")

	//Act
	hndlr.UpdateTask(w, r)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	mockTaskService.AssertCalled(t, "UpdateTask", mock.AnythingOfType("models.Task"), 1)
}

func Test_UpdateTaskFail(t *testing.T) {
	//Arrange
	id := 1
	mockTaskService := new(MockTaskService)
	hndlr := handlers.NewTaskHandler(mockTaskService)
	mockTaskService.On("UpdateTask", mock.AnythingOfType("models.Task"), id).Return(models.Task{}, errors.New(""))
	w := httptest.NewRecorder()
	body := strings.NewReader(`{}`)
	r := httptest.NewRequest("PUT", "/tasks/1", body)
	r.SetPathValue("id", "1")

	//Act
	hndlr.UpdateTask(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockTaskService.AssertCalled(t, "UpdateTask", mock.AnythingOfType("models.Task"), 1)

}

func Test_DeleteTasks(t *testing.T) {
	//Arrange
	id := 1
	mockTaskService := new(MockTaskService)
	hndlr := handlers.NewTaskHandler(mockTaskService)
	mockTaskService.On("DeleteTask", id).Return(nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/tasks/1", nil)
	r.SetPathValue("id", "1")

	//Act
	hndlr.DeleteTask(w, r)

	//Assert
	assert.Equal(t, http.StatusNoContent, w.Code)
	mockTaskService.AssertCalled(t, "DeleteTask", 1)
}

func Test_DeleteTaskFail(t *testing.T) {
	//Arrange
	id := 1
	mockTaskService := new(MockTaskService)
	hndlr := handlers.NewTaskHandler(mockTaskService)
	mockTaskService.On("DeleteTask", id).Return(errors.New("some error"))
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/tasks/1", nil)
	r.SetPathValue("id", "1")

	//Act
	hndlr.DeleteTask(w, r)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockTaskService.AssertCalled(t, "DeleteTask", 1)
}
