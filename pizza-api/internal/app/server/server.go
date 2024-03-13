package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eeQuillibrium/pizza-api/internal/logger"
)

type Server struct {
	log    *logger.Logger
	server *http.Server
}

func New(
	log *logger.Logger,
	restport int,
	router http.Handler,
) *Server {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", restport),
		Handler:        router,
		ReadTimeout:    10000 * time.Millisecond,
		WriteTimeout:   10000 * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		log:    log,
		server: server,
	}
}
func (s *Server) Run() {
	s.log.SugaredLogger.Infof("run rest server on %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("server running problem: %v", err)
	}
}
