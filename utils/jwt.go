package utils

import (
	"fmt"
	"time"
	"tugas4go/app/model"
	"tugas4go/app/model/mongo"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key-1234567890")

// GenerateTokenPostgres → untuk user Postgres
func GenerateTokenPostgres(user model.User) (string, error) {
	userID := fmt.Sprintf("%d", user.ID)

	claims := model.JWTClaims{
		UserID:   userID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// GenerateTokenMongo → untuk user Mongo
func GenerateTokenMongo(user mongo.UserMongo) (string, error) {
	userID := ""
	if !user.ID.IsZero() {
		userID = user.ID.Hex() // convert ObjectID → string
	}

	claims := model.JWTClaims{
		UserID:   userID,
		Username: user.Name,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken memvalidasi token JWT
func ValidateToken(tokenString string) (*model.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}
