package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}

var tasks []Tasks

func allTasks() {
	task := Tasks{
		ID:         "1",
		TaskName:   "Your first task",
		TaskDetail: "Makes your first task by using the 'Add Task' button",
		Date:       "2023-05-03",
	}
	tasks = append(tasks, task)

}
func homePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("i'm homepage")
}
func gettask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if taskId["id"] == tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Tasks
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(1000))
	currentTime := time.Now().Format("01-02-2006")
	task.Date = currentTime
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			flag = true
			json.NewEncoder(w).Encode(map[string]string{"status": "Deleted"})
			return
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}
func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Tasks
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			currentTime := time.Now().Format("01-02-2006")
			task.Date = currentTime
			tasks = append(tasks, task)
			flag = true
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettask/", gettask).Queries("id", "{id}").Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/", deleteTask).Queries("id", "{id}").Methods("DELETE")
	router.HandleFunc("/update/", updateTask).Queries("id", "{id}").Methods("PUT")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	allTasks()
	handleRoutes()
}
