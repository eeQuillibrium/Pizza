package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServ *http.Server
}

func New(
	port int,
	handler http.Handler,
) *Server {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        handler,
		ReadTimeout:    10000 * time.Millisecond,
		WriteTimeout:   10000 * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{server}
}

func (s *Server) Run() {
	err := s.httpServ.ListenAndServe()
	if err != nil {
		log.Fatalf("listen rest kitchen: %v", err)
	}

}
func Stop() {

}
