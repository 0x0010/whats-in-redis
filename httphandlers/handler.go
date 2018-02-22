package httphandler

import (
	"net/http"
	"log"
)

// http请求的处理器声明
type HttpHandler interface {
	GetHandlerPath() string
	GetHandlerFunc() http.HandlerFunc
}

func printRequest(request http.Request) {
	log.Println(request.Method, request.Host, request.URL)
}
