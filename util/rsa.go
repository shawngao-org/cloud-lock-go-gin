package util

import (
	"cloud-lock-go-gin/logger"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

var pack = "rsa"

func Decrypted(base64Key string, base64String string) string {
	originStr, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		logger.LogErr(pack, "Unable to decode public key.")
		logger.LogErr(pack, "%s", err)
		return ""
	}
	cipherText := originStr
	privateKeyBytes, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		logger.LogErr(pack, "Unable to decode private key.")
		logger.LogErr(pack, "%s", err)
		return ""
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBytes)
	if err != nil {
		logger.LogErr(pack, "Unable to parse private key.")
		logger.LogErr(pack, "%s", err)
		return ""
	}
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		logger.LogErr(pack, "Wrong private key type.")
		return ""
	}
	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, cipherText)
	if err != nil {
		logger.LogErr(pack, "%s", err)
		return ""
	}
	return string(decryptedText)
}

func Encrypted(base64Key string, str string) string {
	plainText := []byte(str)
	publicKeyBytes, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		logger.LogErr(pack, "Unable to decode public key.")
		logger.LogErr(pack, "%s", err)
		return ""
	}
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		logger.LogErr(pack, "Unable to parse public key.")
		logger.LogErr(pack, "%s", err)
		return ""
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		logger.LogErr(pack, "Wrong public key type.")
		return ""
	}
	encryptedText, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, plainText)
	if err != nil {
		logger.LogErr(pack, "%s", err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(encryptedText)
}
