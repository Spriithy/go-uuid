package uuid

import (
	"crypto/rand"
)

var hex = "abcdef0123456789"

type UUID string

func format(bytes []byte) UUID {
	var uuid string = ""
	for i := 0; i < 32; i++ {
		switch i {
		case 8, 12, 16: uuid += "-"
		}
		uuid += string(bytes[i])
	}
	return UUID(uuid)
}

func NextUUID() UUID {
	uuid := make([]byte, 32)
	rand.Read(uuid)
	for i, b := range uuid {
		uuid[i] = hex[int(b) % len(hex)]
	}
	return format(uuid)
}

func Match(id1, id2 UUID) bool {
	return id1 == id2
}

func (id1 UUID) Match(id2 UUID) bool {
	return Match(id1, id2)
}