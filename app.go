package main

import (
	"net/http"
	"github.com/0x0010/whats-in-redis/httphandlers"
)

// register http handlers
var handlers = []httphandler.HttpHandler{new(httphandler.IconHandler), new(httphandler.IndexHandler)}

func main() {

	for _, handler := range handlers {
		http.HandleFunc(handler.GetHandlerPath(), handler.GetHandlerFunc())
	}
	http.ListenAndServe(":8080", nil)
}
