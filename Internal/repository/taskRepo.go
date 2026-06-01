package repository

import (
	"TaskAPI/Internal/models"

	"github.com/jmoiron/sqlx"
)

type TaskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Select(&tasks, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepo) GetTaskById(id int) (models.Task, error) {
	var task models.Task
	err := r.db.Get(&task, "SELECT * FROM tasks WHERE id = $1", id)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepo) CreateTask(task models.Task) (models.Task, error) {
	_, err := r.db.Exec(`INSERT INTO tasks(title,descr,status,created_at,updated_at) VALUES($1,$2,$3,$4,$5)`, task.Title, task.Desc, task.Status, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepo) UpdateTask(task models.Task, id int) (models.Task, error) {
	_, err := r.db.Exec(`UPDATE tasks SET title=$1,descr=$2,status=$3,updated_at=$4 WHERE id = $5`, task.Title, task.Desc, task.Status, task.UpdatedAt, id)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepo) DeleteTask(id int) error {
	_, err := r.db.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
