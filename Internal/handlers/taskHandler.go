package handlers

import (
	"TaskAPI/Internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type TaskService interface {
	GetTasks() ([]models.Task, error)
	GetTaskById(id int) (models.Task, error)
	CreateTask(task models.Task) (models.Task, error)
	UpdateTask(task models.Task, id int) (models.Task, error)
	DeleteTask(id int) error
}

type TaskHandler struct {
	Ts TaskService
}

func NewTaskHandler(ts TaskService) *TaskHandler {
	return &TaskHandler{Ts: ts}
}

func (th *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.Ts.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (th *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("id")
	parsedId, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := th.Ts.GetTaskById(parsedId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (th *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	res, err := th.Ts.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (th *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	url := r.PathValue("id")
	parsedId, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	res, err := th.Ts.UpdateTask(task, parsedId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (th *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("id")
	parsedId, err := strconv.Atoi(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = th.Ts.DeleteTask(parsedId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
