package main

import (
	"fmt"
	"pract3/hashing"
)

var sha1_const = []int{0x67452301, 0xEFCDAB89, 0x98BADCFE, 0x10325476, 0xC3D2E1F0}

func main() {
	message := "I can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all day"
	fmt.Println(len(message))
	fmt.Println(hashing.SHA1(message))
}
