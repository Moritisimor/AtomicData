package atomicstack

import (
	"slices"
	"sync"
)

// A Thread-Safe Stack Implementation which follows the LIFO Principle.
type AtomicStack[T any] struct {
	internal_slice []T
	mutex sync.RWMutex
}

// Creates a new Thread-Safe Stack.
func New[T any]() *AtomicStack[T] {
	return &AtomicStack[T]{}
}

// Method for getting the length of the Stack.
func (s *AtomicStack[T]) Len() int {
	return len(s.internal_slice)
}

// Method for pushing an element into the Stack.
func (s *AtomicStack[T]) Push(item T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.internal_slice = append(s.internal_slice, item)
}

// Method for getting the last element which was pushed into the stack.
// Returns T and bool, where T is the element, and bool indicates if all went right.
// Bool could be false if the Stack is empty, for example.
func (s *AtomicStack[T]) Pop() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.Len() == 0 {
		var t T
		return t, false
	}

	escapee := s.internal_slice[s.Len() - 1]
	s.internal_slice = slices.Delete(s.internal_slice, s.Len() - 1, s.Len())
	return escapee, true
}

// Method for getting the last element which was pushed without removing it.
// Returns T and bool, where T is the element and bool indicates if all went right.
// Like Pop, Bool could be false if the Stack is empty (doesn't exist).
func (s *AtomicStack[T]) Peek() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.Len() == 0 {
		var t T
		return t, false
	}

	return s.internal_slice[s.Len() -1], true
}
