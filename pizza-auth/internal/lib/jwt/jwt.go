package jwt

import (
	"time"

	"github.com/eeQuillibrium/pizza-auth/internal/domain/models"
	"github.com/golang-jwt/jwt"
)

var jwtSecretKey = []byte("dfgdf32423tk[ogdf")

func NewToken(user models.User, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   user.Id,
		"login": user.Login,
		"exp":   time.Now().Add(duration).Unix(),
	})

	tokenStr, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
