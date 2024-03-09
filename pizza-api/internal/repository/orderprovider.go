package repository

import (
	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/redis/go-redis/v9"
)

type OPRepo struct {
	repo *redis.Client
}

func NewOPRepo(
	rClient *redis.Client,
) *OPRepo {
	return &OPRepo{repo: rClient}
}

func (r *OPRepo) StoreOrder(order *models.Order) {

}
