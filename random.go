package goutils

import (
	"math/rand"
	"strings"
	"time"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const AlphabeticCharsetLower = "abcdefghijklmnopqrstuvwxyz"
const AlphabeticCharsetUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NumericCharset = "0123456789"

var defaultCharset = AlphabeticCharsetLower + AlphabeticCharsetUpper + NumericCharset

// Create random string with an optional function to check for randomness if not returned true will recursively generate new random string.
// Optionally provide charsets to use when generating the random string. If no charsets are provided then a random alphanumeric string will be generated.
func CreateRandomString(strLength int, checkFunc func(str string) bool, charsets ...string) string {
	charset := strings.Join(charsets, "")

	if charset == "" {
		charset = defaultCharset
	}

	b := make([]byte, strLength)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	randStr := string(b)

	if checkFunc == nil || checkFunc(randStr) {
		return randStr
	}

	return CreateRandomString(strLength, checkFunc)
}

// Create random numeric string with an optional function to check for randomness if not returned true will recursively generate new random numeric string.
func CreateRandomNumericString(strLength int, checkFunc func(str string) bool) string {
	b := make([]byte, strLength)
	for i := range b {
		b[i] = NumericCharset[seededRand.Intn(len(NumericCharset))]
	}

	randStr := string(b)

	if checkFunc == nil || checkFunc(randStr) {
		return randStr
	}

	return CreateRandomNumericString(strLength, checkFunc)
}
