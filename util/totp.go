package util

import (
	"cloud-lock-go-gin/logger"
	totp2 "github.com/pquerna/otp/totp"
	"time"
)

func Totp() (string, error) {
	key, err := totp2.Generate(totp2.GenerateOpts{
		Issuer:      "cloud_lock",
		AccountName: "2519306859@qq.com",
		Period:      0,
		SecretSize:  0,
		Secret:      nil,
		Digits:      0,
		Algorithm:   0,
		Rand:        nil,
	})
	if err != nil {
		logger.LogErr("Cannot generate TOTP key: ", err)
		return "", err
	}
	logger.LogInfo("Key: ", key.URL())
	return key.Secret(), nil
}

func GenerateCode(secret string) (string, error) {
	now := time.Now()
	totpCode, err := totp2.GenerateCode(secret, now)
	if err != nil {
		logger.LogErr("Cannot generate TOTP Key: ", err)
		return "", err
	}
	logger.LogInfo("[%s] TOTP Key: %s", now.Format(time.RFC3339), totpCode)
	return totpCode, nil
}
