package atomicstack

import (
	"slices"
	"sync"
)

type AtomicStack[T any] struct {
	internal_slice []T
	mutex sync.RWMutex
}

func New[T any]() AtomicStack[T] {
	return AtomicStack[T]{}
}

func (s *AtomicStack[T]) Len() int {
	return len(s.internal_slice)
}

func (s *AtomicStack[T]) Push(item T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.internal_slice = append(s.internal_slice, item)
}

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

func (s *AtomicStack[T]) Peek() (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if s.Len() == 0 {
		var t T
		return t, false
	}

	return s.internal_slice[s.Len() -1], true
}
