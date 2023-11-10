package rand

import (
	"log"
	"testing"
)

func TestRandomStringWithLen(t *testing.T) {
	str := RandomStringWithLen(11)
	if str == "" {
		t.Fail()
	}

	log.Println("rand str:", str)
}

func TestRandomStringWithPrefix(t *testing.T) {
	str := RandomStringWithPrefix("veth-", 11)
	if str == "" {
		t.Fail()
	}

	log.Println("rand str:", str)
}
