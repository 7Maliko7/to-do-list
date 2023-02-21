package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"time"
)

// `json:"firstname"`
type Task struct {
	Uuid     string    `json:"uuid"`
	Name     string    `json:"name"`
	Body     string    `json:"body"`
	Status   string    `json:"status"`
	Deadline time.Time `json:"deadline"`
}

var TaskList []Task

func wrightFile(file *os.File, taskList []Task) {
	err := file.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(file).Encode(taskList)

}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		var newTask Task
		_ = json.NewDecoder(r.Body).Decode(&newTask)
		newTask.Uuid = uuid.New().String()
		TaskList = append(TaskList, newTask)
		json.NewEncoder(w).Encode(newTask)
		wrightFile(file, TaskList)
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
				wrightFile(file, TaskList)
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
				wrightFile(file, TaskList)
				return
			}
		}
		http.Error(w, fmt.Sprint("element not found"), http.StatusMethodNotAllowed)
		return
	}
	http.Error(w, fmt.Sprintf("expect method PATCH at /delete, got %v", r.Method), http.StatusMethodNotAllowed)
	return
}

var file *os.File

func main() {
	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/list", ListHandler)
	http.HandleFunc("/", GetHandler)
	http.HandleFunc("/delete", DeleteHandler)
	http.HandleFunc("/update", UpdateHandler)
	var err error
	file, err = os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer func() {
		file.Close()
	}()

	_ = json.NewDecoder(file).Decode(&TaskList)

	log.Fatal(http.ListenAndServe(":8080", nil))
}