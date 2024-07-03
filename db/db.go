package db

import (
	"database/sql"
)

func Connect(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXIST todos(id SERIAL PRIMARY KEY, title TEXT, done BOOLEAN)")
	if err != nil {
		return nil, err
	}

	return db, nil
}
