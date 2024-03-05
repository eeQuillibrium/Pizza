package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServ *http.Server
}

func New(
	port string,
	handler http.Handler,
) *Server {
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    10000 * time.Millisecond,
		WriteTimeout:   10000 * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{server}
}

func Run() {

}
func Stop() {

}
