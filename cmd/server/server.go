package main

import (
	"context"
	"github.com/7Maliko7/to-do-list/internal/core"
	"github.com/7Maliko7/to-do-list/internal/handler"
	"github.com/7Maliko7/to-do-list/internal/storage/driver/db"
	"log"
	"net/http"
)

func main() {
	//fileStorage := file.New("tasks.json")
	DbStorage, err := db.New(context.Background(), "postgresql://postgres:postgres@localhost:5432/todolist")
	if err != nil {
		log.Fatal(err)
	}
	core := core.New(&DbStorage)
	handler.Core = core
	http.HandleFunc("/create", handler.CreateHandler)
	http.HandleFunc("/list", handler.ListHandler)
	http.HandleFunc("/get", handler.GetHandler)
	http.HandleFunc("/delete", handler.DeleteHandler)
	http.HandleFunc("/update", handler.UpdateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
