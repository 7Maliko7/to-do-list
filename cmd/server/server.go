package main

import (
	"fmt"
	"github.com/7Maliko7/to-do-list/internal/core"
	"github.com/7Maliko7/to-do-list/internal/storage/driver/file"
	"log"
)

var Core core.Core

func main() {
	fileStorage := file.New("tasks.json")
	Core = core.New(&fileStorage)
	taskList, err := core.GetTaskList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(taskList)
}
