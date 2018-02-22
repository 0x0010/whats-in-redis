package httphandler

import (
	"net/http"
	"io"
)

type IndexHandler struct {
}

func (h IndexHandler) GetHandlerPath() string {
	return "/"
}

func (h IndexHandler) GetHandlerFunc() http.HandlerFunc {
	return indexHandler
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	printRequest(*r)
	io.WriteString(w, "Index!")
}
