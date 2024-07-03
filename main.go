package main

import (
	"go-todo-postgres/routers"
	"log"
	"net/http"
	"os"
	"todo-app/db"
	"todo-app/middleware"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	dbConn, err := db.Connect(dbURL)

	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	r := routers.SetupRouter(dbConn)
	log.Fatal(http.ListenAndServe(":8080", middleware.JSONContentType(r)))
}
