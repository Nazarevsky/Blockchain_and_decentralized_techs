package main

import (
	"sign/ecdsa"
)

func main() {
	//println(ecdsa.PrivKeyToString(ecdsa.GenPrivKey()))
	ecdsa.GenPublKey()
}
