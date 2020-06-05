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
		fmt.Println("Enter string to convert")
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
			output = sha256.Sum256([]byte(os.Args[1]))
		}
	}
	fmt.Printf("%x\n", output)
}
