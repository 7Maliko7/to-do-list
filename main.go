package main

import (
	"encoding/json"
	"github.com/7Maliko7/to-do-list/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	err := readTaskFile()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/create", handler.CreateHandler)
	http.HandleFunc("/list", handler.ListHandler)
	http.HandleFunc("/", handler.GetHandler)
	http.HandleFunc("/delete", handler.DeleteHandler)
	http.HandleFunc("/update", handler.UpdateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readTaskFile() error {
	file, err := os.OpenFile(handler.TaskFileName, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	err = json.NewDecoder(file).Decode(&handler.TaskList)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
