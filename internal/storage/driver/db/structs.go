package db

import "time"

type TaskList []Task

type Task struct {
	Uuid     string    `json:"uuid"`
	Name     string    `json:"name"`
	Body     string    `json:"body"`
	Status   bool      `json:"status"`
	Deadline time.Time `json:"deadline"`
}
