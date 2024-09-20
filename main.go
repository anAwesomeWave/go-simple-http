package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "It is an index Page\n")
}

func getDataFromMap(w http.ResponseWriter, req *http.Request) {
	var reqVars = mux.Vars(req)

	fmt.Fprintf(w, "data-id: %s", reqVars["id"])
}

func getCustomNotFoundError(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/404error.html")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/data/{id}", getDataFromMap)
	r.PathPrefix("/").HandlerFunc(getCustomNotFoundError)
	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", nil)
}
