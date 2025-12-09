package tests

import (
	"sync"
	"testing"

	"github.com/Moritisimor/AtomicData/pkg/atomicstack"
	"github.com/Moritisimor/AtomicData/pkg/atomiccounter"
)

// Test for determining if the Stack works as it should.
func TestStack(t *testing.T) {
	I := 2000000
	D := 1000000

	stack := atomicstack.New[int]()
	counter := atomiccounter.New()
	wg := sync.WaitGroup{}

	for range I {
		wg.Go(func() {
			stack.Push(int(counter.Get()))
			counter.Increment()
		})
	}

	for range D {
		wg.Go(func() {
			stack.Pop()
			counter.Increment()
		})
	}

	wg.Wait()
	t.Logf("\nPushes: %d\nPops: %d\nExpected: %d\nGot: %d\n", I, D, I - D, stack.Len())
	if stack.Len() != I - D {
		t.Error("The stack is not as large as expected! Possible Race!")
	}
}
