# teivah/bitvector

[![Go Report Card](https://goreportcard.com/badge/github.com/teivah/bitvector)](https://goreportcard.com/report/github.com/teivah/bitvector)

## Overview

A bit vector is an array data structure that compactly stores bits.

This library is statically based on 4 different data structures:
* 8-bit vector: relies on an internal `uint8`
* 16-bit vector: relies on an internal `uint16`
* 32-bit vector: relies on an internal `uint32`
* 64-bit vector: relies on an internal `uint64`

The rationale compared to being based on a `[]byte` is to save memory if you need a static bit vector structure. Hence, you might be interested in this library for memory-bound computation.

Moreover, there is a guarantee that each the internal bit vectors will not escape to the heap and remain only at the stack level.

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