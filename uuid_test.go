package uuid

import (
	"testing"
	"math"
)

func TestCollisions(t *testing.T) {
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
		//println(i, " ┃ ", id, " ┃ ", col, " ┃ ", float64(col) / float64(gen))
	}

	println("Collected ", gen - col, " unique IDs out of ", gen, " generation call")
	println("Found ", col, " collisions for a ratio of ", float64(col) / float64(gen), " collision chance")

	if col != 0 {
		t.Errorf("%d collisions found in %d generations", col, gen)
	}
}