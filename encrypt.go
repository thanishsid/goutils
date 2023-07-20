package goutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

// Encrypt to AES using GCM as cipher.
func EncryptAES(key []byte, data []byte) (*AESEncryptResult, error) {
	// create cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	ct := gcm.Seal(nil, nonce, data, nil)

	return &AESEncryptResult{
		Data:  ct,
		Nonce: nonce,
	}, nil
}

// Decrypt AES GCM encrypted data.
func DecryptAES(key []byte, cipherData []byte, nonce []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	decoded, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

type AESEncryptResult struct {
	Data  []byte
	Nonce []byte
}
