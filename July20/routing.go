package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type employee struct {
	ID       int     `json:ID`
	Name     string  `json:Name`
	Dept     string  `json:Dept`
	BasicPay float64 `json:BasicPay`
	Hra      float64 `json:Hra`
	GrossPay float64 `json:GrossPay`
}

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []task // list of tasks

type allEmployee []employee

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some content",
	},
}
var employees = allEmployee{
	{
		ID:       1,
		Name:     "Carter",
		Dept:     "IT",
		BasicPay: 57000,
		Hra:      5700,
		GrossPay: 62700,
	},
}

func calGross(w http.ResponseWriter, r *http.Request) {
	var newEmp employee
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &newEmp)

	newEmp.ID = len(employees) + 1
	if newEmp.Dept == "IT" && newEmp.Hra == 0 {
		fmt.Println("hit")
		newEmp.Hra = newEmp.BasicPay * .10
	} else if newEmp.Dept == "Security" && newEmp.Hra == 0 {
		newEmp.Hra = newEmp.BasicPay * .14
	}
	newEmp.GrossPay = newEmp.BasicPay + newEmp.Hra
	fmt.Println(newEmp.GrossPay, newEmp.Hra)
	employees = append(employees, newEmp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEmp)

}

func getEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

//HTTP GET
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Fprintf(w, "Task with ID %v has been removed successfully", taskID)
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updateTask task

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}
	json.Unmarshal(reqBody, &updateTask)

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			updateTask.ID = taskID
			tasks = append(tasks, updateTask)

			fmt.Fprintf(w, "The task with id %v has been updated successfully", taskID)
		}
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

//ConfigureRouter setup the router
func ConfigureRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/ping", healthCheck).Methods("GET")

	router.HandleFunc("/task", calGross).Methods("POST")
	router.HandleFunc("/task", getEmp).Methods("GET")
	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
