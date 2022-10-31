package hashing

import (
	"fmt"
	"strconv"
	"strings"
)

var sha1_const = []int{0x67452301, 0xEFCDAB89, 0x98BADCFE, 0x10325476, 0xC3D2E1F0}

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
	return filler(binForm)
}
