package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, req *http.Request) uint16 {
	fmt.Fprintf(w, "It is an index Page\n")
	return http.StatusOK
}

func getEchoDataFromMap(w http.ResponseWriter, req *http.Request) uint16 {
	var reqVars = mux.Vars(req)

	fmt.Fprintf(w, "data-id: %s", reqVars["id"])
	return http.StatusOK
}

func getCustomNotFoundError(w http.ResponseWriter, req *http.Request) uint16 {
	w.WriteHeader(http.StatusNotFound) // before ServeFile
	// quote: Changing the header map after a call to [ResponseWriter.WriteHeader] (or
	// [ResponseWriter.Write]) has no effect unless the HTTP status code was of the
	// 1xx class or the modified headers are trailers.
	http.ServeFile(w, req, "static/404error.html")
	return http.StatusNotFound
}

func getAgeByName(w http.ResponseWriter, req *http.Request) uint16 {
	nameToAgeMap := map[string]int8{"tim": 19, "kate": 18}

	name := mux.Vars(req)["name"]

	age, ok := nameToAgeMap[name]

	if ok {
		fmt.Fprintf(w, "The age of %s is %d", name, age)
	} else {
		fmt.Fprintf(w, "I know nothing about person with name %s", name)
	}
	return http.StatusOK
}

func httpLoggingMiddleware(fn func(w http.ResponseWriter, req *http.Request) uint16) func(w http.ResponseWriter, req *http.Request) {
	// decorator
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s [%s]", r.Method, r.URL.Path)
		var responseCode uint16 = fn(w, r)
		log.Printf("Completed in %v With code %d", time.Since(start), responseCode)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", httpLoggingMiddleware(index))
	r.HandleFunc("/data/{id}", httpLoggingMiddleware(getEchoDataFromMap))
	r.HandleFunc("/ageOf/{name}", httpLoggingMiddleware(getAgeByName))
	r.PathPrefix("/").HandlerFunc(httpLoggingMiddleware(getCustomNotFoundError))
	http.Handle("/", r)
	socket := "localhost:8080"
	log.Println("Serving on socket", socket)
	http.ListenAndServe(socket, nil)
}
