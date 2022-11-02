package hashing

// Keccak-512

import (
	"fmt"
	"strconv"
	"strings"
)

func binaryK(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

// func addZerosToK(mes string, to int) string {
// 	return strings.Repeat("0", to-len(mes)) + "1"
// }

func padding(binMes string) string { // padding maybe 0x80 at the end
	length := len(binMes)
	return binMes + "1" + strings.Repeat("0", 1600-((length+2)%1600)) + "1"
}

func mod(num int, m int) int {
	if num < 0 {
		return m - (num*-1)%5
	}
	return num % m
}

func theta(a [][][]byte) [][][]byte { // really weird: same value as a
	var c [][]byte = make([][]byte, 5)
	var d [][]byte = make([][]byte, 5)
	for i := 0; i < 5; i++ {
		c[i] = make([]byte, 64)
		d[i] = make([]byte, 64)
	}

	for x := 0; x < 5; x++ {
		for z := 0; z < 64; z++ {
			temp := a[x][0][z]
			for y := 1; y < 4; y++ {
				temp ^= a[x][y][z]
			}
		}
	}

	for x := 0; x < 5; x++ {
		for z := 0; z < 64; z++ {
			d[x][z] = c[mod((x-1), 5)][z] ^ c[(x+1)%5][mod((z-1), 64)]
		}
	}

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < 64; z++ {
				a[x][y][z] ^= d[x][z]
			}
		}
	}

	return a
}

func rho(a [][][]byte) [][][]byte {
	for z := 0; z < 24; z++ {
		x, y := 1, 0
		for t := 0; t < 24; t++ {
			a[x][y][z] = a[x][y][(z-(t+1)*(t+2)/2)%64]
			x, y = y, (2*x+3*y)%5
		}
	}
	return a
}

func pi(a [][][]byte) [][][]byte {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < 64; z++ {
				a[x][y][z] = a[(x+3*y)%5][x][z]
			}
		}
	}
	return a
}

func hi(a [][][]byte) [][][]byte {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < 64; z++ {
				a[x][y][z] = a[x][y][z] ^ ((a[(x+1)%5][y][z] ^ 1) * a[(x+2)%5][y][z])
			}
		}
	}
	return a
}

func iota_(a [][][]byte) [][][]byte {
	return a
}

func Keccak(messgae string) string {
	var sponge [][][]byte = make([][][]byte, 5) // sponge stores dec value
	for i := 0; i < 5; i++ {
		sponge[i] = make([][]byte, 5)
		for j := 0; j < 5; j++ {
			sponge[i][j] = make([]byte, 64)
		}
	}

	binaryMes := binaryK(messgae)
	padMes := padding(binaryMes)

	for i := 0; i < len(padMes)/1600; i += 1 {
		chunk := padMes[i*1600 : i*1600+1600]

		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				for z := 0; z < 64; z++ {
					n, _ := strconv.Atoi(string(chunk[64*(5*x+y)+z]))
					sponge[x][y][z] = byte(n)
				}
			}
		}
		sponge = theta(sponge)
		sponge = rho(sponge)
		sponge = pi(sponge)
		sponge = hi(sponge)
	}

	return ""
}
