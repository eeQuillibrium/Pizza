package server

import (
	"log"
	"net/http"
	"time"

	"github.com/eeQuillibrium/pizza-api/internal/handler"
	"github.com/gorilla/mux"
)

type Server struct {
	server  *http.Server
	handler *handler.Handler
}

func New(
	restport string,
	router *mux.Router,
) *Server {
	server := &http.Server{
		Addr:           ":" + restport,
		Handler:        router,
		ReadTimeout:    10000 * time.Millisecond,
		WriteTimeout:   10000 * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		server,
		nil,
	}
}
func (s *Server) Run(grpcport string) {
	s.handler = handler.New(grpcport)

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("server running problem: %v", err)
	}
}
