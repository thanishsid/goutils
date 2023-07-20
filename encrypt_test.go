package goutils

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptDecryptAES(t *testing.T) {
	key := "thisis32bitlongpassphraseimusing"

	text := "hello there how are you today sir !!!"

	encoded, err := EncryptAES([]byte(key), []byte(text))
	require.NoError(t, err)

	t.Logf("cipherText Hex: %s\n", hex.EncodeToString(encoded.Data))
	t.Logf("nonce Hex: %s\n", hex.EncodeToString(encoded.Nonce))
	t.Logf("cipherText Base64: %s\n", Base64Encode(encoded.Data))
	t.Logf("nonce Base64: %s\n", Base64Encode(encoded.Nonce))

	decoded, err := DecryptAES([]byte(key), encoded.Data, encoded.Nonce)

	require.NoError(t, err)

	require.Equal(t, text, string(decoded))
}
