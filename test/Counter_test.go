package tests

import (
	"AtomicData/pkg/atomiccounter"
	"sync"
	"testing"
)

// Test for checking if the atomic counter works as expected.
func TestCounter(t *testing.T) {
	N := 500000
	I := 200000
	wg := sync.WaitGroup{}
	ac := atomiccounter.New()

	for range(N) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ac.Increment()
		}()
	}

	for range(I) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ac.Decrement()
		}()
	}

	wg.Wait()
	if ac.Get() != int64(N - I) {
		t.Errorf("Expected %d, Got %d instead. Possible Race.", N - I, ac.Get())
	}
}
