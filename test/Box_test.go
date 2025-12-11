package tests

import (
	"slices"
	"sync"
	"testing"

	"github.com/Moritisimor/AtomicData/pkg/atomicbox"
)

func TestBox(t *testing.T) {
	N := 10000
	intBox := atomicbox.New(0)
	stringBox := atomicbox.New("")
	sliceBox := atomicbox.New([]int{})
	wg := sync.WaitGroup{}

	for range N {
		wg.Go(func() {
			intBox.WithLock(func(inner *int) {
				*inner++
			})
		})
	}

	for range N {
		wg.Go(func () {
			stringBox.WithLock(func(inner *string) {
				*inner += "c"
			})
		})
	}

	for i := range N {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sliceBox.WithLock(func(inner *[]int) {
				*inner = append(*inner, i)
			})
		}(i)
	}

	wg.Wait()
	intBox.WithLock(func(inner *int) {
		if *inner != N {
			t.Errorf("Counter is not as expected!\nExpected: %d\nGot: %d\n", N, *inner)
		}
	})

	stringBox.WithLock(func(inner *string) {
		if len(*inner) != N {
			t.Errorf("String is not as expected!\nExpected length: %d\nGot: %d", N, len(*inner))
		}
	})

	missingElements := 0
	sliceBox.WithLock(func(inner *[]int) {
		for i := range N {
			if !slices.Contains(*inner, i) {
				missingElements++
			}
		}
	})

	if missingElements > 0 {
		t.Errorf("Slice is missing %d elements!", missingElements)
	}

	sliceBox.WithLock(func(inner *[]int) {
		if len(*inner) != N {
			t.Errorf("Slice is not as large as expected!\nExpected %d\nGot: %d", N, len(*inner))
		}
	})
}
