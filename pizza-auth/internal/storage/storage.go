package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/eeQuillibrium/pizza-auth/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	DB *sqlx.DB
}

func New(dsn string) *Storage {
	log.Print("trying to open postgres database...")

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("db open problem: %v", err)
	}

	log.Print("success database open!")

	return &Storage{db}
}

func (s *Storage) CreateUser(
	ctx context.Context,
	login string,
	passHash string,
) (userId int64, err error) {

	q := fmt.Sprintf("INSERT INTO %s (login, passhash) VALUES($1, $2) returning id", "users")

	rows := s.DB.QueryRow(q, login, passHash)

	if err := rows.Scan(&userId); err != nil {
		log.Fatalf("can't scan from rows")
	}
	
	return userId, nil
}
func (s *Storage) GetUser(
	ctx context.Context,
	login string,
) (user models.User, err error) {
	return models.User{}, nil
}
