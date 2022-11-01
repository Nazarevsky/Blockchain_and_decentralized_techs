package main

import (
	"fmt"
	"pract3/hashing"
)

func main() {
	message := "sha"
	fmt.Println(hashing.SHA1(message))
}
