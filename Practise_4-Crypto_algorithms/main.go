package main

import (
	"crypt/cryption"
	"fmt"
)

func main() {
	cr := cryption.VigenereEncode("AtTACKATDAWN", "LEMONE")
	dcr := cryption.VigenereDecode(cr, "LEMONE")
	fmt.Println(cr)
	fmt.Println(dcr)
}
