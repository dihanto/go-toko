package helper

import (
	"log"
	"time"

	"github.com/dihanto/go-toko/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GenerateCustomerJWTToken(id uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	tokenDuration := getJWTExpired()
	claims["exp"] = time.Now().Add(tokenDuration).Unix()
	claims["role"] = "customer"
	secretKey := getJWTSecret()
	return token.SignedString([]byte(secretKey))
}

func GenerateSellerJWTToken(id uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	tokenDuration := getJWTExpired()
	claims["exp"] = time.Now().Add(tokenDuration).Unix()
	claims["role"] = "seller"
	secretKey := getJWTSecret()
	return token.SignedString([]byte(secretKey))
}

func ParseJWTString(tokenString string) (*jwt.Token, error) {
	secretKey := getJWTSecret()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
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

func getJWTSecret() string {
	config.InitLoadConfiguration()
	secretKey := viper.GetString("jwt.secret_key")
	return secretKey
}

func getJWTExpired() time.Duration {
	config.InitLoadConfiguration()
	duration := viper.GetDuration("jwt.duration")
	return duration * time.Hour
}
