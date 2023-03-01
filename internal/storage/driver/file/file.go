package file

import (
	"errors"
	"encoding/json"
	"github.com/7Maliko7/to-do-list/internal/storage/structs"
	"os"
)

const (
	errUnimplemented = "unimplemented method"
)

type FileStorage struct{
	fileName string
}

func New (fileName string)FileStorage{
	return FileStorage{}
}

func CreateTask(req structs.CreateTaskRequest) error {
	return errors.New(errUnimplemented)
}
func DeleteTask(req structs.DeleteTaskRequest) error {
	return errors.New(errUnimplemented)
}
func GetTask(req structs.GetTaskRequest) (structs.GetTaskResponse, error) {
	return structs.GetTaskResponse{}, errors.New(errUnimplemented)
}
func GetListTask(req structs.GetListTaskRequest) (structs.GetListTaskResponse, error) {
	return structs.GetListTaskResponse{}, errors.New(errUnimplemented)
}
func UpdateTask(req structs.UpdateTaskRequest) error {
	return errors.New(errUnimplemented)
}

func (fs *FileStorage)writeFile(data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(, bytes, 0777)
	if err != nil {
		return err
	}
	return nil
}