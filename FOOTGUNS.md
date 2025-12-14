# Footguns
This library is meant to make go safer by offering thread-safe data structures, however, there are still ways to shoot yourself in the foot.

This file serves to properly document these footguns and warn users of this library of patterns which may cause bugs.

## Storing Pointers
This is arguably the most common type of footgun that can happen. 

While Types like AtomicMap or AtomicSlice themselves are thread-safe, what they store may not be. Therefore, if you must store pointers, use AtomicBox.

### Unsafe Example
```go
type MyStruct struct {
    s string
    i int
}

mySlice := atomicslice.From([]*MyStruct {
	&MyStruct{"Hello", 6},
	&MyStruct{"World", 7},
})

myInstance, _ := mySlice.Get(1) // Unsafe! Pointer without mutex!
```

### Safe Example
```go
type MyStruct struct {
    s string
    i int
}

mySlice := atomicslice.From([]*atomicbox.AtomicBox[MyStruct] {
	atomicbox.New(MyStruct{"Hello", 6}),
	atomicbox.New(MyStruct{"World", 7}),
})

myBox, ok := mySlice.Get(1) // Safe, AtomicBox protects its content.
```

## From Constructor
The From-Constructor is useful for when you want to instantiate a type with starting values, but it can also be misused for keeping a leaked reference to the inside-value.

### Unsafe Example
```go
myMap := map[string]int {
    "One": 1,
    "Two": 2,
    "Three": 3,
}

myAtomicMap := atomicmap.From(myMap) // Reference lives outside of struct.
myMap["Four"] = 4 // Mutex Bypass!
```

Therefore, you should just use an inline declaration like this:

### Safe Example
```go
myAtomicMap := atomicmap.From(map[string]int {
    "One": 1,
    "Two": 2,
    "Three": 3,
})
// No outside-living reference.
```

## Leaked AtomicBox Pointers
The references which AtomicBox stores can easily be leaked to an outside scope.

### Unsafe
```go
myBox := atomicbox.New("Hello World")

var myLeakedPointer *string
myBox.WithLock(func(inner *string) {
    myLeakedPointer = inner // Internal pointer leaked to the outside.
})
```

There is no safe counterpart to this as this is always a really bad pattern.

If you mean to clone or copy the value which the pointer stores, then that depends on the type itself.

## Any other footguns?
Probably, these are just the footguns I discovered. If you discover any more, please let me know by posting an issue!
