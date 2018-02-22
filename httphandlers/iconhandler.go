package httphandler

import (
	"net/http"
	"io"
)

// Icon请求处理器
type IconHandler struct {
}

func (h IconHandler) GetHandlerPath() string {
	return "/icon"
}

func (h IconHandler) GetHandlerFunc() http.HandlerFunc {
	return iconHandler
}

func iconHandler(w http.ResponseWriter, r *http.Request) {
	printRequest(*r)
	io.WriteString(w, "Icon!")
}
