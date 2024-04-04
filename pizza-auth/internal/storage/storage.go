package storage

import (
	"context"
	"fmt"

	"github.com/eeQuillibrium/pizza-auth/internal/domain/models"
	"github.com/eeQuillibrium/pizza-auth/internal/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB  *sqlx.DB
	log *logger.Logger
}

func New(
	dsn string,
	log *logger.Logger,
) *Storage {

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("db open problem: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("db ping error: %v", err)
	}

	return &Storage{
		DB: db, 
		log: log,
	}
}

func (s *Storage) CreateUser(
	ctx context.Context,
	phone string,
	passHash string,
) (userId int64, err error) {

	q := fmt.Sprintf("INSERT INTO %s (phone, passhash) VALUES ($1, $2) RETURNING id", "users")

	row := s.DB.QueryRow(q, phone, passHash)

	if err := row.Scan(&userId); err != nil {
		s.log.Fatalf("can't scan from rows")
	}

	return userId, nil
}
func (s *Storage) GetUser(
	ctx context.Context,
	phone string,
) (user models.User, err error) {

	q := fmt.Sprintf("SELECT (passhash) FROM %s WHERE phone = $1", "users")

	err = s.DB.Get(&user, q, phone)

	return user, err
}
