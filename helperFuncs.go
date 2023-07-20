package goutils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/typ.v4"
)

// Map function to iterate and over an slice and return a new slice
func Map[T any, R any](list []T, mapFunc func(node T) R) []R {
	if list == nil {
		return nil
	}

	results := make([]R, len(list))

	for idx, node := range list {
		results[idx] = mapFunc(node)
	}

	return results
}

// Coalesce function.
func Coalesce[T any](input T, condition bool, output T) T {
	if condition {
		return output
	}

	return input
}

// Generic convert function
func Convert[I any, O any](input I, convertFunc func(i I) O) O {
	return convertFunc(input)
}

// Convert string to signed integer
func StringToInt[T typ.Signed](str string, fallback T) T {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return fallback
	}

	return T(i)
}

// Convert string to unsigned integer
func StringToUint[T typ.Unsigned](str string, fallback T) T {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return fallback
	}

	return T(i)
}

// Convert signed integer to string.
func IntToString[T typ.Signed](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

// Convert unsigned intger to string.
func UintToString[T typ.Unsigned](i T) string {
	return strconv.FormatUint(uint64(i), 10)
}

// Convert string to float
func StringToFloat[T typ.Float](str string, fallback T) T {
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fallback
	}

	return T(i)
}

// Convert string to boolean.
func StringToBool(str string, fallback bool) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return fallback
	}

	return b
}

// Get md5 hash string
func GetMD5Hash(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

// Base64Encode - Base64 Encode
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func GetIpFromHttpReq(req *http.Request) (string, error) {
	headerIps := []string{
		req.Header.Get("CF-Connecting-IP"),
		req.Header.Get("X-Forwarded-For"),
		req.Header.Get("X-Real-IP"),
	}

	for _, hip := range headerIps {
		if hip != "" {
			return hip, nil
		}
	}

	remoteAdr := req.RemoteAddr

	adrParts := strings.Split(remoteAdr, ":")

	if len(adrParts) == 2 {
		return adrParts[1], nil
	}

	return "", errors.New("ip not found")
}

func GetContentTypeFromMultipartHeader(f *multipart.FileHeader) string {
	var contentType string

	cTypes := f.Header["Content-Type"]
	for i, ctype := range cTypes {
		if i == 0 {
			contentType = ctype
		} else {
			contentType = fmt.Sprintf("%s;%s", contentType, ctype)
		}
	}

	return contentType
}
