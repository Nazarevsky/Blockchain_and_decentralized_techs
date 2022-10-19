package main

// if we write a library, we can't use 0x at the beginning!!
import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

func addZeroToLen(str string, length int, state bool) string {
	concat := strings.Repeat("0", length-len(str))

	if state {
		return str + concat
	} else {
		return concat + str
	}
}

func reverseHexNum(value string) string { // startNum if we have 0x at the beginning
	valueChars := []rune(value)
	valLen := int(math.Round(float64(len(value)) / float64(2)))
	for i := 0; i < valLen; i += 2 {
		valueChars[len(value)-i-2], valueChars[i] = valueChars[i], valueChars[len(value)-i-2]
		valueChars[len(value)-i-1], valueChars[i+1] = valueChars[i+1], valueChars[len(value)-i-1]
	}

	return string(valueChars)
}

func hexToLittleEndian(value string) *big.Int {
	revStr := reverseHexNum(value)
	answ, err := big.NewInt(0).SetString(revStr, 16)

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
	str := fmt.Sprintf("%x", n)
	length := 32
	for ; len(str) > length; length *= 2 {
	}

	return addZeroToLen(str, length, true)
}

func littleEndianToHex(value *big.Int) string {
	hexVal := bigintToHex(value)
	hexVal = reverseHexNum(hexVal)
	return hexVal
}

func bigEndianToHex(value *big.Int) string {
	return fmt.Sprintf("%x", value)
}

func main() {
	str := "a0db0000000000000000000000000000"
	fmt.Println("Value: " + str)
	litEnd := hexToLittleEndian(str)
	bigEnd := hexToBigEndian(str)
	fmt.Println("Little Endian: " + litEnd.String())
	fmt.Println("Big Endian: " + bigEnd.String())

	str = littleEndianToHex(litEnd)
	fmt.Println("Hex from Little Endian: " + str)

	str = bigEndianToHex(bigEnd)
	fmt.Println("Hex from Big Endian: " + str)
}
