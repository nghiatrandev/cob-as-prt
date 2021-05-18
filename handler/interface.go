package handler

import "net/http"

type Handler interface {
	Routes() *http.ServeMux
	GetHello(w http.ResponseWriter, req *http.Request)
}
