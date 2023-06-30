package helper

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	JWTSecret     = "mastermind"
	TokenDuration = time.Hour * 5
)

func GenerateCustomerJWTToken(id uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(TokenDuration).Unix()
	claims["role"] = "customer"

	return token.SignedString([]byte(JWTSecret))
}

func GenerateSellerJWTToken(id uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(TokenDuration).Unix()
	claims["role"] = "seller"

	return token.SignedString([]byte(JWTSecret))
}

func ParseJWTString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil

}

func GenerateIdFromToken(tokenString string) (id uuid.UUID, err error) {
	token, err := ParseJWTString(tokenString)
	if err != nil {
		log.Println(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if idString, ok := claims["id"].(string); ok {
			id, _ = uuid.Parse(idString)
		}
	}

	return id, nil
}

func GenerateRoleFromToken(token *jwt.Token) (role string, err error) {

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if role, ok = claims["role"].(string); !ok {
			return
		}
		return role, nil
	}
	return
}
