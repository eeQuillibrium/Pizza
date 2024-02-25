package storage

import (
	"context"

	"github.com/eeQuillibrium/pizza-auth/internal/domain/models"
)

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}


func (s *Storage) CreateUser(
	ctx context.Context,
	login string,
	passHash string,
) (userId int64, err error) {
	return 0, nil
}
func (s *Storage) GetUser(
	ctx context.Context,
	login string,
) (user models.User, err error) {
	return models.User{}, nil
}
