# teivah/bitvector

[![Go Report Card](https://goreportcard.com/badge/github.com/teivah/bitvector)](https://goreportcard.com/report/github.com/teivah/bitvector)

## Overview

A bit vector is an array data structure that compactly stores bits.

This library is based on 5 static different data structures:
* 8-bit vector: relies on an internal `uint8`
* 16-bit vector: relies on an internal `uint16`
* 32-bit vector: relies on an internal `uint32`
* 64-bit vector: relies on an internal `uint64`
* 128-bit vector: relies on two internal `uint64` (for ASCII problems)

The rationale of using a static integer compared to a dynamic `[]byte` is first of all to save memory.
There is no structure and/or slice overhead.
Hence, you might be interested in this library for memory-bound computation.

Also, the operations (get, set, etc.) are way more efficient. 
A simple benchmark shows that it's about 10 times more efficient than using a byte slice.
Moreover, there is a guarantee that the internal bit vectors will not escape to the heap and remain only at the stack level.

Yet, the only drawback is to have a fixed-size bit vector (8, 16, 32, 64 or 128). If you require a dynamic bit vector, you should take a look at [dropbox/godropbox](https://github.com/dropbox/godropbox/tree/master/container/bitvector) for example.

## Installation

```
go get github.com/teivah/bitvector
```

## Documentation

### Initialization

* 8-bit vector: 

```go
var bv bitvector.Len8
```

* 16-bit vector: 

```go
var bv bitvector.Len16
```

* 32-bit vector: 

```go
var bv bitvector.Len32
```

* 64-bit vector: 

```go
var bv bitvector.Len64
```

* 128-bit vector: 

```go
var bv bitvector.Ascii
// Or to reinitialize the bit vector
bv = bitvector.NewAscii()
```

### Operations

* Set ith bit:

```go
bv = bv.Set(i, true)
bv = bv.Set(i, false)
```

* Get ith bit:

```go
b := bv.Get(i) // bool
```

* Toggle (flip) ith bit:

```go
bv = bv.Toggle(i)
```

* Clear bits from index i (included) to index j (excluded):

```go
bv = bv.Clear(i, j)
```

* Count the number of bits set to 1:

```go
i := bv.Count() // uint8
```

* And operator:

```go
bv := bv1.And(bv2)
```

* Or operator:

```go
bv := bv1.Or(bv2)
```

* Xor operator:

```go
bv := bv1.Xor(bv2)
```

* AndNot operator:

```go
bv := bv1.AndNot(bv2)
```

* Push operator (left shift):

```go
bv = bv.Push(2)
```


```go
bv = bv.Pop(2)
```

* Convert the internal bit vector structure to a string:

```go
s := bv.String() // string
```