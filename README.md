# go-autoinc

Package **autoinc** implements a _serial_ `AUTO_INCREMENT` type, for the Go programming language.

This is similar to the `AUTO_INCREMENT` found in many databases.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-autoinc

[![GoDoc](https://godoc.org/github.com/reiver/go-autoinc?status.svg)](https://godoc.org/github.com/reiver/go-autoinc)

## Examples

Here is an example **autoincional-type** that can hold a string:
```go
import "github.com/reiver/go-autoinc"

//

var serial autoinc.AutoInc[uint64]

// ...

next, err := serial.Next()

```

## Import

To import package **autoinc** use `import` code like the following:
```
import "github.com/reiver/go-autoinc"
```

## Installation

To install package **autoinc** do the following:
```
GOPROXY=direct go get github.com/reiver/go-autoinc
```

## Author

Package **autoinc** was written by [Charles Iliya Krempeaux](http://reiver.link)
