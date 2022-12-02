package main

import (
	"bavovnacoin/account"
	"bavovnacoin/ecdsa"
	"bavovnacoin/hashing"
)

// Function for demonstrating key pair and signature generation
// (from Step 2 baby blockchain realization)
func ecdsaExample() {
	ecdsa.InitValues()
	message := "Here is a transaction in Bavovnacoin."
	hash := hashing.SHA1(message)

	// The loop here is just to show that signing and verification work fine for different keys.
	for i := 0; i < 10; i++ {
		priv_key := ecdsa.GenPrivKey()
		pub_key := ecdsa.GenPubKey(priv_key)

		sign := ecdsa.Sign(hash, priv_key)
		ver := ecdsa.Verify(pub_key, sign, hash)

		if ver {
			println("Message verified successfuly!")
		} else {
			println("Message is not verified.")
		}
	}
}

func main() {
	// account.GenAccount("abc1")
	// account.GenAccount("abc2")
	// account.GenAccount("abc3")
	// account.GenAccount("abc4")
	account.InitUTXOValues()             // Getting UTXO
	account.InitAccountsData()           // Getting account data from json file
	accInd := account.GetAccountInd("1") // Choosing account to work with
	//account.AddKeyPairToAccount(accInd, "abc1")

	account.PrintBalance(accInd)
	account.PrintBalance(3)
	res := account.CreatePaymentOp(accInd, "02db1c6c791978448ef671746fab54797495597730b9a68721fca780db002162f0", 10000, "abc1")
	if res != "" {
		println(res)
	} else {
		account.PrintBalance(accInd)
		account.PrintBalance(3)
	}

	mes := "Transaction here"
	hashMes := hashing.SHA1(mes)
	sign, err := account.SignData(accInd, hashMes, 0, "abc1")
	if err {
		println("Wrong password")
	} else {
		verifRes := account.VerifData(accInd, hashMes, 0, sign)
		if verifRes {
			println("Transaction is verified")
		} else {
			println("Transaction is not verified")
		}
	}
}
