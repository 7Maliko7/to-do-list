package file

import (
	"context"
	"encoding/json"
	"github.com/7Maliko7/to-do-list/internal/storage/structs"
	"github.com/jackc/pgx/v5"
	"os"
	"time"
)

const (
	errUnimplemented = "unimplemented method"
)

type DbStorage struct {
	conn *pgx.Conn
}

func (fs *DbStorage) Close(ctx context.Context) error {
	err := fs.conn.Close(ctx)
	if err != nil {
		return err
	}
}

func New(ctx context.Context, url string) (DbStorage, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return DbStorage{}, err
	}

	return DbStorage{conn}, nil
}

func (fs *DbStorage) CreateTask(req structs.CreateTaskRequest) error {
	tasks := TaskList{}
	err := fs.readFile(&tasks)
	if err != nil {
		return err
	}
	tasks = append(tasks, Task{
		Uuid:     req.Uuid,
		Name:     req.Name,
		Body:     req.Body,
		Status:   req.Status,
		Deadline: time.Time{},
	})
	err = fs.writeFile(tasks)
	if err != nil {
		return err
	}
	return nil
}
func (fs *DbStorage) DeleteTask(req structs.DeleteTaskRequest) error {
	tasks := TaskList{}
	err := fs.readFile(&tasks)
	if err != nil {
		return err
	}
	for i, v := range tasks {
		if v.Uuid == req.Uuid {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	err = fs.writeFile(tasks)
	if err != nil {
		return err
	}

	return nil
}
func (fs *DbStorage) GetTask(req structs.GetTaskRequest) (structs.GetTaskResponse, error) {
	tasks := TaskList{}
	err := fs.readFile(&tasks)
	if err != nil {
		return structs.GetTaskResponse{}, err
	}
	for _, v := range tasks {
		if v.Uuid == req.Uuid {
			return structs.GetTaskResponse{
				Uuid:     v.Uuid,
				Name:     v.Name,
				Body:     v.Body,
				Status:   v.Status,
				Deadline: v.Deadline,
			}, nil
		}
	}
	return structs.GetTaskResponse{}, nil
}
func (fs *DbStorage) GetListTask(req structs.GetListTaskRequest) (structs.GetListTaskResponse, error) {
	tasks := TaskList{}
	err := fs.readFile(&tasks)
	if err != nil {
		return structs.GetListTaskResponse{}, err
	}
	list := make([]structs.Task, 0, len(tasks))
	for _, v := range tasks {
		list = append(list, structs.Task{
			Uuid:     v.Uuid,
			Name:     v.Name,
			Body:     v.Body,
			Status:   v.Status,
			Deadline: time.Time{},
		})
	}

	return structs.GetListTaskResponse{List: list}, err
}
func (fs *DbStorage) UpdateTask(req structs.UpdateTaskRequest) error {
	tasks := TaskList{}
	err := fs.readFile(&tasks)
	if err != nil {
		return err
	}
	for i, _ := range tasks {
		if tasks[i].Uuid == req.Uuid {
			tasks[i].Name = req.Name
			tasks[i].Body = req.Body
			tasks[i].Deadline = req.Deadline
			tasks[i].Status = req.Status
		}
	}
	err = fs.writeFile(tasks)
	if err != nil {
		return err
	}
	return nil
}

func (fs *DbStorage) writeFile(data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(fs.dbName, bytes, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (fs *DbStorage) readFile(data interface{}) error {
	file, err := os.OpenFile(fs.dbName, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
