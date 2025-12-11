package tests

import (
	"sync"
	"testing"

	"github.com/Moritisimor/AtomicData/pkg/atomicbox"
)

func TestBox(t *testing.T) {
	N := 10000
	intBox := atomicbox.New(0)
	stringBox := atomicbox.New("")
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
}
