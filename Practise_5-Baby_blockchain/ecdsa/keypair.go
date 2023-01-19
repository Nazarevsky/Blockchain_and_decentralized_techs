package ecdsa

import (
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

func bigToHex(num *big.Int) string {
	return fmt.Sprintf("%x", num)
}

func hexToBig(hex string) *big.Int {
	val, _ := big.NewInt(1).SetString(hex, 16)
	return val
}

func GenPrivKey() string {
	s := rand.NewSource(time.Now().UnixNano()) // Use entropy here
	r := rand.New(s)
	rrr := big.NewInt(1).Rand(r, n)
	return bigToHex(rrr)
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

type KeyPair struct {
	PrivKey, PublKey string
}

var g Point

type P struct {
	x, y int
}

func InitValues() {
	a = big.NewInt(0)
	b = big.NewInt(7)
	p, _ = big.NewInt(0).SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	n, _ = big.NewInt(0).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
	g.x, _ = big.NewInt(0).SetString("55066263022277343669578718895168534326250603453777594175500187360389116729240", 10)
	g.y, _ = big.NewInt(0).SetString("32670510020758816978083085130507043184471273380659243275938904335757337482424", 10)
}

func inverse(a, m *big.Int) *big.Int {
	prevY := big.NewInt(0)
	y := big.NewInt(1)

	if a.Cmp(big.NewInt(0)) == -1 {
		a = big.NewInt(0).Mod(a, m)
	}

	bigOne := big.NewInt(1)
	for a.Cmp(bigOne) == 1 {
		q := big.NewInt(0).Div(m, a)

		qmuly := big.NewInt(0).Mul(q, y)
		tempY := y
		y = big.NewInt(0).Sub(prevY, qmuly)
		prevY = tempY

		tempa := a
		a = big.NewInt(0).Rem(m, a)
		m = tempa
	}

	return y
}

func double(point Point) Point {
	// slope
	x2 := big.NewInt(0).Mul(point.x, point.x)
	x2m3 := big.NewInt(1).Mul(x2, big.NewInt(3))
	up := big.NewInt(0).Add(x2m3, a)
	ym2 := big.NewInt(0).Mul(point.y, big.NewInt(2))
	down := inverse(ym2, p)
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
	if isPointEqual(p1, p2) {
		return double(p1)
	}
	var res Point

	// slope
	up := big.NewInt(0).Sub(p1.y, p2.y)
	p1xsp2x := big.NewInt(0).Sub(p1.x, p2.x)
	down := inverse(p1xsp2x, p)
	umd := big.NewInt(0).Mul(up, down)
	s := big.NewInt(0).Mod(umd, p)

	// x
	s2 := big.NewInt(0).Mul(s, s)
	s2sp1x := big.NewInt(0).Sub(s2, p1.x)
	s2sp1xp2x := big.NewInt(0).Sub(s2sp1x, p2.x)
	x := big.NewInt(0).Mod(s2sp1xp2x, p)

	// y
	p1xsx := big.NewInt(0).Sub(p1.x, x)
	smp1xsx := big.NewInt(0).Mul(s, p1xsx)
	smp1xsxsp1y := big.NewInt(0).Sub(smp1xsx, p1.y)
	y := big.NewInt(0).Mod(smp1xsxsp1y, p)

	res.x = x
	res.y = y
	return res
}

func multiply(pk *big.Int, p Point) Point {
	currP := p
	pkBin := fmt.Sprintf("%b", pk)
	for i := 1; i < len(pkBin); i++ {
		currP = double(currP)
		if pkBin[i] == '1' {
			currP = add(currP, p)
		}
	}
	return currP
}

func fillZero(str string, needLen int) string {
	return strings.Repeat("0", needLen-len(str)) + str
}

func compressPubKey(p Point) string {
	xStrHex := fillZero(bigToHex(p.x), 64)
	if big.NewInt(0).Mod(p.y, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return "02" + xStrHex
	} else {
		return "03" + xStrHex
	}
}

func decompressPubKey(s string) Point {
	pref := s[:2]
	x, _ := big.NewInt(0).SetString(s[2:], 16)
	// y**2 = x**3 + b, a == 0
	ysquared := big.NewInt(0).Mod(big.NewInt(0).Add(big.NewInt(0).Exp(x, big.NewInt(3), nil), b), p)
	// y = (y**2)**((p+1)/4)
	y := big.NewInt(0).Exp(ysquared, big.NewInt(0).Div(big.NewInt(0).Add(p, big.NewInt(1)), big.NewInt(4)), p)
	if (pref == "02" && big.NewInt(0).Mod(y, big.NewInt(2)).Cmp(big.NewInt(0)) != 0) ||
		(pref == "03" && big.NewInt(0).Mod(y, big.NewInt(2)).Cmp(big.NewInt(0)) == 0) {
		y = big.NewInt(0).Mod(big.NewInt(0).Sub(p, y), p)
	}

	var p Point
	p.x = x
	p.y = y
	return p
}

func GenPubKey(privK string) string {
	privKey := hexToBig(privK)
	p := multiply(privKey, g)
	return compressPubKey(p)
}

func GenKeyPair() KeyPair {
	var kp KeyPair
	kp.PrivKey = GenPrivKey()
	kp.PublKey = GenPubKey(kp.PrivKey)
	return kp
}
