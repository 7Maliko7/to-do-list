package structs

import "time"

type GetListTaskResponse struct {
	List []Task `json:"list"`
}

type GetTaskResponse Task

type Task struct {
	Uuid     string    `json:"uuid"`
	Name     string    `json:"name"`
	Body     string    `json:"body"`
	Status   bool      `json:"status"`
	Deadline time.Time `json:"deadline"`
}

type CreateTaskRequest Task

type DeleteTaskRequest struct {
	Uuid string `json:"uuid"`
}

type UpdateTaskRequest Task

type GetTaskRequest struct {
	Uuid string `json:"uuid"`
}
type GetListTaskRequest struct {
}
