package handler

import (
	"encoding/json"
	"fmt"
	"github.com/7Maliko7/to-do-list/internal/core"
	"github.com/7Maliko7/to-do-list/structs"
	"github.com/7Maliko7/to-do-list/task"
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

var Core core.Core

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
		newTask.Uuid, err = Core.CreateTask(newTask.Name, newTask.Body, newTask.Deadline)
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
		taskList, err := Core.GetTaskList()
		if err != nil {
			e := structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
		}

		err = makeResponse(w, taskList)
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
		OneTask, err := Core.GetTask(ReadTask.Uuid)
		if err != nil {
			e := structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}

			makeErrorResponse(w, e)
			log.Print(err)
			return
		}
		err = makeResponse(w, OneTask)
		if err != nil {
			e := structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
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
		taskList, err := Core.GetTaskList()
		if err != nil {
			e := structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
		}
		if taskList == nil {
			e := structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
		}
		err = Core.DeleteTask(DeleteTask.Uuid)
		if err != nil {
			e := structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
		}
		err = makeResponse(w, structs.OkResponse{Code: 200})
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
		Message: fmt.Sprintf("expect method DELETE at /delete, got %v", r.Method),
	}
	makeErrorResponse(w, e)
	return
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		var PatchTask task.Task
		_ = json.NewDecoder(r.Body).Decode(&PatchTask)
		taskList, err := Core.GetTaskList()
		if err != nil {
			e := structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
		}
		if taskList == nil {
			e := structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
		}
		err = Core.UpdateTask(PatchTask.Uuid, PatchTask.Name, PatchTask.Body, PatchTask.Status, PatchTask.Deadline)
		if err != nil {
			e := structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
		}
		err = makeResponse(w, structs.OkResponse{Code: 200})
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
		Code:    http.StatusMethodNotAllowed,
		Message: fmt.Sprintf("expect method PATCH at /patch, got %v", r.Method),
	}
	makeErrorResponse(w, e)
	return
}
