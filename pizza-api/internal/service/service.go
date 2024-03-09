package service

import (
	"context"

	"github.com/eeQuillibrium/pizza-api/internal/repository"
	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
)

// OrderProvider - OP
type OrderProvider interface {
	SendMessage(
		ctx context.Context,
		in *nikita_kitchen1.SendOrderReq,
	) error
}

type Service struct {
	OrderProvider
}

func New(
	repo *repository.Repository,
) *Service {
	return &Service{OrderProvider: NewOPService(repo.OrderProvider)}
}
