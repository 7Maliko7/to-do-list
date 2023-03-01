package main

import (
	"github.com/7Maliko7/to-do-list/internal/core"
	"github.com/7Maliko7/to-do-list/internal/storage/driver/file"
)

func main() {
	fileStorage := file.New("tasks.json")
	core := core.New(&fileStorage)
}
