package handler

import (
	"encoding/json"
	"fmt"
	"github.com/7Maliko7/to-do-list/structs"
	"github.com/7Maliko7/to-do-list/task"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
)

const (
	headerContentType = "Content-Type"
	contentTypeJson   = "application/json"
	taskFileName = "tasks.json"
)

var TaskList []task.Task

func makeResponse(w http.ResponseWriter, data any) error {
	w.Header().Set(headerContentType, contentTypeJson)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func makeErrorResponse(w http.ResponseWriter, e structs.ErrResponse){
	msg,_ := json.Marshal(e)
	http.Error(w, string(msg), e.Code)
}

func writeFile(taskList []task.Task) error {
	bytes, err := json.Marshal(taskList)
	if err != nil {
		return err
	}

	err = os.WriteFile(taskFileName, bytes, 0777)
	if err != nil {
		return err
	}
	return nil
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == http.MethodPost {

		var newTask task.Task
		err = json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil{
			e:= structs.ErrResponse{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
		}
		newTask.Uuid = uuid.New().String()
		TaskList = append(TaskList, newTask)

		err = writeFile(TaskList)
		if err != nil {
			e:= structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
		}
		err = makeResponse(w, newTask)
		if err != nil {
			log.Print(err)
			return
		}
		return
	}
	http.Error(w, fmt.Sprintf("expect method POST at /create, got %v", r.Method), http.StatusMethodNotAllowed)
	return

}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(TaskList)
		return
	}
	http.Error(w, fmt.Sprintf("expect method GET at /create, got %v", r.Method), http.StatusMethodNotAllowed)
	return

}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		var ReadTask Task
		_ = json.NewDecoder(r.Body).Decode(&ReadTask)
		for _, i := range TaskList {
			if i.Uuid == ReadTask.Uuid {
				json.NewEncoder(w).Encode(i)
				return
			}
		}
		http.Error(w, fmt.Sprintf("not found"), http.StatusMethodNotAllowed)
		return

	}
	http.Error(w, fmt.Sprintf("expect method GET at /create, got %v", r.Method), http.StatusMethodNotAllowed)
	return
}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		w.Header().Set("Content-Type", "application/json")
		var DeleteTask Task
		_ = json.NewDecoder(r.Body).Decode(&DeleteTask)
		for i, v := range TaskList {
			if v.Uuid == DeleteTask.Uuid {
				TaskList = append(TaskList[:i], TaskList[i+1:]...)
				json.NewEncoder(w).Encode(TaskList)
				err := writeFile(TaskList)
				if err != nil {
					log.Print(err)
				}
				return
			}
		}
		http.Error(w, fmt.Sprint("element not found"), http.StatusMethodNotAllowed)
		return
	}
	http.Error(w, fmt.Sprintf("expect method DELETE at /delete, got %v", r.Method), http.StatusMethodNotAllowed)
	return
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		w.Header().Set("Content-Type", "application/json")
		var PatchTask Task
		_ = json.NewDecoder(r.Body).Decode(&PatchTask)
		for i, _ := range TaskList {
			if TaskList[i].Uuid == PatchTask.Uuid {
				TaskList[i].Name = PatchTask.Name
				TaskList[i].Body = PatchTask.Body
				TaskList[i].Deadline = PatchTask.Deadline
				TaskList[i].Status = PatchTask.Status
				json.NewEncoder(w).Encode(TaskList[i])
				err := writeFile(TaskList)
				if err != nil {
					log.Print(err)
				}
				return
			}
		}
		http.Error(w, fmt.Sprint("element not found"), http.StatusMethodNotAllowed)
		return
	}
	http.Error(w, fmt.Sprintf("expect method PATCH at /delete, got %v", r.Method), http.StatusMethodNotAllowed)
	return
}
