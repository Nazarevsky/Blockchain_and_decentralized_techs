package main

import (
	"crypt/cryption"
	"fmt"
)

func main() {
	// cr := cryption.VigenereEncode("AtTACKATDAWN", "LEMONE")
	// dcr := cryption.VigenereDecode(cr, "LEMONE")
	// fmt.Println(cr)
	// fmt.Println(dcr)

	n, E, D := cryption.RSA_keygen()
	println(n, E, D)
	cr := cryption.RSA_Encrypt(123, E, n)
	fmt.Println(cr, "crypted")
	fmt.Println(cryption.RSA_Encrypt(cr, D, n), "decrypted")
	//fmt.Println(cryption.RSA())
}
