package main

import (
	"fmt"
	"math"
	"strconv"
)

func hexToDecimal(hex string) {
	strconv.ParseInt(hex, 16, 64)
}

func hexToLittleEndian(value string) int64 {
	var answer int64
	for i := 2; i < len(value); i++ {
		num, err := strconv.ParseInt(string(value[i]), 16, 64)
		if err != nil {
			panic(err)
		}

		answer += num * int64(math.Pow(16, float64(i-2)))
	}
	return answer
}

func main() {
	str := "0xff00000000000000000000000000000000000000000000000000000000000000"
	answ := hexToLittleEndian(str)
	fmt.Println("Answer: " + strconv.FormatInt(answ, 10))
}
