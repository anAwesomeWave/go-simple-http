package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "It is an index Page\n")
}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe("localhost:8080", nil)
}
