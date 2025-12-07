package tests

import (
	"AtomicData/pkg/atomicslice"
	"slices"
	"sync"
	"testing"
)

// Test for determining if the Atomic Slice is really thread-safe
func TestSlice(t *testing.T) {
	N := 50000
	collection := atomicslice.From([]float32{})
	wg := sync.WaitGroup{}
	for i := range(N) {
		wg.Add(1)
		go func(i float32) {
			defer wg.Done()
			collection.Append(i * i)
		}(float32(i))
	}

	wg.Wait()
	clonedSlice := collection.Clone()
	if collection.Len() != N {
		t.Error("The expected length is not met!")
	}
	for i := range(N) {
		if !slices.Contains(clonedSlice, float32(i) * float32(i)) {
			t.Error("Could not find expected element!")
		}
	}
}
