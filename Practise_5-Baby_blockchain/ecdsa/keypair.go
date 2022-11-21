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

var a *big.Int
var b *big.Int
var p big.Int
var n big.Int

type Point struct {
	x, y big.Int
}

var G Point

func modInv(a, m *big.Int) *big.Int {
	mOrig := m
	y := big.NewInt(0)
	x := big.NewInt(1)

	bigOne := big.NewInt(1)
	for a.Cmp(bigOne) == 1 {
		q := big.NewInt(0).Div(a, m)
		temp := m

		m = big.NewInt(0).Rem(a, m)
		a = temp // remove???
		temp = y

		qmuly := big.NewInt(0).Mul(q, y)
		y = big.NewInt(0).Sub(x, qmuly)
		x = temp
	}

	if x.Cmp(big.NewInt(0)) == -1 {
		x = big.NewInt(0).Add(x, mOrig)
	}
	return x
}

func GenPublKey() { //key *big.Int
	a = big.NewInt(0)
	a = big.NewInt(7)
	// P NOT CORRECT!!!
	p.SetString("497323236409786642155382248146820840100456150797347717440463976893159497012533375532079", 10)
	n.SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)

	x, _ := big.NewInt(1).SetString("55066263022277343669578718895168534326250603453777594175500187360389116729240", 10)
	y, _ := big.NewInt(1).SetString("32670510020758816978083085130507043184471273380659243275938904335757337482424", 10)
	G.x = *x
	G.y = *y

	println(modInv(big.NewInt(13), big.NewInt(47)).String())
}
