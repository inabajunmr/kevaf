# kevaf

[![GoDoc](https://godoc.org/github.com/inabajunmr/kevaf?status.svg)](https://godoc.org/github.com/inabajunmr/kevaf)
[![codecov](https://codecov.io/gh/inabajunmr/kevaf/branch/master/graph/badge.svg)](https://codecov.io/gh/inabajunmr/kevaf)
![CircleCI](https://circleci.com/gh/inabajunmr/kevaf/tree/master.svg?style=svg)

Lightweight file-base KVS for golang.

# Install
```
go get github.com/inabajunmr/kevaf
```

# Usage
```go
package main

import (
	"fmt"
	"os"

	kevaf "github.com/inabajunmr/kevaf"
)

func main() {

	// prepare dir for kevaf
	os.Mkdir("./.kvs", 0777)
	kvs := kevaf.FileMap{"./.kvs"}

	// put new value
	err := kvs.Put("key", []byte("value"))
	if err != nil {
		fmt.Println("Failed to put.", err)
		os.Exit(1)
	}

	// get put value
	val, err := kvs.Get("key")
	if err != nil {
		fmt.Println("key is empty")
		os.Exit(1)
	}

	fmt.Println(string(val)) // value
}
```
