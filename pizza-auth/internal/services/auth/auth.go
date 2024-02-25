package auth

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"time"

	"github.com/eeQuillibrium/pizza-auth/internal/domain/models"
	"github.com/eeQuillibrium/pizza-auth/internal/lib/jwt"
)

type Auth struct {
	userProvider UserProvider
	tokenTTL     time.Duration
}

// storage interface
type UserProvider interface {
	CreateUser(
		ctx context.Context,
		login string,
		passHash string,
	) (userId int64, err error)
	GetUser(
		ctx context.Context,
		login string,
	) (user models.User, err error)
}

func New(
	userProvider UserProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		userProvider: userProvider,
		tokenTTL:     tokenTTL,
	}
}

func (a *Auth) Login(
	ctx context.Context,
	login string,
	pass string,
) (string, error) {
	log.Print("starting to login user...")

	hash, err := hashPassword(pass)
	if err != nil {
		log.Fatalf("hash generating error: %v", err)
		return "", err
	}

	user, err := a.userProvider.GetUser(ctx, login)
	if err != nil {
		return "", err
	}

	if user.PassHash != hash {
		log.Print("don't have that user")
	}

	log.Print("user was identificated")
	log.Print("try to create jwt...")

	jwtStr, err := jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		return "", err
	}
	
	return jwtStr, nil
}

func (a *Auth) Register(
	ctx context.Context,
	login string,
	pass string,
) (int64, error) {
	log.Print("try to create user...")

	hash, err := hashPassword(pass)
	if err != nil {
		log.Fatalf("hash generating error: %v", err)
		return 0, err
	}

	userId, err := a.userProvider.CreateUser(ctx, login, hash)
	if err != nil {
		log.Fatalf("user creating error: %v", err)
	}

	log.Print("log war registered")

	return userId, nil
}
func (a *Auth) IsAdmin(
	ctx context.Context,
	UserId int64,
) (bool, error) {
	return false, nil
}

const salt = "fdlkgldfSDASdfglkj@#$23rkjgfdkjg"

func hashPassword(pass string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(pass))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum([]byte(salt))), nil
}
