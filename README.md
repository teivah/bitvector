# teivah/bitvector

[![Go Report Card](https://goreportcard.com/badge/github.com/teivah/bitvector)](https://goreportcard.com/report/github.com/teivah/bitvector)

## Overview

A bit vector is an array data structure that compactly stores bits.

This library is based on 4 static different data structures:
* 8-bit vector: relies on an internal `uint8`
* 16-bit vector: relies on an internal `uint16`
* 32-bit vector: relies on an internal `uint32`
* 64-bit vector: relies on an internal `uint64`

The rationale of using a static integer compared to a dynamic `[]byte` is first of all to save memory.
There is no structure and/or slice overhead.
Hence, you might be interested in this library for memory-bound computation.

Also, the operations (get, set, etc.) are way more efficient. 
A simple benchmark shows that it's about 10 times more efficient than using a byte slice.
Moreover, there is a guarantee that the internal bit vectors will not escape to the heap and remain only at the stack level.

Yet, the only drawback is to have a fixed-size bit vector (8, 16, 32 or 64). If you require a dynamic bit vector, you should take a look at [dropbox/godropbox](https://github.com/dropbox/godropbox/tree/master/container/bitvector) for example.

## Installation

```
go get github.com/teivah/bitvector
```

## Documentation

### Initialization

* 8-bit vector: 

```go
bv := bitvector.New8()
```

* 16-bit vector: 

```go
bv := bitvector.New16()
```

* 32-bit vector: 

```go
bv := bitvector.New32()
```

* 64-bit vector: 

```go
bv := bitvector.New64()
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

* Convert the internal bit vector structure to a string:

```go
s := bv.String() // string
```