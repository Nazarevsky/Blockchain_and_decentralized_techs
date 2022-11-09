package cryption

import (
	"fmt"
	"strconv"
	"strings"
)

var blocks [][][]byte
var stateKey [][]byte

var mulMatr [][]byte = [][]byte{
	{2, 3, 1, 1},
	{1, 2, 3, 1},
	{1, 1, 2, 3},
	{3, 1, 1, 2}}

func divMesIntoBlocks(bitMes string, countBlocks int) {
	blocks = make([][][]byte, countBlocks)
	ind := 0
	for i := 0; i < countBlocks; i++ {
		blocks[i] = make([][]byte, 4)
		for y := 0; y < 4; y++ {
			blocks[i][y] = make([]byte, 4)
			for x := 0; x < 4; x++ {
				num, _ := strconv.ParseInt(bitMes[ind:ind+8], 2, 8)
				blocks[i][y][x] = byte(num)
				ind += 8
			}
		}
	}
}

func intitKey(bitMes string) {
	stateKey = make([][]byte, 4)
	ind := 0
	for y := 0; y < 4; y++ {
		stateKey[y] = make([]byte, 4)
		for x := 0; x < 4; x++ {
			num, _ := strconv.ParseInt(bitMes[ind:ind+8], 2, 8)
			stateKey[y][x] = byte(num)
			ind += 8
		}
	}
}

func complete(mes string, to int) string {
	return strings.Repeat("0", to-len(mes)) + mes
}

func pad(mes string, to int) string {
	return strings.Repeat("0", to-(len(mes)%to)) + mes
}

func mesToBits(mes string) string {
	str := ""
	for _, c := range mes {
		str += fmt.Sprintf("%b", c)
	}
	return str
}

func printBlock() {
	for i := 0; i < len(blocks); i++ {
		for y := 0; y < 4; y++ {
			line := ""
			for x := 0; x < 4; x++ {
				line += fmt.Sprintf("%d ", blocks[i][y][x])
			}
			println(line)
		}
		println()
	}
}

func addKey() {
	for i := 0; i < len(blocks); i++ {
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				blocks[i][y][x] ^= stateKey[y][x]
			}
		}
	}
}

func getDecByHex(val string) int {
	num, _ := strconv.ParseInt(val, 16, 8)
	return int(num)
}

func subBytes() {
	for i := 0; i < len(blocks); i++ {
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				hex := complete(fmt.Sprintf("%x", blocks[i][y][x]), 2)
				blocks[i][y][x] = sbox[getDecByHex(string(hex[0]))][getDecByHex(string(hex[1]))]
			}
		}
	}
}

func shiftArr(arr []byte, shift int) []byte {
	return append(arr[shift:], arr[:shift]...)
}

func shiftBlock() {
	for i := 0; i < len(blocks); i++ {
		for y := 1; y < 4; y++ {
			blocks[i][y] = shiftArr(blocks[i][y], y)
		}
	}
}

func multHex(hex byte, mul byte) byte {
	if mul == 1 {
		return hex
	}
	hexStr := complete(fmt.Sprintf("%x", hex), 2)
	i, _ := strconv.ParseInt(string(hexStr[0]), 16, 8)
	j, _ := strconv.ParseInt(string(hexStr[1]), 16, 8)

	if mul == 2 {
		return mul2[i][j]
	}
	return mul3[i][j]
}

func mixColumns() {
	var m [][]byte = [][]byte{
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0xbf, 0xb4, 0x41, 0x27},
		{0x5d, 0x52, 0x11, 0x98},
		{0x30, 0xae, 0xf1, 0xe5}}
	var res [][]byte = [][]byte{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			for k := 0; k < 4; k++ {
				res[y][x] ^= multHex(m[k][x], mulMatr[y][k])
			}
		}
	}

}

func AES_crypt(mes string, key string) string {
	bitMes := pad(mesToBits(mes), 128)
	// divMesIntoBlocks(bitMes, len(bitMes)/128)

	// bitKey := pad(mesToBits(key), 128)
	// intitKey(bitKey)

	//addKey()
	//subBytes()
	//shiftBlock()
	//printBlock()
	mixColumns()
	//printBlock()
	//println(fmt.Sprintf("%b", 0x1554^0x11b))
	return bitMes
}
