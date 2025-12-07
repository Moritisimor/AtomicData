package atomiccounter

import "sync/atomic"

// A wrapper over an Atomic 64-bit Integer.
// This counter is thread-safe and mutex-free since it uses atomic operations.
type AtomicCounter struct {
	internalcounter atomic.Int64
}

// This function will make a new Synced Counter, with its internal counter set to 0.
func New() *AtomicCounter {
	return &AtomicCounter{}
}

// This function is like New, except that the counter will be set to the parameter of the function.
func At(i int64) *AtomicCounter {
	temp := AtomicCounter{}
	temp.internalcounter.Store(i)
	return &temp
}

func (c *AtomicCounter) Get() int64 {
	return c.internalcounter.Load()
}

// Method for incrementing the counter.
// It will return the new value which the counter holds.
func (c *AtomicCounter) Increment() int64 {
	return c.internalcounter.Add(1)
}

// Method for incermenting the counter by a set number.
// Just like Increment, it will return the new value of the counter.
func (c *AtomicCounter) IncrementBy(i int64) int64 {
	return c.internalcounter.Add(i)
}

// Method for decrementing the counter.
// It will return the new value of the counter.
func (c *AtomicCounter) Decrement() int64 {
	return c.internalcounter.Add(-1)
}

// Method for decrementing the counter by a set number.
// Just like Decrement, it will return the new value of the counter.
func (c *AtomicCounter) DecrementBy(i int64) int64 {
	return c.internalcounter.Add(-i)
}

// Method for setting the counter back to 0.
func (c *AtomicCounter) Reset() {
	c.internalcounter.Store(0)
}
