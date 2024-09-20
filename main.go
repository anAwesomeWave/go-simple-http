package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "It is an index Page\n")
}

func getEchoDataFromMap(w http.ResponseWriter, req *http.Request) {
	var reqVars = mux.Vars(req)

	fmt.Fprintf(w, "data-id: %s", reqVars["id"])
}

func getCustomNotFoundError(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/404error.html")
}

func getAgeByName(w http.ResponseWriter, req *http.Request) {
	nameToAgeMap := map[string]int8{"tim": 19, "kate": 18}

	name := mux.Vars(req)["name"]

	age, ok := nameToAgeMap[name]

	if ok {
		fmt.Fprintf(w, "The age of %s is %d", name, age)
	} else {
		fmt.Fprintf(w, "I know nothing about person with name %s", name)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/data/{id}", getEchoDataFromMap)
	r.HandleFunc("/ageOf/{name}", getAgeByName)
	r.PathPrefix("/").HandlerFunc(getCustomNotFoundError)
	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", nil)
}
