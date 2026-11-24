package accountService

import (
	"errors"
	"strconv"
	"time"
	"github.com/golang-jwt/jwt/v5"
)


type TokenData struct {
	Role string `json:"role"`
	jwt.RegisteredClaims // стандартные поля для джот токена 
}


func GenerateJWT(userID int64, role string, secret []byte, ttl time.Duration) (string, error) {
	if ttl <= 0 {
		return "", errors.New("срок действия JWT должен быть положительным")
	}

	now := time.Now()
	data := TokenData{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: strconv.FormatInt(userID, 10),
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl))},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	return token.SignedString(secret)
}


func ValidateJWT(tokenString string, secret []byte) (*TokenData, error) {
	data := &TokenData{}

	getSecret := func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}

	_, err := jwt.ParseWithClaims(tokenString, data, getSecret, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil {
		return nil, err
	}
	return data, nil
}