// The atomicslice package provides all necessary types, methods and functions for using AtomicSlices.
// An AtomicSlice is a thread-safe slice implemenation.
package atomicslice

import (
	"slices"
	"sync"
)

// A thread-safe slice-wrapper.
// Internally it works by wrapping a slice and an RWMutex and offering methods for accessing and manipulating data.
type AtomicSlice[T any] struct {
	internal_slice []T
	mutex          sync.RWMutex
}

// Method for appending data to the slice.
// It takes the parameter data and appends it to the slice.
func (s *AtomicSlice[T]) Append(data T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.internal_slice = append(s.internal_slice, data)
}

// Function for creating a new empty slice.
// It returns a pointer to the AtomicSlice-object in the heap.
func New[T any]() *AtomicSlice[T] {
	return &AtomicSlice[T]{
		internal_slice: []T{},
	}
}

// Function for creating a Synced Slice from an existing slice.
// The supplied parameter will be shallowly cloned, as such, storing raw pointers is not recommended.
// Better yet, if you must store pointers, try AtomicBox.
func From[T any](s []T) *AtomicSlice[T] {
	return &AtomicSlice[T]{
		internal_slice: slices.Clone(s),
	}
}

// Method for cloning the internal slice of a Synced Slice.
// It will return the internal slice, which the struct stores.
// The clone is shallow, which means that if you were to, say, store pointers, a cloned map's pointers would still point to the same object.
func (s *AtomicSlice[T]) Clone() []T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return slices.Clone(s.internal_slice)
}

// Method for getting data from the slice by index.
// Get takes the index you want to access as a parameter and, if possible, returns the value at that index.
// It returns T and bool, where T is the value at that index and bool reports if the access was successful or not.
func (s *AtomicSlice[T]) Get(index int) (T, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if len(s.internal_slice) <= int(index) {
		var t T
		return t, false
	}

	return s.internal_slice[index], true
}

// Method for clearing the internal slice which the struct holds.
func (s *AtomicSlice[T]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.internal_slice = []T{}
}

// Method for deleting an entry at an index.
// Returns bool, which reports if the delete went well or not.
func (s *AtomicSlice[T]) Delete(index int) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.internal_slice) <= index || index < 0 {
		return false
	}

	s.internal_slice = slices.Delete(s.internal_slice, index - 1, index)
	return true
}

// Method for getting the length of the slice.
func (s *AtomicSlice[T]) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.internal_slice)
}
