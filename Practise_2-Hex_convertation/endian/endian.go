package endian

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

func reverseHexNum(value string) string {
	valueChars := []rune(value)
	valLen := int(math.Round(float64(len(value)) / float64(2)))
	for i := 0; i < valLen; i += 2 {
		valueChars[len(value)-i-2], valueChars[i] = valueChars[i], valueChars[len(value)-i-2]
		valueChars[len(value)-i-1], valueChars[i+1] = valueChars[i+1], valueChars[len(value)-i-1]
	}

	return string(valueChars)
}

func HexToLittleEndian(value string) *big.Int {
	var revStr string
	revStr = value
	if len(value)%2 == 1 {
		revStr = addZeroToLen(value, len(value)+1, true)
	}
	revStr = reverseHexNum(revStr)
	answ, _ := big.NewInt(0).SetString(revStr, 16)
	return answ
}

func HexToBigEndian(value string) *big.Int {
	answer, _ := big.NewInt(0).SetString(value, 16)
	return answer
}

func bigIntToHex(n *big.Int) string {
	str := fmt.Sprintf("%x", n)
	length := 32
	for ; len(str) > length; length *= 2 {
	}

	return addZeroToLen(str, length, true)
}

func LittleEndianToHex(value *big.Int) string {
	hexVal := bigIntToHex(value)
	hexVal = reverseHexNum(hexVal)
	return hexVal
}

func BigEndianToHex(value *big.Int) string {
	return fmt.Sprintf("%x", value)
}
