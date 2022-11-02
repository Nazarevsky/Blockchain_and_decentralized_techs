package hashing

// Keccak-512, 1600 bit 24 rounds

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// var rc = []uint64{0x0000000000000001, 0x0000000000008082, 0x800000000000808A, 0x8000000080008000,
// 	0x000000000000808B, 0x0000000080000001, 0x8000000080008081, 0x8000000000008009,
// 	0x000000000000008A, 0x0000000000000088, 0x0000000080008009, 0x000000008000000A,
// 	0x000000008000808B, 0x800000000000008B, 0x8000000000008089, 0x8000000000008003,
// 	0x8000000000008002, 0x8000000000000080, 0x000000000000800A, 0x800000008000000A,
// 	0x8000000080008081, 0x8000000000008080, 0x0000000080000001, 0x8000000080008008}

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
		return m - (num*-1)%m
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
	for z := 0; z < 64; z++ {
		x, y := 1, 0
		for t := 0; t < 24; t++ {
			fmt.Println((z - (t+1)*(t+2)/2), mod((z-(t+1)*(t+2)/2), 64))
			a[x][y][z] = a[x][y][(mod((z - (t+1)*(t+2)/2), 64))]
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

func replAtInd(in string, r byte, i int) string {
	out := []rune(in)
	n, _ := strconv.Atoi(string(r))
	out[i] = rune(n)
	return string(out)
}

func rc(t int) byte {
	if t%255 == 0 {
		return 1
	}
	R := "10000000"
	for i := 0; i < t%255; i++ {
		R = "0" + R
		R = replAtInd(R, R[0]^R[8], 0)
		R = replAtInd(R, R[4]^R[8], 4)
		R = replAtInd(R, R[5]^R[8], 5)
		R = replAtInd(R, R[6]^R[8], 6)
		R = R[:len(R)-1]
	}
	n, _ := strconv.Atoi(string(R[0]))
	return byte(n)
}

func iota_(a [][][]byte, ir int) [][][]byte {
	var RC []byte = make([]byte, 64)

	for j := 0; j < 6; j++ {
		RC[int(math.Pow(2., float64(j)))-1] = rc(j + 7*ir)
	}

	for k := 0; k < 64; k++ {
		a[0][0][k] ^= RC[k]
	}

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

		for ir := 12 + 2*6 - 24; ir != 12+2*6-1; ir-- { // change to constants?
			sponge = iota_(hi(pi(rho(theta(sponge)))), ir)
		}

		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				for z := 0; z < 64; z++ {
					fmt.Println(sponge[x][y][z])
				}
			}
		}

	}

	return ""
}
