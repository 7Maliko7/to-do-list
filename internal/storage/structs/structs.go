package structs

import "time"

type GetListTaskResponse struct {
	List []Task   `json:"list"`
}

type GetTaskResponse struct {
	Task
}
type Task struct {
	Uuid     string    `json:"uuid"`
	Name     string    `json:"name"`
	Body     string    `json:"body"`
	Status   string    `json:"status"`
	Deadline time.Time `json:"deadline"`
}

type CreateTaskRequest struct {
	Task
}

type DeleteTaskRequest struct {
	Uuid     string    `json:"uuid"`
}

type UpdateTaskRequest struct {
	Task
}

type GetTaskRequest struct {
	Uuid     string    `json:"uuid"`
}
type GetListTaskRequest struct {

}

