package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/eeQuillibrium/pizza-auth/internal/domain/models"
	"github.com/golang-jwt/jwt"
)

const (
	signingKey = "dfgdf32423tk[ogdf"
	zeroInt    = 0
	zeroStr    = ""
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"uid"`
}

func NewToken(
	user models.User,
	duration time.Duration,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		int(user.Id),
	})

	tokenStr, err := token.SignedString(signingKey)
	if err != nil {
		return zeroStr, err
	}

	return tokenStr, nil
}

func ParseToken(ctx context.Context, tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return signingKey, nil
	})
	if err != nil {
		return zeroInt, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return zeroInt, errors.New("can't transform tokenClaims")
	}

	return claims.UserId, nil
}
