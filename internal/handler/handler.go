package handler

import (
	"encoding/json"
	"fmt"
	"github.com/7Maliko7/to-do-list/internal/core"
	"github.com/7Maliko7/to-do-list/internal/handler/structs"
	"log"
	"net/http"
)

const (
	headerContentType = "Content-Type"
	contentTypeJson   = "application/json"
)

var Core core.Core

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == http.MethodPost {
		var newTask structs.CreateTaskRequest
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
		res := structs.CreateTaskResponse{}
		res.Uuid, err = Core.CreateTask(newTask.Name, newTask.Body, newTask.Deadline)
		if err != nil {
			e := structs.ErrResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
			makeErrorResponse(w, e)
			log.Print(err)
			return
		}

		err = makeResponse(w, res)
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
		list := make([]structs.Task, 0, len(taskList))
		for _, v := range taskList {
			list = append(list, structs.Task{
				Uuid:     v.Uuid,
				Name:     v.Name,
				Body:     v.Body,
				Status:   v.Status,
				Deadline: v.Deadline,
			})
		}

		err = makeResponse(w, structs.ListTaskResponse{
			List: list,
		})
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
		var ReadTask structs.GetTaskRequest
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
		if OneTask.Uuid == "" {
			e := structs.ErrResponse{
				Code:    http.StatusNotFound,
				Message: "Task not found",
			}
			makeErrorResponse(w, e)
			return
		}

		err = makeResponse(w, structs.GetTaskResponse{
			Uuid:     OneTask.Uuid,
			Name:     OneTask.Name,
			Body:     OneTask.Body,
			Status:   OneTask.Status,
			Deadline: OneTask.Deadline,
		})
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
		Message: fmt.Sprintf("expect method GET at /get, got %v", r.Method),
	}
	makeErrorResponse(w, e)
	return
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		var DeleteTask structs.DeleteTaskRequest
		_ = json.NewDecoder(r.Body).Decode(&DeleteTask)
		if !isExist(DeleteTask.Uuid) {
			e := structs.ErrResponse{
				Code:    http.StatusNotFound,
				Message: "Task not found",
			}
			makeErrorResponse(w, e)
			return
		}
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
		err = makeResponse(w, structs.DeleteTaskResponse{Code: 200})
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
		var PatchTask structs.UpdateTaskRequest
		_ = json.NewDecoder(r.Body).Decode(&PatchTask)
		if !isExist(PatchTask.Uuid) {
			e := structs.ErrResponse{
				Code:    http.StatusNotFound,
				Message: "Task not found",
			}
			makeErrorResponse(w, e)
			return
		}

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
		err = makeResponse(w, structs.UpdateTaskResponse{Code: 200})
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
		Message: fmt.Sprintf("expect method PATCH at /patch, got %v", r.Method),
	}
	makeErrorResponse(w, e)
	return
}

func makeResponse(w http.ResponseWriter, data any) error {
	w.Header().Set(headerContentType, contentTypeJson)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func makeErrorResponse(w http.ResponseWriter, e structs.ErrResponse) {
	w.Header().Set(headerContentType, contentTypeJson)
	w.WriteHeader(e.Code)
	_ = json.NewEncoder(w).Encode(e)
}

func isExist(uuid string) bool {
	OneTask, err := Core.GetTask(uuid)
	if err != nil {
		log.Print(err)
		return false
	}
	if OneTask.Uuid == "" {
		return false
	}
	return true
}
