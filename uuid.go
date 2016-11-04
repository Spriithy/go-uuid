package main

import (
	"crypto/rand"
	"math"
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

func main() {
	ids := map[UUID]bool{}
	gen := 0
	col := 0
	id := NextUUID()
	for i := 0; i < math.MaxUint16; i++ {
		id = NextUUID()
		if ids[id] {
			col++
		}
		ids[id] = true
		gen++
		println(i, " ┃ ", id, " ┃ ", col, " ┃ ", float64(col) / float64(gen))
	}

	println("Collected ", gen - col, " unique IDs out of ", gen, " generation call")
	println("Found ", col, " collisions for a ratio of ", float64(col) / float64(gen), " collision chance")
}
