package account

import (
	"bavovnacoin/cryption"
	"bavovnacoin/ecdsa"
	"bavovnacoin/hashing"
	"fmt"
	"sort"
)

var CurrAccount Account

type Account struct {
	Id          string
	HashPass    string
	KeyPairList []ecdsa.KeyPair
	ArrId       int    `json:"-"`
	Balance     uint64 `json:"-"`
}

type UtxoForInput struct {
	Address string
	Index   int
}

// Generates new account and set up a password to encode a private key
func GenAccount(password string) Account {
	ecdsa.InitValues()
	var newAcc Account

	newAcc.HashPass = hashing.SHA1(password)

	newKeyPair := ecdsa.GenKeyPair()
	newKeyPair.PrivKey = cryption.AES_encrypt(newKeyPair.PrivKey, password)

	newAcc.Id = fmt.Sprint(RightBoundAccNum + 1)
	RightBoundAccNum++
	newAcc.KeyPairList = append(newAcc.KeyPairList, newKeyPair)

	Wallet = append(Wallet, newAcc)
	WriteAccounts()
	return newAcc
}

func AddKeyPairToAccount(password string) string {
	if CurrAccount.HashPass == hashing.SHA1(password) {
		ecdsa.InitValues()
		newKeyPair := ecdsa.GenKeyPair()
		newKeyPair.PrivKey = cryption.AES_encrypt(newKeyPair.PrivKey, password)
		CurrAccount.KeyPairList = append(CurrAccount.KeyPairList, newKeyPair)
		Wallet[CurrAccount.ArrId] = CurrAccount
		WriteAccounts()
	} else {
		return "Wrong password!"
	}
	return ""
}

func GetAccUtxo() []UTXO {
	var accUtxo []UTXO
	for i := 0; i < len(CurrAccount.KeyPairList); i++ {
		for j := 0; j < len(UtxoList); j++ {
			if UtxoList[j].PubKey == CurrAccount.KeyPairList[i].PublKey {
				accUtxo = append(accUtxo, UtxoList[j])
			}
		}
	}
	sort.Slice(accUtxo, func(i, j int) bool {
		return accUtxo[i].Sum > accUtxo[j].Sum
	})
	return accUtxo
}

func getNextInpIndex(addressInp string, utxoInputs []UTXO, utxoInd int) int {
	ind := -1
	for i := 0; i <= utxoInd; i++ {
		if utxoInputs[i].PubKey == addressInp {
			ind++
		}
	}
	return ind
}

/*
Algorithm of effective transaction inputs search:
iterate utxo of a specific account and check two neighboring values.
At the beginning we add a stub UTXO (it will not be added to the database).
We are looking for a minimum value (checking left neighbor)
that is higher or equal to the required sum (minus sum that we have already found).
If a right neighbor is less than needed sum, we keep iterating, because there is a chance
of finding better variant.
*/
func GetTransInputs(sum uint64, accUtxo []UTXO) ([]UtxoForInput, []UTXO, uint64) {
	if accUtxo == nil {
		accUtxo = GetAccUtxo()
	}

	accUtxo = append(accUtxo, UTXO{PubKey: "", Sum: 0}) // Stub value for searching
	var utxoInput []UtxoForInput
	tempSum := uint64(0)

	if len(accUtxo) == 1 && accUtxo[0].Sum >= sum {
		return append(utxoInput, UtxoForInput{accUtxo[0].PubKey, getNextInpIndex(accUtxo[0].PubKey, accUtxo, 0)}),
			accUtxo, accUtxo[0].Sum
	}

	for i := 1; i < len(accUtxo); i++ {
		if accUtxo[i-1].Sum >= sum-tempSum {
			if sum-tempSum > accUtxo[i].Sum {
				utxoInput = append(utxoInput, UtxoForInput{Address: accUtxo[i-1].PubKey, Index: getNextInpIndex(accUtxo[i-1].PubKey, accUtxo, i-1)})
				return utxoInput, accUtxo, accUtxo[i-1].Sum + tempSum
			} else {
				continue
			}
		}
		utxoInput = append(utxoInput, UtxoForInput{accUtxo[i-1].PubKey, getNextInpIndex(accUtxo[i-1].PubKey, accUtxo, i-1)})
		tempSum += accUtxo[i-1].Sum
	}
	return nil, accUtxo, tempSum
}

// Future realization: transaction creation.
// A stub operation just for demo of account work
func CreatePaymentOp(recieverAdr string, sum uint64, pass string) string {
	if CurrAccount.HashPass != hashing.SHA1(pass) {
		return "Wrong password"
	}
	GetBalance()
	if sum > CurrAccount.Balance {
		return "Not enough coins to make a payment operation"
	}
	if sum < 0 {
		return "Incorrect value of sum of the operation"
	}

	accUtxo := GetAccUtxo()
	transInputs, accUtxo, inpSum := GetTransInputs(sum, accUtxo)
	if transInputs != nil {
		UtxoList = append(UtxoList, UTXO{PubKey: recieverAdr, Sum: sum}) // Send coins

		if inpSum-sum != 0 { // generating change
			accKeys := CurrAccount.KeyPairList
			AddKeyPairToAccount(pass) // generate new keypair for the change
			UtxoList = append(UtxoList, UTXO{Id: UtxoList[len(UtxoList)-1].Id + 1,
				PubKey: CurrAccount.KeyPairList[len(accKeys)].PublKey, Sum: inpSum - sum})
		}
	} else {
		return "Not enough coins for sending."
	}

	return ""
}

func GetBalByKeyHash(keyHash string, outInd int) uint64 {
	ind := -1
	for j := 0; j < len(UtxoList); j++ {
		if keyHash == hashing.SHA1(UtxoList[j].PubKey) {
			ind++
		}
		if ind == outInd {
			return UtxoList[j].Sum
		}
	}
	return 0
}

func getKeyBal(pubKey string) uint64 {
	bal := uint64(0)
	for j := 0; j < len(UtxoList); j++ {
		if pubKey == UtxoList[j].PubKey {
			bal += UtxoList[j].Sum
		}
	}
	return bal
}

// A function counts all the UTXOs that is on specific public keys on user's account
func GetBalance() uint64 {
	CurrAccount.Balance = 0
	for i := 0; i < len(CurrAccount.KeyPairList); i++ {
		CurrAccount.Balance += getKeyBal(CurrAccount.KeyPairList[i].PublKey)
	}
	return CurrAccount.Balance
}

func PrintBalance() {
	GetBalance()
	var bal float64 = float64(CurrAccount.Balance) / 100000000.
	fmt.Printf("Balance: %.8f BVC\n", bal)
}

func getAccountInd(accountId string) int {
	for i := 0; i < len(Wallet); i++ {
		if Wallet[i].Id == accountId {
			Wallet[i].ArrId = i
			return i
		}
	}
	return -1
}

func InitAccount(accountId string) bool {
	ecdsa.InitValues()
	accInd := getAccountInd(accountId)
	if accInd != -1 {
		CurrAccount = Wallet[accInd]
		return true
	}
	return false
}

func SignData(hashMes string, kpInd int, pass string) (string, bool) {
	if CurrAccount.HashPass != hashing.SHA1(pass) {
		return "", true
	}
	kp := CurrAccount.KeyPairList[kpInd]
	kp.PrivKey = cryption.AES_decrypt(kp.PrivKey, pass)

	return ecdsa.Sign(hashMes, kp.PrivKey), false
}

func VerifData(hashMes string, kpInd int, signature string) bool {
	kp := CurrAccount.KeyPairList[kpInd]
	return ecdsa.Verify(kp.PublKey, signature, hashMes)
}
