package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Worker struct {
	DB *sql.DB
}

func CreateWorker(connection string) (*Worker, error) {
	DB, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	return &Worker{DB: DB}, nil
}