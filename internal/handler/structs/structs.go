package structs

import "time"

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type OkResponse struct {
	Code int `json:"code"`
}

type CreateTaskRequest struct {
	Name     string    `json:"name"`
	Body     string    `json:"body"`
	Deadline time.Time `json:"deadline"`
}

type CreateTaskResponse struct {
	Uuid string `json:"uuid"`
}

type ListTaskRequest struct {
}

type ListTaskResponse struct {
	List []Task `json:"list"`
}

type Task struct {
	Uuid     string    `json:"uuid"`
	Name     string    `json:"name"`
	Body     string    `json:"body"`
	Status   string    `json:"status"`
	Deadline time.Time `json:"deadline"`
}

type GetTaskRequest struct {
	Uuid string `json:"uuid"`
}

type GetTaskResponse Task

type DeleteTaskRequest struct {
	Uuid string `json:"uuid"`
}

type DeleteTaskResponse OkResponse

type UpdateTaskRequest Task

type UpdateTaskResponse OkResponse
