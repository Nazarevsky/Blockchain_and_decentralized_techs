package main

// if we write a library, we can't use 0x at the beginning!!
import (
	"fmt"
	"math"
	"math/big"
)

func reverseHexNum(value string, startNum int) string { // startNum if we have 0x at the beginning
	valueChars := []rune(value)
	valLen := int(math.Round(float64(len(value)) / float64(2)))
	for i := startNum; i < valLen; i += 2 {
		valueChars[len(value)-i], valueChars[i] = valueChars[i], valueChars[len(value)-i]
		valueChars[len(value)-i+1], valueChars[i+1] = valueChars[i+1], valueChars[len(value)-i+1]
	}
	return string(valueChars)
}

func hexToLittleEndian(value string) *big.Int {
	revStr := reverseHexNum(value, 2)
	answ, err := big.NewInt(0).SetString(revStr[2:], 16)

	if !err {
		panic("Помилка! Конвертується не число.")
	}
	return answ
}

func hexToBigEndian(value string) *big.Int {
	answer, err := big.NewInt(0).SetString(value, 16) // remake 0x should not get in!
	if !err {
		panic("Помилка! Конвертується не число.")
	}

	return answer
}

func bigintToHex(n *big.Int) string {
	return fmt.Sprintf("%x", n) // or %x or upper case
}

func littleEndianToHex(value *big.Int) string {
	hexVal := bigintToHex(value)
	hexVal = reverseHexNum(hexVal, 1)
	return hexVal
}

func main() {
	str := "0x0d0f000000000000000000000000000000000000000000000000000000000000"
	answ := hexToLittleEndian(str)
	fmt.Println("Answer 1: " + answ.String())

	str = littleEndianToHex(answ)
	fmt.Println("Answer 2: " + str)
}
