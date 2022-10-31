package main

import (
	"fmt"
	"pract3/hashing"
)

func main() {
	message := "Some text"
	fmt.Println(hashing.SHA1(message))
}
