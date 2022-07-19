package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var s = new(store)

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
	} else if len(s.getData()) == 0 {
		t := time.Now().Format(time.ANSIC)
		s.storeData(t)
		fmt.Fprint(w, "IP: "+p+", first accessed: "+t)
	} else {
		fmt.Fprintln(w, "IP: "+p+", last accessed: "+s.getData()[len(s.getData())-2])
		fmt.Fprint(w, "All times accessed: ")
		fmt.Fprint(w, strings.Join(s.getData(), ", "))
		t := time.Now().Format(time.ANSIC)
		s.storeData(t)
	}

}

type store struct {
	data []string
}

func (a *store) storeData(s string) {
	a.data = append(a.data, s)
}
func (a *store) getData() []string {
	return a.data
}
