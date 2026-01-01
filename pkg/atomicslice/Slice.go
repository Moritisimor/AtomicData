// Package atomicslice provides all necessary types, methods and functions for using AtomicSlices.
// An AtomicSlice is a thread-safe slice implemenation.
package atomicslice

import (
	"slices"
	"sync"
)

// AtomicSlice is a thread-safe slice-wrapper.
// Internally it works by wrapping a slice and an RWMutex and offering methods for accessing and manipulating data.
type AtomicSlice[T any] struct {
	internalSlice []T
	mutex         sync.RWMutex
}

// Append Method for appending data to the slice.
// It takes the parameter data and appends it to the slice.
func (s *AtomicSlice[T]) Append(data T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.internalSlice = append(s.internalSlice, data)
}

// New function for creating a new empty slice.
// It returns a pointer to the AtomicSlice-object in the heap.
func New[T any]() *AtomicSlice[T] {
	return &AtomicSlice[T]{
		internalSlice: []T{},
	}
}

// From function for creating a Synced Slice from an existing slice.
// The supplied parameter will be shallowly cloned, as such, storing raw pointers is not recommended.
// Better yet, if you must store pointers, try AtomicBox.
func From[T any](s []T) *AtomicSlice[T] {
	return &AtomicSlice[T]{
		internalSlice: slices.Clone(s),
	}
}

// Clone Method for cloning the internal slice of a Synced Slice.
// It will return the internal slice, which the struct stores.
// The clone is shallow, which means that if you were to, say, store pointers, a cloned map's pointers would still
// point to the same object.
func (s *AtomicSlice[T]) Clone() []T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return slices.Clone(s.internalSlice)
}

// Get Method for getting data from the slice by index.
// Get takes the index you want to access as a parameter and, if possible, returns the value at that index.
// It returns T and bool, where T is the value at that index and bool reports if the access was successful or not.
func (s *AtomicSlice[T]) Get(index int) (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if len(s.internalSlice) <= index {
		var t T
		return t, false
	}

	return s.internalSlice[index], true
}

// Clear Method for clearing the internal slice which the struct holds.
func (s *AtomicSlice[T]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.internalSlice = []T{}
}

// Delete Method for deleting an entry at an index.
// Returns bool, which reports if the delete went well or not.
func (s *AtomicSlice[T]) Delete(index int) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.internalSlice) <= index || index < 0 {
		return false
	}

	s.internalSlice = slices.Delete(s.internalSlice, index-1, index)
	return true
}

// Len Method for getting the length of the slice.
func (s *AtomicSlice[T]) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.internalSlice)
}
