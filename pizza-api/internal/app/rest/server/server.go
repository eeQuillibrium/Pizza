package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/eeQuillibrium/pizza-api/internal/logger"
)

type GatewayServer struct {
	*http.Server
}

func New(
	log *logger.Logger,
	restport int,
	router http.Handler,
) *GatewayServer {
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", restport),
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
	return s.Server.ListenAndServe()
}
func (s *GatewayServer) Stop(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
