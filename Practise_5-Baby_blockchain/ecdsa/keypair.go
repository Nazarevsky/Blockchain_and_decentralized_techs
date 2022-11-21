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
var p *big.Int
var n *big.Int

type Point struct {
	x, y *big.Int
}

var G Point

func inverse(a, m *big.Int) *big.Int {
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

func double(point Point) Point {
	// slope
	x2 := big.NewInt(0).Mul(point.x, point.x)
	x2m3 := big.NewInt(0).Mul(x2, big.NewInt(3))
	up := big.NewInt(0).Add(x2m3, a)
	ym2 := big.NewInt(0).Mul(point.y, big.NewInt(2))
	invym2 := inverse(ym2, p)
	down := big.NewInt(0).Mod(invym2, p)
	umd := big.NewInt(0).Mul(up, down)
	s := big.NewInt(0).Mod(umd, p)

	// x
	s2 := big.NewInt(0).Mul(s, s)
	xm2 := big.NewInt(0).Mul(point.x, big.NewInt(2))
	s2mxm2 := big.NewInt(0).Sub(s2, xm2)
	x := big.NewInt(0).Mod(s2mxm2, p)

	// y
	pxsx := big.NewInt(0).Sub(point.x, x)
	sm_pxsx := big.NewInt(0).Mul(s, pxsx)
	sm_pxsx_sy := big.NewInt(0).Sub(sm_pxsx, point.y)
	y := big.NewInt(0).Mod(sm_pxsx_sy, p)

	point.x = x
	point.y = y
	return point
}

func isPointEqual(p1, p2 Point) bool {
	if p1.x.Cmp(p2.x) == 0 && p1.y.Cmp(p2.y) == 0 {
		return true
	}
	return false
}

func add(p1, p2 Point) Point {
	if isPointEqual(p1, p2) { // try to check this: remove and run. If similar, it`s correct
		return double(p1)
	}
	var res Point
	// slope
	//slope = ((point1[:y] - point2[:y]) * inverse(point1[:x] - point2[:x], $p)) % $p
	up := big.NewInt(0).Sub(p1.y, p2.y)
	p1xsp2x := big.NewInt(0).Sub(p1.x, p2.x)
	down := inverse(p1xsp2x, p)
	umd := big.NewInt(0).Mul(up, down)
	s := big.NewInt(0).Mod(umd, p)
	println(s.String())
	return res
}

func GenPublKey() { //key *big.Int
	a = big.NewInt(0)
	a = big.NewInt(7)
	p, _ = big.NewInt(0).SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	n, _ = big.NewInt(0).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
	x, _ := big.NewInt(0).SetString("55066263022277343669578718895168534326250603453777594175500187360389116729240", 10)
	y, _ := big.NewInt(0).SetString("32670510020758816978083085130507043184471273380659243275938904335757337482424", 10)
	G.x = x
	G.y = y

	var a Point
	a.x = big.NewInt(1)
	a.y = big.NewInt(2)

	var b Point
	b.x = big.NewInt(1)
	b.y = big.NewInt(32)

	add(a, b)
}
