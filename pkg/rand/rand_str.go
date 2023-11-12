package rand

import (
	"crypto/rand"
	"fmt"
)

func RandomStringWithPrefix(prefix string, n int) string {
	return prefix + RandomStringWithLen(n)
}

func RandomStringWithLen(n int) string {
	if str, err := randomStringFromHexWithLen(n); err != nil {
		return ""
	} else {
		return str
	}
}

func randomStringFromHexWithLen(n int) (string, error) {
	var l int
	if (n % 2) != 0 {
		l = n/2 + 1
	} else {
		l = n / 2
	}
	entropy := make([]byte, l)
	_, err := rand.Read(entropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate random string: %#v", err)
	}

	return fmt.Sprintf("%x", entropy)[:n], nil
}
