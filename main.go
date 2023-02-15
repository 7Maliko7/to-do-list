package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// `json:"firstname"`
type Task struct {
	Uuid     string    `json:"uuid"`
	Name     string    `json:"name"`
	Body     string    `json:"body"`
	Deadline time.Time `json:"deadline"`
}

var TaskList []Task

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var NewTask Task
	_ = json.NewDecoder(r.Body).Decode(&NewTask)
	NewTask.Uuid = strconv.Itoa(rand.Intn(1000000))
	TaskList = append(TaskList, NewTask)
	json.NewEncoder(w).Encode(NewTask)

}

//viewHandler (){
//
//}
//editHandler (){
//
//}

func main() {
	http.HandleFunc("/create", CreateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
