package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type GatewayServer struct {
	http.Server
}

func New(
	port int,
	router http.Handler,
) *GatewayServer {
	server := http.Server{
		Addr:           fmt.Sprintf("localhost:%d", port),
		Handler:        router,
		ReadTimeout:    10000 * time.Millisecond,
		WriteTimeout:   10000 * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}
	return &GatewayServer{
		server,
	}
}
func (s *GatewayServer) Run() error {
	return s.ListenAndServe()
}
func (s *GatewayServer) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}
