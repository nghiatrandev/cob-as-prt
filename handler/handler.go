package handler

import (
	"github.com/prometheus/common/log"
	"net/http"
)

type handler struct {
	Log log.Logger
}

func NewHandler() Handler {
	return handler{Log: nil}
}

func (h handler) Routes() *http.ServeMux {
	srvMux := http.NewServeMux()
	srvMux.HandleFunc("/hello", h.GetHello)

	return srvMux
}

func (h handler) GetHello(w http.ResponseWriter, req *http.Request) {
	h.Log.Info("get hello success")
	w.Write([]byte("hello"))
}
