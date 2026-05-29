package db

import (
	"TaskAPI/Internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(c config.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=%s", c.DB.Port, c.DB.Host, c.DB.User, c.Password, c.DB.Name, c.DB.SSLMod)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
