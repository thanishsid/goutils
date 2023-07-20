package goutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func EncryptAES(key []byte, body string) (string, error) {
	// create cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	ct := gcm.Seal(nonce, nonce, []byte(body), nil)

	return string(ct), nil
}

func DecryptAES(key []byte, ciphertext string) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plainText, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
