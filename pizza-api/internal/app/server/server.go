package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/eeQuillibrium/pizza-api/internal/logger"
)

type Server struct {
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
		server: server,
	}
}
func (s *Server) Run() error {
	return s.server.ListenAndServe()
}
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
