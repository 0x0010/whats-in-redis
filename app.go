package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "Hello World!")
}
