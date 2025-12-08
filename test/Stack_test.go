package tests

import (
	"sync"
	"testing"

	"github.com/Moritisimor/AtomicData/pkg/atomicstack"
	"github.com/Moritisimor/AtomicData/pkg/atomiccounter"
)

func TestStack(t *testing.T) {
	I := 500000
	D := 500000

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
	if stack.Len() != I - D {
		t.Error("The Stack is not as large as expected. Either the atomic counter failed or the stack did.")
	} else {
		t.Logf("Pushes: %d\nPops: %d\nExpected: %d\nGot: %d\n", I, D, I - D, stack.Len())
	}
}
