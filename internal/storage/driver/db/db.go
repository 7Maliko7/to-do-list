package db

import (
	"context"
	"errors"
	"github.com/7Maliko7/to-do-list/internal/storage/structs"
	"github.com/jackc/pgx/v5"
)

const (
	errUnimplemented = "unimplemented method"
)

type DbStorage struct {
	conn *pgx.Conn
	ctx  context.Context
}

func (fs *DbStorage) Close(ctx context.Context) error {
	err := fs.conn.Close(ctx)
	if err != nil {
		return err
	}
	return nil
}

func New(ctx context.Context, url string) (DbStorage, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return DbStorage{}, err
	}

	return DbStorage{conn, ctx}, nil
}

func (fs *DbStorage) CreateTask(req structs.CreateTaskRequest) error {
	var ok bool
	err := fs.conn.QueryRow(fs.ctx, "insert into task.list (uuid, name,body,status,deadline) values ($1,$2,$3,$4,$5);", req.Uuid, req.Name, req.Body, req.Status, req.Deadline).Scan(&ok)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	return errors.New("Not Inserted")
}
func (fs *DbStorage) DeleteTask(req structs.DeleteTaskRequest) error {
	var ok bool
	err := fs.conn.QueryRow(fs.ctx, "delete from task.list where uuid=$1", req.Uuid).Scan(&ok)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	return errors.New("Not Deleted")
}
func (fs *DbStorage) GetTask(req structs.GetTaskRequest) (structs.GetTaskResponse, error) {
	task := structs.GetTaskResponse{}
	err := fs.conn.QueryRow(fs.ctx, "select uuid,name, body, status, deadline from task.list where uuid=$1", req.Uuid).Scan(&task.Uuid, &task.Name, &task.Body, &task.Status, &task.Deadline)
	if err != nil {
		return structs.GetTaskResponse{}, err
	}
	return task, nil
}
func (fs *DbStorage) GetListTask(req structs.GetListTaskRequest) (structs.GetListTaskResponse, error) {
	taskList := structs.GetListTaskResponse{}
	err := fs.conn.QueryRow(fs.ctx, "select * from task.list").Scan(&taskList.List)
	if err != nil {
		return structs.GetListTaskResponse{}, err
	}
	return taskList, nil
}

func (fs *DbStorage) UpdateTask(req structs.UpdateTaskRequest) error {
	var ok bool
	err := fs.conn.QueryRow(fs.ctx, "update task.list set (name,body,status,deadline) = ($2,$3,$4,$5) where uuid=$1;", req.Uuid, req.Name, req.Body, req.Status, req.Deadline).Scan(&ok)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	return errors.New("Not Updated")
}

func (fs *DbStorage) writeFile(data interface{}) error {

	return nil
}

func (fs *DbStorage) readFile(data interface{}) error {

	return nil
}
