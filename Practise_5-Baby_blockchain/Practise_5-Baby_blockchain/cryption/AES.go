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

func addKey(block [][]byte, key [][]byte) [][]byte {
	for y := 0; y < len(block); y++ {
		for x := 0; x < len(block); x++ {
			block[y][x] ^= key[y][x]
		}
	}
	return block
}

func getValInBox(hex byte, box [][]byte) byte {
	hexStr := complete(fmt.Sprintf("%x", hex), 2)
	i, _ := strconv.ParseInt(string(hexStr[0]), 16, 8)
	j, _ := strconv.ParseInt(string(hexStr[1]), 16, 8)
	return box[i][j]
}

func subBytes(block [][]byte, box [][]byte) [][]byte {
	for y := 0; y < len(block); y++ {
		for x := 0; x < len(block); x++ {
			block[y][x] = getValInBox(block[y][x], box)
		}
	}
	return block
}

func shiftArr(arr []byte, shift int) []byte {
	return append(arr[shift:], arr[:shift]...)
}

func shiftBlock(block [][]byte) [][]byte {
	for y := 1; y < len(block); y++ {
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

	for y := 0; y < len(block); y++ {
		for x := 0; x < len(block); x++ {
			var val byte = 0
			for k := 0; k < len(block); k++ {
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
	for y := 0; y < len(block); y++ {
		for x := 0; x < len(block); x++ {
			res += complete(fmt.Sprintf("%x", block[x][y]), 2)
		}
	}
	return res
}

// crypt

func crToBlocks(blocks [][][]byte, hexMes string, countBlocks int) [][][]byte {
	blocks = make([][][]byte, countBlocks)
	ind := 0
	for i := 0; i < countBlocks; i++ {
		blocks[i] = make([][]byte, 4)
		for j := 0; j < len(blocks[i]); j++ {
			blocks[i][j] = make([]byte, 4)
		}
		for y := 0; y < len(blocks[i]); y++ {
			for x := 0; x < len(blocks[i]); x++ {
				num, _ := strconv.ParseInt(hexMes[ind:ind+2], 16, 64)
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
	for y := 1; y < len(block); y++ {
		block[y] = shiftArr(block[y], len(block)-y)
	}
	return block
}

func hexToString(block [][]byte) string {
	res := ""
	for y := 0; y < len(block); y++ {
		for x := 0; x < len(block); x++ {
			res += string(block[x][y])
		}
	}
	return res
}

func AES_encrypt(mes string, key string) string {
	var blocks [][][]byte
	var stateKey [][]byte
	var roundKeys [][][]byte
	res := ""

	bitMes := pad(mesToBits(mes), 128)
	blocks = divMesIntoBlocks(blocks, bitMes, len(bitMes)/128)
	bitKey := pad(mesToBits(key), 128)
	stateKey = intitKey(stateKey, bitKey)

	roundKeys = append(roundKeys, blockNewInstance(stateKey))
	for i := 0; i < 10; i++ {
		stateKey = genRoundKey(stateKey, i)
		roundKeys = append(roundKeys, blockNewInstance(stateKey))
	}

	for i := 0; i < len(blocks); i++ {
		addKey(blocks[i], roundKeys[0])

		for r := 0; r < 9; r++ {
			blocks[i] = subBytes(blocks[i], sbox)
			blocks[i] = shiftBlock(blocks[i])
			blocks[i] = mixColumns(blocks[i], mulMatr)
			blocks[i] = addKey(blocks[i], roundKeys[r+1])
		}

		blocks[i] = subBytes(blocks[i], sbox)
		blocks[i] = shiftBlock(blocks[i])
		blocks[i] = addKey(blocks[i], roundKeys[len(roundKeys)-1])

		res += reassemble(blocks[i])
	}

	return res
}

func AES_decrypt(mes string, key string) string {
	var blocks [][][]byte
	var stateKey [][]byte
	var roundKeys [][][]byte

	blocks = crToBlocks(blocks, mes, len(mes)/32)
	bitKey := pad(mesToBits(key), 128)
	stateKey = intitKey(stateKey, bitKey)
	roundKeys = append(roundKeys, blockNewInstance(stateKey))
	res := ""

	for i := 0; i < 10; i++ {
		stateKey = genRoundKey(stateKey, i)
		roundKeys = append(roundKeys, blockNewInstance(stateKey))
	}

	for i := 0; i < len(blocks); i++ {
		blocks[i] = addKey(blocks[i], roundKeys[len(roundKeys)-1])
		blocks[i] = invShiftBlock(blocks[i])
		blocks[i] = subBytes(blocks[i], invsbox)
		for r := 8; r >= 0; r-- {
			blocks[i] = addKey(blocks[i], roundKeys[r+1])
			blocks[i] = mixColumns(blocks[i], invMulMatr)
			blocks[i] = invShiftBlock(blocks[i])
			blocks[i] = subBytes(blocks[i], invsbox)
		}
		blocks[i] = addKey(blocks[i], roundKeys[0])
		res += hexToString(blocks[i])
	}
	return res
}
