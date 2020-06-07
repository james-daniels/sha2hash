# sha2hash

Lesson Excercise 4.2 from [The Go Programming Language](http://www.gopl.io/) book.

### Usage:
Build the binary
```Bash
$ go build -o sha2hash
```

Run the program
```Bash
$ ./sha2hash 56
7688b6ef52555962d008fff894223582c484517cea7da49ee67800adc7fc8866
```

OR
```Bash
$ ./sha2hash 56 384
0c5fb0ea6eba72f2dcbe1985e4bd011132c5d099f2486f63e3a8f554ebe50e6fae1e2b0e454b695b5acd318a7aa6d5e9
```

Options are 256 (Default), 384, and 512.


```Go
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	var output interface{}

	if len(os.Args) == 1 {
		fmt.Println("Enter value to convert.")
		os.Exit(0)
	}

	if len(os.Args) == 2 {
		output = sha256.Sum256([]byte(os.Args[1]))
	} else {
		switch os.Args[2] {
		case "256":
			output = sha256.Sum256([]byte(os.Args[1]))
		case "384":
			output = sha512.Sum384([]byte(os.Args[1]))
		case "512":
			output = sha512.Sum512([]byte(os.Args[1]))
		default:
			fmt.Println("Only 256 (Default), 384, and 512.")
			os.Exit(0)
		}
	}
	fmt.Printf("%x\n", output)
}
```