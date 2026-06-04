package main

import (
	"TaskAPI/Internal/config"
	"TaskAPI/Internal/db"
	"TaskAPI/Internal/handlers"
	"TaskAPI/Internal/repository"
	"TaskAPI/Internal/service"
	"fmt"
	"net/http"
)

func main() {
	//Load config
	cfg := config.MustLoad()
	database, err := db.Connect(*cfg)
	if err != nil {
		panic(err)
	}

	//Create repo,services and handlers
	taskRepo := repository.NewTaskRepo(database)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	//Routes
	http.HandleFunc("GET /tasks", taskHandler.GetTasks)
	http.HandleFunc("GET /tasks/{id}", taskHandler.GetTaskById)
	http.HandleFunc("POST /tasks", taskHandler.CreateTask)
	http.HandleFunc("PUT /tasks/{id}", taskHandler.UpdateTask)
	http.HandleFunc("DELETE /tasks/{id}", taskHandler.DeleteTask)

	//Starting server
	fmt.Println("Сервер запущен на порту 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
