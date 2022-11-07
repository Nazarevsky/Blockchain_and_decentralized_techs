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
	cr := cryption.RSA(4564564, E, n)
	fmt.Println(cr, "crypted")
	fmt.Println(cryption.RSA(cr, D, n), "decrypted")
}
