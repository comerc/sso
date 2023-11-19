package jwt

import (
	"time"

	"grpc-service-ref/internal/domain/models"

	"github.com/golang-jwt/jwt/v5"
)

// NewToken creates new JWT token for given user and app.
func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"uid":    user.ID,
			"email":  user.Email,
			"exp":    time.Now().Add(duration).Unix(),
			"app_id": app.ID,
		},
	)

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
