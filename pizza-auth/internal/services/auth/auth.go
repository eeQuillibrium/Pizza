package auth

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/eeQuillibrium/pizza-auth/internal/domain/models"
	"github.com/eeQuillibrium/pizza-auth/internal/lib/jwt"
	"github.com/eeQuillibrium/pizza-auth/internal/logger"
)

const salt = "fdlkgldfSDASdfglkj@#$23rkjgfdkjg"

type Auth struct {
	userProvider UserProvider
	tokenTTL     time.Duration
	log          *logger.Logger
}

// storage interface
type UserProvider interface {
	CreateUser(
		ctx context.Context,
		phone string,
		passHash string,
	) (userId int64, err error)
	GetUser(
		ctx context.Context,
		phone string,
	) (user models.User, err error)
}

func New(
	userProvider UserProvider,
	tokenTTL time.Duration,
	log *logger.Logger,
) *Auth {
	return &Auth{
		userProvider: userProvider,
		tokenTTL:     tokenTTL,
		log:          log,
	}
}

func (a *Auth) Login(
	ctx context.Context,
	phone string,
	pass string,
) (string, error) {
	a.log.SugaredLogger.Info("starting to login user...")

	hash, err := hashPassword(pass)
	if err != nil {
		a.log.Fatalf("hash generating error: %v", err)
		return "", err
	}

	user, err := a.userProvider.GetUser(ctx, phone)
	if err != nil {
		return "", err
	}

	if user.PassHash != hash {
		a.log.SugaredLogger.Info("don't have that user")
	}

	a.log.SugaredLogger.Info("user was identified")
	a.log.SugaredLogger.Info("try to create jwt")

	jwtStr, err := jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		return "", err
	}

	return jwtStr, nil
}

func (a *Auth) Register(
	ctx context.Context,
	phone string,
	pass string,
) (int64, error) {
	a.log.SugaredLogger.Info("trying to create user")

	hash, err := hashPassword(pass)
	if err != nil {
		a.log.Fatalf("hash generating error: %v", err)
		return 0, err
	}

	userId, err := a.userProvider.CreateUser(ctx, phone, hash)
	if err != nil {
		a.log.Fatalf("user creating error: %v", err)
	}

	a.log.SugaredLogger.Info("log was registered")

	return userId, nil
}
func (a *Auth) IsAdmin(
	ctx context.Context,
	UserId int64,
) (bool, error) {
	return false, nil
}


func (a *Auth) UserIdentify(
	ctx context.Context,
	token string,
) (int, error) {
	return jwt.ParseToken(ctx, token)
}

func hashPassword(pass string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(pass))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum([]byte(salt))), nil
}
