package core

import (
	"github.com/7Maliko7/to-do-list/internal/storage"
	"time"
)

type Core struct {
	storage storage.Storager
}

type TaskList []Task

type Task struct {
	Uuid     string
	Name     string
	Body     string
	Status   string
	Deadline time.Time
}

func New(storage storage.Storager) Core {
	return Core{storage: storage}
}

func (c *Core) GetTaskList() TaskList {
	tasks := c.storage.GetListTask()
}
