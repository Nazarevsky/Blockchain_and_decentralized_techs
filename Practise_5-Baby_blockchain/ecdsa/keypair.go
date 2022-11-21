package ecdsa

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

var symbols string = "0123456789abcdef"

func bigToHex(num *big.Int) string {
	return fmt.Sprintf("%x", num)
}

func hexToBig(hex string) *big.Int {
	val, _ := big.NewInt(1).SetString(hex, 16)
	return val
}

func GenPrivKey() *big.Int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	hexPriv := ""
	hexPriv += string(symbols[r.Intn(15)+1])

	for i := 1; i < 64; i++ {
		hexPriv += string(symbols[r.Intn(16)])
	}
	return hexToBig(hexPriv)
}

func PrivKeyToString(key *big.Int) string {
	return bigToHex(key)
}
