// Package atomicboxcontains types, methods and functions for instantiating and interacting with AtomicBoxes.
// The AtomicBox type itself is a thread-safe pointer which uses mutexes.
package atomicbox

import "sync"

// AtomicBox is arguably the most primitive type in this library.
// At its core, it's nothing more than a struct which wraps a value T and a Mutex.
// However, AtomicBox can be used to build almost any other thread-safe structure if you want to.
type AtomicBox[T any] struct {
	val   T
	mutex sync.RWMutex
}

// New method initiates a new AtomicBox with its value being its only parameter.
// It returns a reference to the AtomicBox object which is stored on the heap.
func New[T any](t T) *AtomicBox[T] {
	return &AtomicBox[T]{
		val: t,
	}
}

// WithLock method will lock the box and execute the function fn.
// fn's signature demands that it gets inner, which is of type *T, aka a pointer to the type which the box holds.
// inner will be the alias of the internal value of box, and will represent it in the body of fn.
// While fn is being executed, box is locked, meaning no other goroutine can access it.
func (box *AtomicBox[T]) WithLock(fn func(inner *T)) {
	box.mutex.Lock()
	defer box.mutex.Unlock()
	fn(&box.val)
}

// WithReadLock Method works similarly to the regular WithLock except that other threads may read, but not write.
// It is important to note that mutations can still happen, however, they shouldn't, as this may cause data races.
// Only use this method if fn does not mutate inner.
func (box *AtomicBox[T]) WithReadLock(fn func(inner *T)) {
	box.mutex.RLock()
	defer box.mutex.RUnlock()
	fn(&box.val)
}
