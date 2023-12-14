# go-collections [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
## Overview
This is util library for Go.  
The idea is from Python's collections library.  

## How to install
```
go get github.com/marubontan/go-collections
```

## How to use
```Go
package main

import (
	"fmt"

	counter "github.com/marubontan/go-collections/counter"
	defaultdict "github.com/marubontan/go-collections/defaultdict"
)

func main() {
	// Counter example
	sampleSlice := []int{1, 1, 2, 2, 2, 3, 3, 3, 3}
	counter, err := counter.NewCounter[int](sampleSlice)
	if err == nil {
		fmt.Println("Most common items:", counter.MostCommon(2))
	}

	// Defaultdict example
	dd := defaultdict.NewDefaultDict[string, int]()
	dd.Data["a"] += 1
	fmt.Println("Defaultdict value a:", dd.Get("a"))
	fmt.Println("Defaultdict value b:", dd.Get("b"))
}

```