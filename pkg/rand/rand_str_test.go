package rand

import (
	"log"
	"testing"
)

func TestRandomStringWithLen(t *testing.T) {
	str := RandomStringWithLen(12)
	if str == "" {
		t.Fail()
	}

	log.Println("rand str:", str)
}

func TestRandomStringWithPrefix(t *testing.T) {
	str := RandomStringWithPrefix("hey-", 12)
	if str == "" {
		t.Fail()
	}

	log.Println("rand str:", str)
}
