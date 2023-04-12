package db

import (
	"context"
	"errors"
	"github.com/7Maliko7/to-do-list/internal/storage/structs"
	"github.com/jackc/pgx/v5"
	"time"
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
	tag, err := fs.conn.Exec(fs.ctx, "insert into task.list (uuid, name,body,status,deadline) values ($1,$2,$3,$4,$5);", req.Uuid, req.Name, req.Body, req.Status, req.Deadline)
	if err != nil {
		return err
	}
	if tag.RowsAffected() > 0 {
		return nil
	}
	return errors.New("Not Created")
}
func (fs *DbStorage) DeleteTask(req structs.DeleteTaskRequest) error {
	tag, err := fs.conn.Exec(fs.ctx, "delete from task.list where uuid=$1", req.Uuid)
	if err != nil {
		return err
	}
	if tag.RowsAffected() > 0 {
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
	var rowCount uint64
	err := fs.conn.QueryRow(fs.ctx, "select count(Uuid) from task.list").Scan(&rowCount)
	if err != nil {
		return structs.GetListTaskResponse{}, err
	}

	taskList := structs.GetListTaskResponse{}
	rows, err := fs.conn.Query(fs.ctx, "select * from task.list")
	if err != nil {
		return structs.GetListTaskResponse{}, err
	}
	taskList.List = make([]structs.Task, 0, rowCount)
	for rows.Next() {
		var Uuid string
		var Name string
		var Body string
		var Status bool
		var Deadline time.Time
		err := rows.Scan(&Uuid, &Name, &Body, &Status, &Deadline)
		if err != nil {
			return structs.GetListTaskResponse{}, err
		}
		taskList.List = append(taskList.List, structs.Task{
			Uuid:     Uuid,
			Name:     Name,
			Body:     Body,
			Status:   Status,
			Deadline: Deadline,
		})
	}
	return taskList, nil

}

func (fs *DbStorage) UpdateTask(req structs.UpdateTaskRequest) error {
	tag, err := fs.conn.Exec(fs.ctx, "update task.list set (name,body,status,deadline) = ($1,$2,$3,$4) where uuid=$5", req.Name, req.Body, req.Status, req.Deadline, req.Uuid)
	if err != nil {
		return err
	}
	if tag.RowsAffected() > 0 {
		return nil
	}
	return errors.New("Not Updated")
}
