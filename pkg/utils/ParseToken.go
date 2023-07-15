package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-delivery/pkg/models"
)

func ParseToken(tokenString string) (claims *models.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("nfdnf5534784nmnmdfj"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok {
		return nil, err
	}
	return claims, nil
}
