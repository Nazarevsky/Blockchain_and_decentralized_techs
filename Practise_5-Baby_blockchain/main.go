package main

import (
	"bavovnacoin/account"
	"bavovnacoin/blockchain"
	"bavovnacoin/transaction"
	"bavovnacoin/utxo"
)

// Function to demonstate transaction operations
// (for Step 4 baby blockchain realization)
func transactionExample() {
	account.InitAccountsData() // Getting account data from json file
	utxo.InitUTXOValues()      // Getting UTXO
	isAccExists := account.InitAccount("1")
	account.PrintBalance()
	if isAccExists {
		tx, resMes := transaction.CreateTransaction("abc1",
			[]string{"ce390e1683d0fa81568d2042f6f84ac0b0bf0bb9",
				"52ccca777a1ab46ddd07257b7def18f3c33e1fec"},
			[]uint64{458, 200},
			14,
			0)
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

func main() {
	account.InitAccountsData()
	account.InitAccount("1")

	utxo.InitUTXOValues()
	utxo.ShowCoinDatabase()

	blockchain.InitBlockchain()

	tx, _ := transaction.CreateTransaction("abc1",
		[]string{"ce390e1683d0fa81568d2042f6f84ac0b0bf0bb9",
			"52ccca777a1ab46ddd07257b7def18f3c33e1fec"},
		[]uint64{458, 200},
		3,
		5)
	println(blockchain.AddTxToMempool(tx))

	/*
	 Locktime is 5-th block, so it won't be added until 5-th block appears.
	 Reward tx only will be added
	*/
	// tx, _ := transaction.CreateTransaction("abc1",
	// 	[]string{"ce390e1683d0fa81568d2042f6f84ac0b0bf0bb9",
	// 		"52ccca777a1ab46ddd07257b7def18f3c33e1fec"},
	// 	[]uint64{458, 200},
	// 	3,
	// 	5)
	// println(blockchain.AddTxToMempool(tx))

	tx, _ = transaction.CreateTransaction("abc1",
		[]string{"ce390e1683d0fa81568d2042f6f84ac0b0bf0bb9"},
		[]uint64{458, 200},
		1,
		0)
	println(blockchain.AddTxToMempool(tx)) // false, because this outputs are already in mempool

	newBlock := blockchain.CreateBlock(uint(len(blockchain.Blockchain)), blockchain.GetTransactionsFromMempool(),
		"0accd57e23dfd44edc1e8a60f8a0889aa4b277a8")

	isAdded := blockchain.AddBlockToBlockchain(newBlock)
	if isAdded {
		println("Block is added to blockchain")
	} else {
		println("Block is not added to blockchain")
	}
	println("\nChanged coin database:")
	utxo.ShowCoinDatabase()
}
