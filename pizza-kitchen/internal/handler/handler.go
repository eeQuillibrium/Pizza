package handler

import (
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}
