package task

import "time"

type Task struct {
	Uuid     string    `json:"uuid"`
	Name     string    `json:"name"`
	Body     string    `json:"body"`
	Status   string    `json:"status"`
	Deadline time.Time `json:"deadline"`
}

type TaskCollection struct {
	List []Task
}

func (tc *TaskCollection) DeleteTask(uuid string) {

}