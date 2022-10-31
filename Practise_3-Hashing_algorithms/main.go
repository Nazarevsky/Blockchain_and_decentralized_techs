package main

import (
	"fmt"
	"pract3/hashing"
)

func main() {
	message := "I can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all dayI can literally talk all day"
	fmt.Println(len(message))
	fmt.Println(hashing.SHA1(message))
}
