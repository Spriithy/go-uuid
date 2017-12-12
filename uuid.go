package uuid

import (
	"crypto/rand"
	"errors"
)

// A UUID is a Universal Unique IDentifier
type UUID = string

var hex = "abcdef0123456789"

// Length is the generic UUID length
const Length = len("xxxxxxxx-xxxx-xxxx-xxxxxxxxxxxxxxxx")

func format(bytes []byte) UUID {
	uuid := ""
	for i := 0; i < 32; i++ {
		switch i {
		case 8, 12, 16:
			uuid += "-"
		}
		uuid += UUID(bytes[i])
	}
	return uuid
}

func isHex(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

// ValidateUUID parses an input string and checks its validity
func ValidateUUID(src string) (UUID, error) {
	for i := 0; i < Length; i++ {
		switch i {
		case 8, 12, 16:
			if src[i] != '-' {
				return "", errors.New("input string doesn't match uuid.UUID pattern")
			}
			continue
		}
		if !isHex(src[i]) {
			return "", errors.New("non-hexadecimal character encountered while parsing uuid.UUID")
		}
	}
	return UUID(src), nil
}

// NextUUID is the actual UUID generator
func NextUUID() UUID {
	uuid := make([]byte, 32)
	rand.Read(uuid)
	for i, b := range uuid {
		uuid[i] = hex[int(b)%len(hex)]
	}
	return format(uuid)
}
