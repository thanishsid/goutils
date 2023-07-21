package goutils

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptDecryptAESWithoutNoncePrefix(t *testing.T) {
	key := "thisis32bitlongpassphraseimusing"

	text := "hello there how are you today sir !!!"

	encoded, err := EncryptAES([]byte(key), []byte(text), false)
	require.NoError(t, err)

	log.Printf("encoded: %s", hex.EncodeToString(encoded.Data))
	log.Printf("nonce: %s", hex.EncodeToString(encoded.Nonce))

	decoded, err := DecryptAES([]byte(key), encoded.Data, encoded.Nonce)

	require.NoError(t, err)

	require.Equal(t, text, string(decoded))
}

func TestEncryptDecryptAESWithNoncePrefix(t *testing.T) {
	key := "thisis32bitlongpassphraseimusing"

	text := "hello there how are you today sir !!!"

	encoded, err := EncryptAES([]byte(key), []byte(text), true)
	require.NoError(t, err)

	log.Printf("encoded: %s", hex.EncodeToString(encoded.Data))
	log.Printf("nonce: %s", hex.EncodeToString(encoded.Nonce))

	decoded, err := DecryptAES([]byte(key), encoded.Data, nil)

	require.NoError(t, err)

	require.Equal(t, text, string(decoded))
}
