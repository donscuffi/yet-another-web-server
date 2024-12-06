package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "counter is", strconv.Itoa(counter))
	} else {
		fmt.Fprintln(w, "Method GET allowed only")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		counter++
		fmt.Fprintln(w, "counter increased by 1")
	} else {
		fmt.Fprintln(w, "Method POST allowed only")
	}
}

func main() {
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/post", PostHandler)

	http.ListenAndServe("localhost:8080", nil)
}
