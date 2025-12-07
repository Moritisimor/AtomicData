package tests

import (
	"AtomicData/pkg/atomicslice"
	"fmt"
	"sync"
	"testing"
)

// Test for determining if the Atomic Slice is really thread-safe
func TestSlice(t *testing.T) {
	collection := atomicslice.From([]float32{})
	wg := sync.WaitGroup{}
	for i := range(50) {
		wg.Add(1)
		go func(i float32) {
			defer wg.Done()
			collection.Append(i * i)
		}(float32(i))
	}

	wg.Wait()
	for i := range(50) {
		n, _ := collection.Get(i)
		fmt.Println(n)
	}
}
