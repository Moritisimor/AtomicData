// The atomicmap package contains functions, methods and types for the AtomicMap type.
// AtomicMaps themselves are thread-safe map implementations.
package atomicmap

import (
	"maps"
	"sync"
)

// A thread-safe map-wrapper.
// Internally it works by wrapping a map and a mutex and offering methods for interacting with the data.
type AtomicMap[K comparable, V any] struct {
	internalmap map[K]V
	mutex        sync.RWMutex
}

// Creates a new empty, thread-safe map.
// It returns a pointer to the AtomicMap-object in the heap.
func New[K comparable, V any]() *AtomicMap[K, V] {
	return &AtomicMap[K, V]{
		internalmap: map[K]V{},
		mutex:        sync.RWMutex{},
	}
}

// Creates a new thread-safe map from an existing map.
// The supplied parameter will be shallowly cloned, as such, storing pointers is not recommended.
// If you must store pointers, it is recommended to use AtomicBox
func From[K comparable, V any](m map[K]V) *AtomicMap[K, V] {
	return &AtomicMap[K, V]{
		internalmap: maps.Clone(m),
		mutex:        sync.RWMutex{},
	}
}

// Method for Clearing the Map.
func (m *AtomicMap[K, V]) Clear() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.internalmap = map[K]V{}
}

// Method for getting a value with a key.
// Returns V and bool, where V is the value of the key and bool represents if access was successful or not.
func (m *AtomicMap[K, V]) Get(key K) (V, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if v, ok := m.internalmap[key]; ok {
		return v, true
	}

	var v V
	return v, false
}

// Method for setting a value with a key.
// The value of the key is not immutable and overwriting data may happen.
// If overwrites are not allowed to happen, consider using SetIfNotExists instead.
func (m *AtomicMap[K, V]) Set(key K, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.internalmap[key] = value
}

// Method for setting value with a key if it does not exist.
// Similar to the regular Set method, but keys are treated as immutable.
// Returns a bool. True if the key was set, else false.
func (m *AtomicMap[K, V]) SetIfNotExists(key K, value V) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, ok := m.internalmap[key]; ok {
		return false
	}

	m.internalmap[key] = value
	return true
}

// Method for cloning the internal map.
// It will return the internal map, which the struct stores.
// The map is copied-by-value, not a reference.
// Since the clone is shallow, storing pointers is not recommended as those will still point to the same object.
func (m *AtomicMap[K, V]) Clone() map[K]V {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	temp := map[K]V{}
	maps.Copy(temp, (m.internalmap))
	return temp
}

// Method for updating a value with its key.
// This method is exclusively meant for explicitly modifying an existing key.
// If the key does not exist, it will not set the value either. 
// It returns bool, which reports whether an update went well or not.
func (m *AtomicMap[K, V]) Update(key K, value V) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, ok := m.internalmap[key]; ok {
		m.internalmap[key] = value
		return true
	}

	return false
}

func (m *AtomicMap[K, V]) Len() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.internalmap)
}

// Method for deleting value with a key.
// Returns bool, which reports if a delete went well or not.
func (m *AtomicMap[K, V]) Delete(key K) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if _, ok := m.internalmap[key]; ok {
		delete(m.internalmap, key)
		return true
	}

	return false
}
