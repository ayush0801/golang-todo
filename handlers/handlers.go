package handlers

import (
	"database/sql"
	"encoding/json"
	"go-todo-postgres/models"
	"log"
	"net/http"
)

func getTodos(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM todos")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		todos := []models.Todo{}

		for rows.Next() {
			var temp models.Todo
			err := rows.Scan(&temp.ID, temp.Item, temp.Done)
			if err != nil {
				log.Fatal(err)
			}
			todos = append(todos, temp)
		}
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(todos)
	}
}
