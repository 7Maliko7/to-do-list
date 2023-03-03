package storage

import "github.com/7Maliko7/to-do-list/internal/storage/structs"

type Storager interface {
	CreateTask(req structs.CreateTaskRequest) error
	DeleteTask(req structs.DeleteTaskRequest) error
	GetTask(req structs.GetTaskRequest) (structs.GetTaskResponse, error)
	GetListTask(req structs.GetListTaskRequest) (structs.GetListTaskResponse, error)
	UpdateTask(req structs.UpdateTaskRequest) error
}
