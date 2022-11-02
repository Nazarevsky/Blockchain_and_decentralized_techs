package hashing

// Keccak-512

import (
	"fmt"
	"strings"
)

func binaryK(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

// func addZerosToK(mes string, to int) string {
// 	return strings.Repeat("0", to-len(mes)) + "1"
// }

func padding(binMes string) string {
	length := len(binMes)
	return binMes + "1" + strings.Repeat("0", 1600-((length+2)%1600)) + "1"
}

func Keccak(messgae string) string {
	binaryMes := binaryK(messgae)
	padMes := padding(binaryMes)

	for i := 0; i < len(padMes)/1600; i += 1 {
		chunk := padMes[i*1600 : i*1600+1600] //????
		fmt.Println(chunk)
		// make state array (the sponge) and fill it with 0
	}

	return ""
}
