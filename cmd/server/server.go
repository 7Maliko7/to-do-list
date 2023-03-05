package main

import (
	"github.com/7Maliko7/to-do-list/internal/core"
	"github.com/7Maliko7/to-do-list/internal/handler"
	"github.com/7Maliko7/to-do-list/internal/storage/driver/file"
	"log"
	"net/http"
)

func main() {
	fileStorage := file.New("tasks.json")
	core := core.New(&fileStorage)
	handler.Core = core
	http.HandleFunc("/create", handler.CreateHandler)
	http.HandleFunc("/list", handler.ListHandler)
	http.HandleFunc("/get", handler.GetHandler)
	http.HandleFunc("/delete", handler.DeleteHandler)
	http.HandleFunc("/update", handler.UpdateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
