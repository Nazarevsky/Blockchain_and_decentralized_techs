package account

import (
	"bavovnacoin/cryption"
	"bavovnacoin/ecdsa"
	"bavovnacoin/hashing"
	"fmt"
	"sort"
)

type Account struct {
	Id          string
	HashPass    string
	KeyPairList []ecdsa.KeyPair
	ArrId       int `json:"-"`
	Balance     int `json:"-"`
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

func AddKeyPairToAccount(arrInd int, password string) string {
	if Wallet[arrInd].HashPass == hashing.SHA1(password) {
		ecdsa.InitValues()
		newKeyPair := ecdsa.GenKeyPair()
		newKeyPair.PrivKey = cryption.AES_encrypt(newKeyPair.PrivKey, password)
		Wallet[arrInd].KeyPairList = append(Wallet[arrInd].KeyPairList, newKeyPair)
		WriteAccounts()
	} else {
		return "Wrong password!"
	}
	return ""
}

// Future realization: transaction creation.
// A stub operation just for demo of account work
func CreatePaymentOp(accountId int, recieverPubKey string, sum int, pass string) string {
	if Wallet[accountId].HashPass != hashing.SHA1(pass) {
		return "Wrong password"
	}
	getBalance(accountId)
	if sum > Wallet[accountId].Balance {
		return "Not enough coins to make a payment operation"
	}
	if sum < 0 {
		return "Incorrect value of sum of the operation"
	}

	var accUtxo []UTXO
	for i := 0; i < len(Wallet[accountId].KeyPairList); i++ {
		for j := 0; j < len(utxoList); j++ {
			if utxoList[j].PubKey == Wallet[accountId].KeyPairList[i].PublKey {
				accUtxo = append(accUtxo, utxoList[j])
			}
		}
	}
	sort.Slice(accUtxo, func(i, j int) bool {
		return accUtxo[i].Sum < accUtxo[j].Sum
	})

	tempSum := 0
	for i := 0; i < len(accUtxo); i++ {
		DelFromUtxo(accUtxo[i].Id)
		if tempSum+accUtxo[i].Sum > sum {
			tempSum += accUtxo[i].Sum
			utxoList = append(utxoList, UTXO{PubKey: recieverPubKey, Sum: sum})
			if tempSum-sum != 0 {
				accKeys := Wallet[accountId].KeyPairList
				AddKeyPairToAccount(accountId, pass) // generate new keypair for the change
				utxoList = append(utxoList, UTXO{Id: utxoList[len(utxoList)-1].Id + 1, PubKey: Wallet[accountId].KeyPairList[len(accKeys)].PublKey, Sum: tempSum - sum})
			}
			return ""
		}
		tempSum += accUtxo[i].Sum
	}

	return ""
}

func getKeyBal(pubKey string, arrInd int) int {
	bal := 0
	for j := 0; j < len(utxoList); j++ {
		if pubKey == utxoList[j].PubKey {
			bal += utxoList[j].Sum
		}
	}
	return bal
}

// A function counts all the UTXOs that is on specific public keys on user's account
func getBalance(arrInd int) {
	Wallet[arrInd].Balance = 0
	for i := 0; i < len(Wallet[arrInd].KeyPairList); i++ {
		Wallet[arrInd].Balance += getKeyBal(Wallet[arrInd].KeyPairList[i].PublKey, arrInd)
	}
}

func PrintBalance(arrInd int) {
	getBalance(arrInd)
	var bal float64 = float64(Wallet[arrInd].Balance) / 100000000.
	fmt.Printf("Balance: %.8f BVC\n", bal)
}

func SignData(accountInd int, hashMes string, kpInd int, pass string) (string, bool) {
	if Wallet[accountInd].HashPass != hashing.SHA1(pass) {
		return "", true
	}
	kp := Wallet[accountInd].KeyPairList[kpInd]
	kp.PrivKey = cryption.AES_decrypt(kp.PrivKey, pass)

	return ecdsa.Sign(hashMes, kp.PrivKey), false
}

func VerifData(accountInd int, hashMes string, kpInd int, signature string) bool {
	kp := Wallet[accountInd].KeyPairList[kpInd]
	return ecdsa.Verify(kp.PublKey, signature, hashMes)
}
