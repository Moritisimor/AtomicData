package tests

import (
	"AtomicData/pkg/atomicslice"
	"fmt"
	"sync"
)

// This function will square each number from 0 to iters concurrently via goroutines and collect them in a slice.
// The purpose is not to collect the numbers in correct order, but to demonstrate that syncedslice is really thread-safe.
func CacTest(iters int) {
	collection := atomicslice.From([]float32{})
	wg := sync.WaitGroup{}
	for i := range(iters) {
		wg.Add(1)
		go func(i float32) {
			defer wg.Done()
			collection.Append(i * i)
		}(float32(i))
	}

	wg.Wait()
	for i := range(iters) {
		n, _ := collection.Get(i)
		fmt.Println(n)
	}
}
