package main

import (
	"bavovnacoin/account"
	"bavovnacoin/blockchain"
	"bavovnacoin/ecdsa"
	"bavovnacoin/hashing"
	"bavovnacoin/transaction"
	"bavovnacoin/utxo"
	"fmt"
	"math/big"
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

// Function to demonstate account and wallet operations
// (for Step 3 baby blockchain realization)
func walletAndAccountExample() {
	// account.GenAccount("abc1")
	// account.GenAccount("abc2")
	// account.GenAccount("abc3")
	// account.GenAccount("abc4")
	utxo.InitUTXOValues()                  // Getting UTXO
	account.InitAccountsData()             // Getting account data from json file
	isAccExist := account.InitAccount("1") // Initialization account to work with
	//account.AddKeyPairToAccount(accInd, "abc1")
	if isAccExist {
		account.PrintBalance()
		res := transaction.CreatePaymentOp("02db1c6c791978448ef671746fab54797495597730b9a68721fca780db002162f0", 10000, "abc1")
		if res != "" {
			println(res)
		} else {
			account.PrintBalance()
		}

		mes := "Transaction here"
		hashMes := hashing.SHA1(mes)
		sign, err := account.SignData(hashMes, 0, "abc1")
		if err {
			println("Wrong password")
		} else {
			verifRes := account.VerifData(hashMes, 0, sign)
			if verifRes {
				println("Transaction is verified")
			} else {
				println("Transaction is not verified")
			}
		}
	} else {
		println("Such an account is not exists")
	}
}

// Function to demonstate transaction operations
// (for Step 4 baby blockchain realization)
func transactionExample() {
	account.InitAccountsData() // Getting account data from json file
	utxo.InitUTXOValues()      // Getting UTXO
	isAccExists := account.InitAccount("1")
	account.PrintBalance()
	if isAccExists {
		tx, resMes := transaction.CreateTransaction("abc1",
			[]string{"0284cbd0bcf8a34035b71c5a72e37924cb960aaa0b69df4c41d50628734b8e1408",
				"02487318bd34dc641708fd3300e2f5cca03cebe9ce7bea8973a40fc5b383de951d"},
			[]uint64{458, 200},
			14)
		if resMes == "" {
			transaction.PrintTransaction(tx)
		} else {
			println(resMes)
		}
		println()

		txVerifRes := transaction.VerifyTransaction(tx)
		if txVerifRes {
			println("Transaction is valid")
		} else {
			println("Transaction is not valid")
		}
	} else {
		println("Account is not found")
	}
}

func findNonce() {
	mes := "Hello world"
	target := new(big.Int)
	target.SetString("0000400000000000000000000000000000000000", 16)
	hashmesNonceBig := new(big.Int)
	hashmesNonceBig.SetString(hashing.SHA1(mes+fmt.Sprint(0)), 16)
	n := 1
	//println(hashing.SHA1(mes + fmt.Sprint(n)))
	for ; hashmesNonceBig.Cmp(target) == 1; n++ {
		hashmesNonceBig.SetString(hashing.SHA1(mes+fmt.Sprint(n)), 16)
	}
	println(n)
	println(fmt.Sprintf("%x", target))
	println(fmt.Sprintf("%x", hashmesNonceBig))
}

func main() {
	//findNonce()
	blockchain.CreateBlock(0, "", nil, "")
}
