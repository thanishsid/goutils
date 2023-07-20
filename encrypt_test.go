package goutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptAES(t *testing.T) {
	key := "thisis32bitlongpassphraseimusing"

	text := "hello there how are you today sir !!!"

	encoded, err := EncryptAES([]byte(key), text)
	require.NoError(t, err)

	t.Log(encoded)

	decoded, err := DecryptAES([]byte(key), encoded)
	require.NoError(t, err)

	require.Equal(t, text, decoded)
}
