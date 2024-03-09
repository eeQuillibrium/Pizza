package service

import (
	"context"

	"github.com/eeQuillibrium/pizza-api/internal/repository"
	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
)

type OPService struct {
	repo repository.OrderProvider
}

func NewOPService(
	repo repository.OrderProvider,
) *OPService {
	return &OPService{repo: repo}
}
func (s *OPService) SendMessage(
	ctx context.Context,
	in *nikita_kitchen1.SendOrderReq,
) error {
	return nil
}
