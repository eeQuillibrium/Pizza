package server

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func New(
	restport string,
	router http.Handler,
) *Server {
	server := &http.Server{
		Addr:           ":" + restport,
		Handler:        router,
		ReadTimeout:    10000 * time.Millisecond,
		WriteTimeout:   10000 * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{server}
}
func (s *Server) Run() {
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("server running problem: %v", err)
	}
}
