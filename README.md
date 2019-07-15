# teivah/bitvector

[![Go Report Card](https://goreportcard.com/badge/github.com/teivah/bitvector)](https://goreportcard.com/report/github.com/teivah/bitvector)

## Overview

A bit vector is an array data structure that compactly stores bits.

`teivah/bitvector` is based on 4 different data structures:
* 8-bit vector: based on the Go `uint8` type 
* 16-bit vector: based on the Go `uint16` type
* 32-bit vector: based on the Go `uint32` type
* 64-bit vector: based on the Go `uint64` type

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
bv.Set(i, true)
```

* Get ith bit:

```go
b := bv.Get(i)
```

* Toggle ith bit:

```go
bv.Toggle(i)
```

* Clear bits from index i (included) to index j (excluded):

```go
bv.Clear(i, j)
```

* Count the number of bits set to 1:

```go
i := bv.Count()
```

* Reset the internal bit vector:

```go
bv.Reset()
```

* Copy the internal bit vector to another bit vector structure:

```go
bv2 := bv.Copy()
```

* Convert the internal bit vector structure to a string:

```go
s := bv.String()
```