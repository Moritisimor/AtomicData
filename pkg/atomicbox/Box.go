// The atomicbox package contains types, methods and functions for instantiating and interacting with AtomicBoxes.
// The AtomicBox type itself is a thread-safe pointer which uses mutexes.
package atomicbox

import "sync"

// AtomicBox is arguably the most primitive type in this library.
// At its core, it's nothing more than a struct which wraps a value T and a Mutex.
// However, AtomicBox can be used to build almost any other thread-safe structure if you want to.
type AtomicBox[T any] struct {
	val T
	mutex sync.Mutex
}

// This method initiates a new AtomicBox with its value being its only parameter.
// It returns a reference to the AtomicBox object which is stored on the heap.
func New[T any](t T) *AtomicBox[T] {
	return &AtomicBox[T]{
		val: t,
	}
}

// This method will lock the box and execute the function fn.
// fn's signature demands that it gets inner, which is of type *T, aka a pointer to the type which the box holds.
// inner will be the alias of the internal value of box, and will represent it in the body of fn.
// While fn is being executed, box is locked, meaning no other goroutine can access it.
func (box *AtomicBox[T]) WithLock(fn func(inner *T)) {
	box.mutex.Lock()
	defer box.mutex.Unlock()
	fn(&box.val)
}

