// Package atomicstack package contains functions, methods and types for using AtomicStacks.
// An AtomicStack is a thread-safe Last-in First-out stack implementation.
package atomicstack

import (
	"slices"
	"sync"
)

// AtomicStack is a Thread-Safe Stack Implementation which follows the LIFO Principle.
type AtomicStack[T any] struct {
	internalSlice []T
	mutex         sync.RWMutex
}

// New function Creates a new Thread-Safe Stack.
func New[T any]() *AtomicStack[T] {
	return &AtomicStack[T]{}
}

// Len Method for getting the length of the Stack.
func (s *AtomicStack[T]) Len() int {
	return len(s.internalSlice)
}

// Push Method for pushing an element into the Stack.
func (s *AtomicStack[T]) Push(item T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.internalSlice = append(s.internalSlice, item)
}

// Pop Method for getting the last element which was pushed into the stack.
// Returns T and bool, where T is the element, and bool indicates if all went right.
// Bool could be false if the Stack is empty, for example.
func (s *AtomicStack[T]) Pop() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.Len() == 0 {
		var t T
		return t, false
	}

	escapee := s.internalSlice[s.Len()-1]
	s.internalSlice = slices.Delete(s.internalSlice, s.Len()-1, s.Len())
	return escapee, true
}

// Peek Method for getting the last element which was pushed without removing it.
// Returns T and bool, where T is the element and bool indicates if all went right.
// Like Pop, Bool could be false if the Stack is empty (doesn't exist).
func (s *AtomicStack[T]) Peek() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.Len() == 0 {
		var t T
		return t, false
	}

	return s.internalSlice[s.Len()-1], true
}
