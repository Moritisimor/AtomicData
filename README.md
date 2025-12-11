# AtomicData
A Go Library for thread-safe collections and other data types.

## Why does this project exist?
During the time I used Rust, I grew aware of how many hidden dangers lurk in virtually all other languages, not just Go, and as such, I decided to make concurrency easier with this library, at least for Go.

At their cores, the types which this library provides are simply structs with a value and a mutex, unless the value is atomic, then it is without mutex.

The types offer a public API for interacting with the values which they encapsulate.

Fair warning: "atomic" refers to logical atomicity, meaning threads cannot interrupt each other, not atomic as in atomic CPU instructions.

## What types are there?
Since this library is in development, types may always be added, but currently the list looks like this:
- ```AtomicMap```
- ```AtomicSlice```
- ```AtomicStack```
- ```AtomicCounter```
- ```AtomicBox```

## How do I use it?
First, you will need to import the module. To do this, run ```go get github.com/Moritisimor/AtomicData``` in the folder of your project.

And that was basically it, you can now import the packages into go files and use the types. A small example:
```go
import github.com/Moritisimor/AtomicData/pkg/atomicmap
```
Or you can let the IDE figure it out like we all do.

## Practical Examples

### AtomicMap
```go
mymap := atomicmap.New[string, int]()
mykey := "one"
mymap.Set(mykey, 1)

val, ok := mymap.Get(mykey)
if !ok {
    fmt.Println("Key not found!")
} else {
    fmt.Printf("The Key %s holds the value %d\n", mykey, val)
}
```

### AtomicSlice
```go
myslice := atomicslice.New[int]()
wg := sync.WaitGroup{}
for i := range 200 {
    wg.Add(1)
    go func(n int) { // Start some goroutines to prove thread-safety
        defer wg.Done()
        myslice.Append(n)
    }(i)
}

wg.Wait()
for i := range myslice.Len() {
    val, ok := myslice.Get(i)
    if !ok {
        fmt.Println("Accessed out-of-bounds index!")
        return
    }

    fmt.Printf("Index %d holds value: %d", i, val)
    // The results will most likely not be ordered.
}

mysortedclone := myslice.Clone()
slices.Sort(mysortedclone)
for i, n := range mysortedclone {
	fmt.Printf("Index %d holds value: %d\n", i, n)
	// Sorted now.
}
```

### AtomicStack
```go
mystack := atomicstack.New[float64]()
for i := range 1000 {
	mystack.Push(math.Pi * float64(i))
}

fmt.Println(mystack.Peek())
fmt.Println(mystack.Pop())
fmt.Println(mystack.Peek())
```

### AtomicCounter
```go
mycounter := atomiccounter.New()
I := 20000
D := 10000
wg := sync.WaitGroup{}
for range I {
	wg.Go(func() {
		mycounter.Increment()
	})
}

for range D {
	wg.Go(func() {
		mycounter.Decrement()	
	})
}

wg.Wait()
if mycounter.Get() != int64(I - D) {
	fmt.Println("Counter's value is not as expected!")
} else {
	fmt.Println("Counter's value is as expected!")
}
```

### AtomicBox
```go
type mystruct struct {
	s string
	i int
	f float32
}

wg := sync.WaitGroup{}
mybox := atomicbox.New(mystruct {
	s: "Goodbye World",
	i: 67,
	f: 6.7,
})

for range 100 {
	wg.Go(func() {
		mybox.WithLock(func(inner *mystruct) {
			inner.i++
		})
	})
}

for range 100 {
	wg.Go(func() {
		mybox.WithLock(func(inner *mystruct) {
			inner.f--
		})
	})
}

mybox.WithLock(func(inner *mystruct) {
	inner.s = "Hello thread-safe World!"
})

wg.Wait()
mybox.WithLock(func(inner *mystruct) {
	fmt.Printf("s: %s\ni: %d\nf: %f", inner.s, inner.i, inner.f)
})
```
