package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:8000", nil)

}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	p := r.RemoteAddr
	if p == "127.0.0.1:2658" {
		fmt.Fprint(w, "IP Blocked")
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	}

}
