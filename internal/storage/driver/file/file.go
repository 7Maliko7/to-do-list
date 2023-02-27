package file

import (
	"encoding/json"
	"github.com/7Maliko7/to-do-list/internal/storage/structs"
	"github.com/7Maliko7/to-do-list/task"
	"os"
)
type FileStorage struct{
	fileName string
}


func New (fileName string)FileStorage{

}

func CreateTask(req structs.CreateTaskRequest) error {

}
func DeleteTask(req structs.DeleteTaskRequest) error {

}
func GetTask(req structs.GetTaskRequest) (structs.GetTaskResponse, error) {

}
func GetListTask(req structs.GetListTaskRequest) (structs.GetListTaskResponse, error) {

}
func UpdateTask(req structs.UpdateTaskRequest) error {

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