package main

import (
	"fmt"
	"math/big"
)

func hexToLittleEndian(value string) *big.Int {
	valueChars := []rune(value)
	for i := 2; i < len(value)/2; i += 2 {
		valueChars[len(value)-i], valueChars[i] = valueChars[i], valueChars[len(value)-i]
		valueChars[len(value)-i+1], valueChars[i+1] = valueChars[i+1], valueChars[len(value)-i+1]
	}
	answ, err := big.NewInt(0).SetString(string(valueChars)[2:], 16)

	if !err {
		panic("Помилка! Конвертується не число.")
	}
	return answ
}

func hexToBigEndian(value string) *big.Int {
	answer, err := big.NewInt(0).SetString(value, 16) // remake 0x should not get in!
	if !err {
		panic("Error when converting!")
	}

	return answer
}

func main() {
	str := "0x0d0f000000000000000000000000000000000000000000000000000000000000"
	answ := hexToLittleEndian(str)
	fmt.Println("Answer: " + answ.String())
}
