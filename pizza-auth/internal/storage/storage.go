package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/eeQuillibrium/pizza-auth/internal/domain/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	if err := db.Ping(); err != nil {
		log.Fatalf("db ping error: %v", err)
	}

	log.Print("success database open!")

	return &Storage{db}
}

func (s *Storage) CreateUser(
	ctx context.Context,
	login string,
	passHash string,
) (userId int64, err error) {

	q := fmt.Sprintf("INSERT INTO %s (login, passhash) VALUES ($1, $2) RETURNING id", "users")

	row := s.DB.QueryRow(q, login, passHash)

	if err := row.Scan(&userId); err != nil {
		log.Fatalf("can't scan from rows")
	}

	return userId, nil
}
func (s *Storage) GetUser(
	ctx context.Context,
	login string,
) (user models.User, err error) {

	q := fmt.Sprintf("SELECT (passhash) FROM %s WHERE login = $1", "users")

	err = s.DB.Get(&user, q, login)

	return user, err
}
