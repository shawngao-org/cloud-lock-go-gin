package jwt

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/logger"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = username
	claims["exp"] = time.Now().Add(time.Duration(config.Conf.Security.Jwt.Timeout)).Unix()
	secret := []byte(config.Conf.Security.Jwt.Secret)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		logger.LogErr("[Token] Token generate -----> FAILED")
		logger.LogErr("%s", err)
		return "", err
	}
	return tokenString, nil
}
