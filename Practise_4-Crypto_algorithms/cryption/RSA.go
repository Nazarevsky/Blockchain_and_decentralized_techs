package cryption

import (
	"math"
	"math/rand"
	"time"
)

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
	println(p, q)
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

func RSA_Encrypt(mes, E, n uint64) uint64 {
	return pow_mod_Knuth(mes, E, n)
}

func RSA_Decrypt(crmes, D, n uint64) uint64 {
	return pow_mod_Knuth(crmes, D, n)
}
