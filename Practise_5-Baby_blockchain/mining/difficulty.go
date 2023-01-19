package mining

import (
	"math/big"
)

func GenBits() {

}

func BitsToTarget(bits string) *big.Int {
	shift, _ := new(big.Int).SetString(bits[6:], 16)
	shift.Sub(shift, big.NewInt(3))
	shift.Mul(shift, big.NewInt(8))
	powBase := big.NewInt(2)
	shift = powBase.Exp(powBase, shift, nil)

	target, _ := new(big.Int).SetString(bits[2:6]+bits[:2], 16)
	target.Mul(target, shift)
	return target
}
