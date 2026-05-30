package models

import "time"

type Task struct {
	Id        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Desc      string    `json:"desc" db:"desc"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
