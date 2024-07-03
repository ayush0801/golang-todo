package routers

import (
	"database/sql"s
	"go-todo-postgres/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) *mux.routers {
	router := mux.NewRouter()
	router.Handlefunc("/todos", handlers.getTodos(db)).Methods("GET")
	return router
}
