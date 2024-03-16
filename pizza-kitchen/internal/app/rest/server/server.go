package server

import (
	"context"
	"fmt"
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

func (s *Server) Run() error {
	return s.httpServ.ListenAndServe()
}
func (s *Server) Stop(ctx context.Context) error {
	return s.httpServ.Shutdown(ctx)
}
