package hashing

import (
	"fmt"
	"strconv"
	"strings"
)

var sha1_h = []int{0x67452301, 0xEFCDAB89, 0x98BADCFE, 0x10325476, 0xC3D2E1F0}
var w []int64

func binary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

func addZerosTo(mes string, to int) string {
	return strings.Repeat("0", to-len(mes)) + mes
}

func filler(binMes string) string {
	length := len(binMes)
	binLen := binary(strconv.Itoa(length))
	return binMes + "1" + strings.Repeat("0", 512-((length+65)%512)) + addZerosTo(binLen, 64)
}

func SHA1(message string) string {
	binForm := binary(message)
	binForm = filler(binForm)

	for i := 0; i < len(binForm)%512; i += 0 {
		chunk := binForm[i*512 : i*512+512]
		for i := 0; i < 16; i++ {
			num, _ := strconv.ParseInt(chunk[i*32:i*32+32], 2, 64)
			w = append(w, num)
		}

		for i := 16; i < 80; i++ {
			w = append(w, (w[i-3]^w[i-8]^w[i-14]^w[i-16])<<1) // or left shift by 5????
		}

		arr := make([]int, len(sha1_h))
		copy(arr, sha1_h)
	}

	return filler(binForm)
}
