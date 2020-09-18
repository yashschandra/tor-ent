package main

import (
	"net/http"
)

const (
	port = ":17666"
	addEntryURL = "/add"
	getEntryURL = "/get"
)

func start() {
	http.HandleFunc(addEntryURL, addEntry)
	http.HandleFunc(getEntryURL, getEntry)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func main() {
	start()
}