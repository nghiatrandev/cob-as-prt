package handler

import (
	"net/http"
)

type handler struct {
	//Log log.Logger
}

func NewHandler() Handler {
	return handler{}
}

func (h handler) Routes() *http.ServeMux {
	srvMux := http.NewServeMux()
	srvMux.HandleFunc("/hello", h.GetHello)

	return srvMux
}

func (h handler) GetHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
}
