package rand

import (
	"crypto/rand"
	"math/big"
)

var (
	runes = []rune("1234567890abcdefghijklmnopqrstuvwxyz")
)

func RandomStringWithPrefix(prefix string, n int) string {
	return prefix + RandomStringWithLen(n)
}

func RandomStringWithLen(n int) string {
	return randomStringFromSliceWithLen(n)
}

func randomStringFromSliceWithLen(n int) string {
	limit := new(big.Int).SetInt64(int64(len(runes)))
	buf := make([]rune, n)

	for i := range buf {
		j, err := rand.Int(rand.Reader, limit)
		if err != nil {
			return ""
		}

		buf[i] = runes[j.Int64()]
	}

	return string(buf)
}
