package atomicbox

import "sync"

type AtomicBox[T any] struct {
	val T
	mutex sync.RWMutex
}

func New[T any](t T) AtomicBox[T] {
	return AtomicBox[T]{
		val: t,
	}
}

func (box *AtomicBox[T]) WithLock(fn func(inner *T)) {
	box.mutex.Lock()
	defer box.mutex.Unlock()
	fn(&box.val)
}

