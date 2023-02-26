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
	TaskFileName      = "tasks.json"
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

func makeErrorResponse(w http.ResponseWriter, e structs.ErrResponse) {
	msg, _ := json.Marshal(e)
	http.Error(w, string(msg), e.Code)
}

func writeFile(taskList []task.Task) error {
	bytes, err := json.Marshal(taskList)
	if err != nil {
		return err
	}

	err = os.WriteFile(TaskFileName, bytes, 0777)
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
		if err != nil {
			e := structs.ErrResponse{
				Code:    http.StatusBadRequest,
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
			e := structs.ErrResponse{
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
	e := structs.ErrResponse{
		Code:    http.StatusMethodNotAllowed,
		Message: fmt.Sprintf("expect method POST at /create, got %v", r.Method),
	}
	makeErrorResponse(w, e)
	return

}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := makeResponse(w, TaskList)
		if err != nil {
			e := structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
		}
		return
	}
	e := structs.ErrResponse{
		Code:    http.StatusMethodNotAllowed,
		Message: fmt.Sprintf("expect method GET at /list, got %v", r.Method),
	}
	makeErrorResponse(w, e)
	return

}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var ReadTask task.Task
		_ = json.NewDecoder(r.Body).Decode(&ReadTask)
		for _, v := range TaskList {
			if v.Uuid == ReadTask.Uuid {
				err := makeResponse(w, v)
				if err != nil {
					e := structs.ErrResponse{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					}
					makeErrorResponse(w, e)
					log.Print(err)
					return
				}
				return
			}
		}
		e := structs.ErrResponse{
			Code:    http.StatusNotFound,
			Message: "Task not found",
		}
		makeErrorResponse(w, e)
		return
	}
	e := structs.ErrResponse{
		Code:    http.StatusMethodNotAllowed,
		Message: fmt.Sprintf("expect method GET at /get, got %v", r.Method),
	}
	makeErrorResponse(w, e)
	return

}
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		var DeleteTask task.Task
		_ = json.NewDecoder(r.Body).Decode(&DeleteTask)
		for i, v := range TaskList {
			if v.Uuid == DeleteTask.Uuid {
				TaskList = append(TaskList[:i], TaskList[i+1:]...)
				err := writeFile(TaskList)
				if err != nil {
					e := structs.ErrResponse{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					}
					makeErrorResponse(w, e)
					log.Print(err)
					return
				}
				err = makeResponse(w, TaskList)
				if err != nil {
					e := structs.ErrResponse{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					}
					makeErrorResponse(w, e)
					log.Print(err)
					return
				}
				return
			}
		}
		e := structs.ErrResponse{
			Code:    http.StatusNotFound,
			Message: "element not found",
		}
		makeErrorResponse(w, e)
		return
	}
	e := structs.ErrResponse{
		Code:    http.StatusMethodNotAllowed,
		Message: fmt.Sprintf("expect method DELETE at /delete, got %v", r.Method),
	}
	makeErrorResponse(w, e)
	return
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		var PatchTask task.Task
		_ = json.NewDecoder(r.Body).Decode(&PatchTask)
		for i, _ := range TaskList {
			if TaskList[i].Uuid == PatchTask.Uuid {
				TaskList[i].Name = PatchTask.Name
				TaskList[i].Body = PatchTask.Body
				TaskList[i].Deadline = PatchTask.Deadline
				TaskList[i].Status = PatchTask.Status
				err := writeFile(TaskList)
				if err != nil {
					e := structs.ErrResponse{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					}
					makeErrorResponse(w, e)
					log.Print(err)
					return
				}
				err = makeResponse(w, TaskList)
				if err != nil {
					e := structs.ErrResponse{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					}
					makeErrorResponse(w, e)
					log.Print(err)
					return
				}
			}
			e := structs.ErrResponse{
				Code:    http.StatusNotFound,
				Message: "element not found",
			}
			makeErrorResponse(w, e)
			return
		}
		e := structs.ErrResponse{
			Code:    http.StatusMethodNotAllowed,
			Message: fmt.Sprintf("expect method PATCH at /patch, got %v", r.Method),
		}
		makeErrorResponse(w, e)
		return
	}
}
