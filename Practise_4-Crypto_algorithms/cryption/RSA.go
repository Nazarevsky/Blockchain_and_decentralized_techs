package cryption

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var symbols string = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-=[];',./*-+!@#$%^&*()`~йцукенгшщзхъфывапролджэячсмитьбюёіїЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЯЧСМИТЬБЮІЇÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜÝÞßàáâãäåæçèéêëìíîïóô"
var lower_bound uint64 = 999
var higher_bound uint64 = 9999

func isPrime(num uint64) bool {
	if num < 0 || num%2 == 0 {
		return false
	}

	var sqrtNum uint64 = uint64(math.Sqrt(float64(num)))
	var n uint64 = 3
	for ; n <= sqrtNum; n++ {
		if num%n == 0 {
			return false
		}
	}

	return true
}

func genRandPrime() uint64 {
	rand.Seed(time.Now().UnixNano())
	random := rand.Uint64() % ((higher_bound - lower_bound) + lower_bound)

	for ; !isPrime(random); random++ {
	}

	return random
}

func gcd(n1, n2 uint64) uint64 {
	for n1 != 0 && n2 != 0 {
		if n1 > n2 {
			n1 %= n2
		} else {
			n2 %= n1
		}
	}
	return n1 + n2
}

func RSA_keygen() (uint64, uint64, uint64) {
	p := genRandPrime()
	q := genRandPrime()

	for p == q {
		p = genRandPrime()
		q = genRandPrime()
	}

	n := p * q
	fi := (p - 1) * (q - 1)
	var E, D uint64

	for E = 3; gcd(E, fi) != 1; E += 2 {
	}
	for D = 3; (E*D)%fi != 1; D += 1 {
	}

	return n, E, D
}

func pow_mod_Knuth(val, pow, mod uint64) uint64 {
	var res uint64 = 1
	for pow != 0 {
		if (pow & 1) == 1 {
			res = ((res % mod) * (val % mod)) % mod
		}

		val = ((val % mod) * (val % mod)) % mod
		pow >>= 1
	}

	return res
}

func strToDec(mes string) []uint64 {
	binMes := ""
	for _, c := range mes {
		binMes = fmt.Sprintf("%s%b", binMes, c)
	}

	var decMesArr []uint64
	for i := 0; i < len(binMes)/64+1; i++ {
		if len(binMes)-(i*64+64) > 0 {
			val, _ := strconv.ParseUint(binMes[i*64:i*64+64], 2, 64)
			decMesArr = append(decMesArr, val)
		} else {
			val, _ := strconv.ParseUint(binMes[:len(binMes)%64], 2, 64)
			decMesArr = append(decMesArr, val)
		}
	}
	return decMesArr
}

func pad(mes string, c int) string {
	return strings.Repeat("0", c) + mes
}

func decToStr(crArr []uint64) string {
	binMes := ""
	hexMes := ""
	for i := 0; i < len(crArr); i++ {
		binStr := fmt.Sprintf("%b", crArr[i])
		if len(binStr)%16 != 0 {
			binStr = pad(binStr, 16-len(binStr)%16)
		}
		binMes += binStr
	}
	println(binMes)
	for i := 0; i < len(binMes)/16; i++ {
		ui, _ := strconv.ParseUint(binMes[i*16:i*16+16], 2, 16)
		hexMes += fmt.Sprintf("%x", ui)
	}
	return hexMes
}

func RSA_Encrypt(mes string, key, n uint64) string {
	var mesIntArr []uint64 = strToDec(mes)
	var crArr []uint64

	for i := 0; i < len(mesIntArr); i++ {
		crArr = append(crArr, pow_mod_Knuth(mesIntArr[i], key, n))
	}

	return decToStr(crArr)
}

func RSA_Decrypt(mes string, key, n uint64) string {
	var mesIntArr []uint64 = strToDec(mes)
	var crArr []uint64

	for i := 0; i < len(mesIntArr); i++ {
		crArr = append(crArr, pow_mod_Knuth(mesIntArr[i], key, n))
	}

	return decToStr(crArr)
}
