package tests

import (
	"slices"
	"sync"
	"testing"

	"github.com/Moritisimor/AtomicData/pkg/atomicslice"
)

// Test for determining if the Atomic Slice is really thread-safe
func TestSlice(t *testing.T) {
	N := 50000
	atomicslice := atomicslice.New[float32]()
	wg := sync.WaitGroup{}
	for i := range(N) {
		wg.Add(1)
		go func(i float32) {
			defer wg.Done()
			atomicslice.Append(i * i)
		}(float32(i))
	}

	wg.Wait()
	t.Logf("\nExpected: %d\nGot: %d\n", N, atomicslice.Len())
	if atomicslice.Len() != N {
		t.Error("The expected length is not met!")
	}

	clonedSlice := atomicslice.Clone()
	missingElements := 0
	for i := range(N) {
		if !slices.Contains(clonedSlice, float32(i) * float32(i)) {
			missingElements++
		}
	}

	if missingElements != 0 {
		t.Errorf("The slice is missing %d expected Elements!", missingElements)
	}
}
