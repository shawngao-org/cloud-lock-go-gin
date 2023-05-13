package jwt

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/logger"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var pack = "jwt"

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = username
	claims["exp"] = time.Now().Add(time.Duration(config.Conf.Security.Jwt.Timeout) * time.Second).Unix()
	secret := []byte(config.Conf.Security.Jwt.Secret)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		logger.LogErr(pack, "Token generate -----> FAILED")
		logger.LogErr(pack, "%s", err)
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	if tokenString == "" {
		return nil, errors.New("invalid token")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return []byte(config.Conf.Security.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
