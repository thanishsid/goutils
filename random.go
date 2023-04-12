package goutils

import (
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var randCharset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + numericCharset
var numericCharset = "0123456789"

// Create random string with an optional function to check for randomness if not returned true will recursively generate new random string
func CreateRandomString(strLength int, checkFunc func(str string) bool) string {
	b := make([]byte, strLength)
	for i := range b {
		b[i] = randCharset[seededRand.Intn(len(randCharset))]
	}

	randStr := string(b)

	if checkFunc == nil || checkFunc(randStr) {
		return randStr
	}

	return CreateRandomString(strLength, checkFunc)
}

// Create random numeric string with an optional function to check for randomness if not returned true will recursively generate new random numeric string
func CreateRandomNumericString(strLength int, checkFunc func(str string) bool) string {
	b := make([]byte, strLength)
	for i := range b {
		b[i] = numericCharset[seededRand.Intn(len(numericCharset))]
	}

	randStr := string(b)

	if checkFunc == nil || checkFunc(randStr) {
		return randStr
	}

	return CreateRandomNumericString(strLength, checkFunc)
}
