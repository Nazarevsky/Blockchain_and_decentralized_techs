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

func printBlock(block [][]byte) {
	for y := 0; y < 4; y++ {
		line := ""
		for x := 0; x < 4; x++ {
			line += fmt.Sprintf("%x ", block[y][x])
		}
		println(line)
	}
	println()
}

func addKey(block [][]byte, key [][]byte) [][]byte {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			block[y][x] ^= key[y][x] //statekey
		}
	}
	return block
}

func getDecByHex(val string) int {
	num, _ := strconv.ParseInt(val, 16, 8)
	return int(num)
}

func getValInBox(hex byte, box [][]byte) byte {
	hexStr := complete(fmt.Sprintf("%x", hex), 2)
	i, _ := strconv.ParseInt(string(hexStr[0]), 16, 8)
	j, _ := strconv.ParseInt(string(hexStr[1]), 16, 8)
	return box[i][j]
}

func subBytes(block [][]byte) [][]byte {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			block[y][x] = getValInBox(block[y][x], sbox)
			// hex := complete(fmt.Sprintf("%x", block[y][x]), 2)
			// block[y][x] = sbox[getDecByHex(string(hex[0]))][getDecByHex(string(hex[1]))]
		}
	}
	return block
}

func shiftArr(arr []byte, shift int) []byte {
	return append(arr[shift:], arr[:shift]...)
}

func shiftBlock(block [][]byte) [][]byte {
	for y := 1; y < 4; y++ {
		block[y] = shiftArr(block[y], y)
	}
	return block
}

func multHex(hex byte, mul byte) byte {
	if mul == 1 {
		return hex
	} else if mul == 2 {
		return getValInBox(hex, mul2)
	}
	return getValInBox(hex, mul3)
}

func mixColumns(block [][]byte) [][]byte {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			var val byte = 0
			for k := 0; k < 4; k++ {
				val ^= multHex(block[k][x], mulMatr[y][k])
			}
			block[x][y] = val
		}
	}
	return block
}

func genRoundKey(key [][]byte, round int) [][]byte {
	w3 := []byte{0, 0, 0, 0}
	w3[0] = key[0][3]

	// get G(w3) and save w3
	tempW := getValInBox(key[0][3], sbox)
	for i := 1; i < len(key); i++ {
		w3[i] = key[i][3]
		key[i-1][3] = getValInBox(key[i][3], sbox)
	}
	key[3][3] = tempW
	key[0][3] ^= rcon[round]

	// xor words
	for w := 0; w < len(key); w++ {
		for i := 0; i < len(key); i++ {
			if w != len(key)-1 {
				key[i][w] ^= key[i][(3+w)%4]
			} else {
				key[i][3] = w3[i] ^ key[i][2]
			}
		}
	}

	return key
}

func AES_crypt(mes string, key string) string {
	bitMes := pad(mesToBits(mes), 128)
	divMesIntoBlocks(bitMes, len(bitMes)/128)

	bitKey := pad(mesToBits(key), 128)
	intitKey(bitKey)

	//addKey()
	//subBytes()
	//shiftBlock()
	//mixColumns()
	//genRoundKey()
	//printBlock()

	return bitMes
}
