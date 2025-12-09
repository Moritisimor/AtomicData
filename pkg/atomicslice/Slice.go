package atomicslice

import (
	"slices"
	"sync"
)

// A thread-safe slice-wrapper.
// Internally it works by wrapping a slice and an RWMutex and offering methods for accessing and manipulating data.
type SyncedSlice[T any] struct {
	internal_slice []T
	mutex          sync.RWMutex
}

// Method for appending data to the slice.
// It takes the parameter data and appends it to the slice.
func (synced_slice *SyncedSlice[T]) Append(data T) {
	synced_slice.mutex.Lock()
	defer synced_slice.mutex.Unlock()
	synced_slice.internal_slice = append(synced_slice.internal_slice, data)
}

// Function for creating a new empty slice.
func New[T any]() SyncedSlice[T] {
	return SyncedSlice[T]{
		internal_slice: []T{},
		mutex: 			sync.RWMutex{},
	}
}

// Function for creating a Synced Slice from an existing slice.
// It is generally not recommended to use an existing slice as a parameter, as this allows for circumventing mutex locks.
func From[T any](s []T) SyncedSlice[T] {
	return SyncedSlice[T]{
		internal_slice: s,
		mutex:          sync.RWMutex{},
	}
}

// Method for cloning the internal slice of a Synced Slice.
// It will return the internal slice, which the struct stores.
// The clone is shallow, which means that if you were to, say, store pointers, a cloned map's pointers would still point to the same object.
func (s *SyncedSlice[T]) Clone() []T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return slices.Clone(s.internal_slice)
}

// Method for getting data from the slice by index.
// Get takes the index you want to access as a parameter and, if possible, returns the value at that index.
// It returns T and bool, where T is the value at that index and bool reports if the access was successful or not.
func (synced_slice *SyncedSlice[T]) Get(index int) (T, bool) {
	synced_slice.mutex.RLock()
	defer synced_slice.mutex.RUnlock()
	if len(synced_slice.internal_slice) <= int(index) {
		var t T
		return t, false
	}

	return synced_slice.internal_slice[index], true
}

// Method for clearing the internal slice which the struct holds.
func (synced_slice *SyncedSlice[T]) Clear() {
	synced_slice.mutex.Lock()
	defer synced_slice.mutex.Unlock()
	synced_slice.internal_slice = []T{}
}

// Method for deleting an entry at an index.
// Returns bool, which reports if the delete went well or not.
func (synced_slice *SyncedSlice[T]) Delete(index int) bool {
	synced_slice.mutex.Lock()
	defer synced_slice.mutex.Unlock()
	if len(synced_slice.internal_slice) <= index || index < 0 {
		return false
	}

	synced_slice.internal_slice = slices.Delete(synced_slice.internal_slice, index - 1, index)
	return true
}

// Method for getting the length of the slice.
func (synced_slice *SyncedSlice[T]) Len() int {
	synced_slice.mutex.RLock()
	defer synced_slice.mutex.RUnlock()
	return len(synced_slice.internal_slice)
}
