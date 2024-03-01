# Go Data Units

[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/daegalus/go-dataunits/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/daegalus/go-dataunits)](https://goreportcard.com/report/github.com/daegalus/go-dataunits)
[![GoDoc](https://godoc.org/github.com/daegalus/go-dataunits?status.svg)](https://godoc.org/github.com/daegalus/go-dataunits)

This is a library that provides utility functions for working with various data units in Go. Primarily for Storage sizes and Memory sizes.

## Features

- Conversion between different data units (bytes, kilobytes, megabytes, etc.)
- Formatting data sizes for display
- Parsing data sizes from strings

## Usage

To use the library, use the following command:

```shell
go get github.com/daegalus/go-dataunits
```

then import it

```go
import (
  "github.com/daegalus/go-dataunits"
)
```

## Example

```go
package main

import (
 "fmt"
 "github.com/daegalus/go-dataunits"
)

func main() {
 ds, err := dataunits.ParseSize[storage.StorageUnit]("1.5GiB")
 if err != nil {
  fmt.Println("Error parsing size:", err)
  return
 }

 fmt.Println("Parsed size:", ds.Size) // 1.5
 fmt.Println("Parsed size in bytes:", ds.Bytes) // 1610612736.0
 fmt.Println("Parsed Unit in bytes:", ds.Unit) // 1073741824 (GB in bytes)
}
```
