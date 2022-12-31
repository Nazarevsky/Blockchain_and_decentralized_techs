package blockchain

import (
	"bavovnacoin/transaction"
	"bavovnacoin/utxo"
)

var Mempool []transaction.Transaction

func ValidateTransaction(tx transaction.Transaction) bool {
	if !transaction.VerifyTransaction(tx) || tx.Locktime < uint(len(Blockchain)) && tx.Locktime != 0 {
		return false
	}

	for j := 0; j < len(tx.Inputs); j++ {
		for i := 0; i < len(Mempool); i++ { // Check same input in mempool (TODO: find more effective way)
			for k := 0; k < len(Mempool[i].Inputs); k++ {
				if Mempool[i].Inputs[k].HashAdr == tx.Inputs[j].HashAdr &&
					Mempool[i].Inputs[k].OutInd == tx.Inputs[j].OutInd { // hash address and ind
					return false
				}
			}
		}

		isExist := false
		for i := 0; i < len(utxo.UtxoList); i++ {
			if utxo.UtxoList[i].Address == tx.Inputs[j].HashAdr {
				isExist = true
			}
		}

		if isExist {
			return false
		}
	}
	return true
}

func AddTxToMempool(tx transaction.Transaction) bool {
	if ValidateTransaction(tx) {
		fee := transaction.GetTxFee(tx)
		insInd := findIndexSorted(fee, tx.Locktime)

		if len(Mempool) != 0 {
			if insInd < len(Mempool) {
				Mempool = append(Mempool[:insInd+1], Mempool[insInd:]...)
				Mempool[insInd] = tx
				return true
			} else {
				Mempool = append(Mempool, tx)
				return true
			}
		} else {
			Mempool = append(Mempool, tx)
			return true
		}
	}
	return false
}

// Make binary search???
func findIndexSorted(fee uint64, locktime uint) int {
	for i := 0; i < len(Mempool); i++ {
		txFee := transaction.GetTxFee(Mempool[i])
		if txFee == fee {
			if Mempool[i].Locktime < locktime {
				return i
			}
		}
		if txFee < fee { // add locktime check
			return i
		}
	}
	return len(Mempool)
}

func GetTransactionsFromMempool() []transaction.Transaction {
	var txForBlock []transaction.Transaction
	allSize := 0
	MempoolInd := 0

	for allSize < 1000000 && MempoolInd < len(Mempool) {
		allSize += transaction.ComputeTxSize(Mempool[MempoolInd])
		// Check locktime
		if Mempool[MempoolInd].Locktime < uint(len(Mempool)) {
			transaction.PrintTransaction(Mempool[MempoolInd])
			txForBlock = append(txForBlock, Mempool[MempoolInd])
			Mempool = append(Mempool[:MempoolInd], Mempool[MempoolInd+1:]...)
		} else {
			MempoolInd++
		}
	}
	return txForBlock
}
