package tests

import (
	"sync"
	"testing"

	"github.com/Moritisimor/AtomicData/pkg/atomicmap"
	"github.com/Moritisimor/AtomicData/pkg/atomiccounter"
)

// Test for checking if the Atomic Map works as intended.
func TestMap(t *testing.T) {
	N := 50000
	syncedmap := atomicmap.New[int64, float64]()
	wg := sync.WaitGroup{}

	conflicts := atomiccounter.New()
	for i := range N {
		wg.Add(1)
		go func (n int64) {
			defer wg.Done()
			if !syncedmap.SetIfNotExists(n, float64(n)) {
				conflicts.Increment()
			}
		}(int64(i))
	}

	wg.Wait()
	if conflicts.Get() != 0 {
		t.Errorf("\nDetected %d Conflicts! Possible Race!", conflicts)
	}

	t.Logf("\nExpected length: %d\nGot: %d", N, syncedmap.Len())
	if syncedmap.Len() != N {
		t.Errorf("\nThe expected length was not met!")
	}
}
