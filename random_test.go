package goutils

import "testing"

func TestCreateRandomString(t *testing.T) {
	t.Parallel()
	str := CreateRandomString(6, func(str string) bool { return true })

	if len(str) != 6 {
		t.Error("random string incorrect length")
	}
}
