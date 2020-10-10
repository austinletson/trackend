package world

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqlLiteDB struct {
	db *sql.DB
}

func NewSqlLiteDB() (*SqlLiteDB, error) {
	db, err := sql.Open("sqlite3", "./trackend.db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &SqlLiteDB{db}, nil
}

func CloseDb(data SqlLiteDB) {
	data.db.Close()
}
