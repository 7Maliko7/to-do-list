package core

import (
	"github.com/7Maliko7/to-do-list/internal/storage"
	"github.com/7Maliko7/to-do-list/internal/storage/structs"
	"github.com/google/uuid"
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

func (c *Core) GetTaskList() (TaskList, error) {
	tasks, err := c.storage.GetListTask(structs.GetListTaskRequest{})
	if err != nil {
		return nil, err
	}
	List := make([]Task, 0, len(tasks.List))
	for _, v := range tasks.List {
		List = append(List, Task{
			Uuid:     v.Uuid,
			Name:     v.Name,
			Body:     v.Body,
			Status:   v.Status,
			Deadline: v.Deadline,
		})
	}
	return List, err
}

func (c *Core) CreateTask(Name, Body string, deadline time.Time) (string, error) {
	task := structs.CreateTaskRequest{
		Uuid:     uuid.New().String(),
		Name:     Name,
		Body:     Body,
		Deadline: deadline,
	}
	err := c.storage.CreateTask(task)
	if err != nil {
		return "", err
	}
	return task.Uuid, err
}

func (c *Core) DeleteTask(uuid string) error {
	task := structs.DeleteTaskRequest{
		Uuid: uuid,
	}
	err := c.storage.DeleteTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (c *Core) GetTask(uuid string) (Task, error) {
	req := structs.GetTaskRequest{
		Uuid: uuid,
	}
	task, err := c.storage.GetTask(req)
	if err != nil {
		return Task{}, err
	}
	return Task{
		Uuid:     task.Uuid,
		Name:     task.Name,
		Body:     task.Body,
		Status:   task.Status,
		Deadline: task.Deadline,
	}, nil
}

func (c *Core) UpdateTask(Uuid, Name, Body, Status string, deadline time.Time) error {
	req := structs.UpdateTaskRequest{
		Uuid:     Uuid,
		Name:     Name,
		Body:     Body,
		Status:   Status,
		Deadline: deadline,
	}
	err := c.storage.UpdateTask(req)
	if err != nil {
		return err
	}
	return nil
}
