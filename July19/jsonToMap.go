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
func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	data := `{"id": "BM-1347", "title": "The underage storm", "Content": "The creatives' careers can easily get uncreative but yet creative...", "Summary": "Seeking freedom"}`
	Map := map[string]string{}
	json.Unmarshal([]byte(data), &Map)
	fmt.Fprint(w, Map)

}
