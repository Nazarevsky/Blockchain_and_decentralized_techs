package cryption

import (
	"fmt"
	"strconv"
	"strings"
)

func divMesIntoBlocks(blocks [][][]byte, bitMes string, countBlocks int) [][][]byte {
	blocks = make([][][]byte, countBlocks)
	ind := 0
	for i := 0; i < countBlocks; i++ {
		blocks[i] = make([][]byte, 4)
		for j := 0; j < 4; j++ {
			blocks[i][j] = make([]byte, 4)
		}
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				num, _ := strconv.ParseInt(bitMes[ind:ind+8], 2, 8)
				blocks[i][x][y] = byte(num)
				ind += 8
			}
		}
	}
	return blocks
}

func intitKey(stateKey [][]byte, bitMes string) [][]byte {
	stateKey = make([][]byte, 4)
	ind := 0
	for j := 0; j < 4; j++ {
		stateKey[j] = make([]byte, 4)
	}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			num, _ := strconv.ParseInt(bitMes[ind:ind+8], 2, 8)
			stateKey[x][y] = byte(num)
			ind += 8
		}
	}
	return stateKey
}

func complete(mes string, to int) string {
	return strings.Repeat("0", to-len(mes)) + mes
}

func pad(mes string, to int) string {
	if len(mes)%to == 0 {
		return mes
	}
	return strings.Repeat("0", to-(len(mes)%to)) + mes
}

func mesToBits(mes string) string {
	str := ""
	for _, c := range mes {
		str += complete(fmt.Sprintf("%b", c), 8)
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

func subBytes(block [][]byte, box [][]byte) [][]byte {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			block[y][x] = getValInBox(block[y][x], box)
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
	} else if mul == 3 {
		return getValInBox(hex, mul3)
	} else if mul == 9 {
		return getValInBox(hex, mul9)
	} else if mul == 11 {
		return getValInBox(hex, mulb)
	} else if mul == 13 {
		return getValInBox(hex, muld)
	}
	return getValInBox(hex, mule)
}

func mixColumns(block [][]byte, mulmatr [][]byte) [][]byte {
	res := [][]byte{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			var val byte = 0
			for k := 0; k < 4; k++ {
				val ^= multHex(block[k][x], mulmatr[y][k])
			}
			res[y][x] = val
		}
	}
	return res
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

func reassemble(block [][]byte) string {
	res := ""
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			res += complete(fmt.Sprintf("%x", block[x][y]), 2)
		}
	}
	return res
}

func AES_crypt(mes string, key string) string {
	var blocks [][][]byte
	var stateKey [][]byte

	bitMes := pad(mesToBits(mes), 128)
	blocks = divMesIntoBlocks(blocks, bitMes, len(bitMes)/128)

	bitKey := pad(mesToBits(key), 128)
	stateKey = intitKey(stateKey, bitKey)

	for i := 0; i < len(blocks); i++ {
		addKey(blocks[i], stateKey)
		for r := 0; r < 9; r++ {
			blocks[i] = subBytes(blocks[i], sbox)
			blocks[i] = shiftBlock(blocks[i])
			blocks[i] = mixColumns(blocks[i], mulMatr)
			stateKey = genRoundKey(stateKey, r)
			blocks[i] = addKey(blocks[i], stateKey)
		}

		blocks[i] = subBytes(blocks[i], sbox)
		blocks[i] = shiftBlock(blocks[i])
		stateKey = genRoundKey(stateKey, 9)
		blocks[i] = addKey(blocks[i], stateKey)

	}
	res := ""
	for i := 0; i < len(blocks); i++ {
		res += reassemble(blocks[0])
	}

	return res
}

func crToBlocks(blocks [][][]byte, hexMes string, countBlocks int) [][][]byte {
	blocks = make([][][]byte, countBlocks)
	ind := 0
	for i := 0; i < countBlocks; i++ {
		blocks[i] = make([][]byte, 4)
		for j := 0; j < 4; j++ {
			blocks[i][j] = make([]byte, 4)
		}
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				num, _ := strconv.ParseInt(hexMes[ind:ind+2], 16, 2)
				blocks[i][x][y] = byte(num)
				ind += 2
			}
		}
	}
	return blocks
}

func blockNewInstance(block [][]byte) [][]byte {
	res := [][]byte{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}

	for y := 0; y < len(block); y++ {
		for x := 0; x < len(block); x++ {
			res[y][x] = block[y][x]
		}
	}
	return res
}

func invShiftBlock(block [][]byte) [][]byte {
	for y := 1; y < 4; y++ {
		block[y] = shiftArr(block[y], 4-y)
	}
	return block
}

func AES_decrypt(mes string, key string) string {
	var blocks [][][]byte
	var stateKey [][]byte
	var roundKeys [][][]byte

	blocks = crToBlocks(blocks, mes, len(mes)/32)

	bitKey := pad(mesToBits(key), 128)
	stateKey = intitKey(stateKey, bitKey)
	roundKeys = append(roundKeys, blockNewInstance(stateKey))

	for i := 0; i < 10; i++ {
		stateKey = genRoundKey(stateKey, i)
		roundKeys = append(roundKeys, blockNewInstance(stateKey))
	}
	//println(len(blocks))
	for i := 0; i < len(blocks); i++ {
		blocks[i] = addKey(blocks[i], roundKeys[len(roundKeys)-1])
		for r := 9; r > 8; r-- { // i swap to 0???
			blocks[i] = invShiftBlock(blocks[i])
			blocks[i] = subBytes(blocks[i], invsbox)
			blocks[i] = addKey(blocks[i], roundKeys[r])
			printBlock(blocks[i])
			blocks[i] = mixColumns(blocks[i], invMulMatr)
			printBlock(blocks[i])
		}
		blocks[i] = invShiftBlock(blocks[i])
		blocks[i] = subBytes(blocks[i], invsbox)
		blocks[i] = addKey(blocks[i], roundKeys[0])
	}
	//printBlock(blocks[0])
	return ""
}
