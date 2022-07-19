package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/us", TimeHandler1).Methods("GET")
	r.HandleFunc("/jp", TimeHandler2).Methods("GET")
	r.HandleFunc("/uk", TimeHandler3).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:8000", nil)

}
func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello world")
}
func TimeHandler1(w http.ResponseWriter, _ *http.Request) {
	now := time.Now()
	p := make(map[string]string)
	p["now"] = now.Format(time.ANSIC)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
func TimeHandler2(w http.ResponseWriter, _ *http.Request) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)
	p := make(map[string]string)
	p["now"] = now.Format(time.ANSIC)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
func TimeHandler3(w http.ResponseWriter, _ *http.Request) {
	loc, _ := time.LoadLocation("Europe/London")
	now := time.Now().In(loc)
	p := make(map[string]string)
	p["now"] = now.Format(time.ANSIC)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
