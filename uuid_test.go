package uuid

import (
	"fmt"
	"math"
	"testing"
)

func TestCollisions(t *testing.T) {
	ids := map[UUID]bool{}
	id := NextUUID()
	col := false
	max := math.MaxUint16
	for i := 0; i < max; i++ {
		fmt.Printf("\rgen: %d (%.1f%%)", i, float64(i+1)/float64(max)*100)

		id = NextUUID()
		col = ids[id]
		ids[id] = true

		if col {
			fmt.Printf("\r")
			t.Errorf("collision found after %d generations", i)
		}
	}
	fmt.Println()
}
