package blockchain

import (
	"bavovnacoin/hashing"
	"bavovnacoin/transaction"
	"bavovnacoin/utxo"
	"fmt"
	"time"
)

var Blockchain []Block

type Block struct {
	Id            uint
	Blocksize     uint
	Version       uint
	HashPrevBlock string
	Time          time.Time
	// Bits and Nonce fields - depends on Bavovnacoin concept thw will be choosen further
	// Bits          uint
	// Nonce         uint64
	Transactions []transaction.Transaction
}

func BlockToString(block Block) string {
	str := ""
	str += fmt.Sprint(block.Id)
	str += fmt.Sprint(block.Blocksize)
	str += fmt.Sprint(block.Version)
	str += block.HashPrevBlock
	str += block.Time.String()
	for i := 0; i < len(block.Transactions); i++ {
		str += transaction.GetCatTxFields(block.Transactions[i])
	}
	return str
}

func AddBlockToBlockchain(block Block) bool {
	isBlockValid := ValidateBlock(block)
	if isBlockValid {
		for i := 0; i < len(block.Transactions); i++ {
			txInpList := block.Transactions[i].Inputs
			//transaction.PrintTransaction(block.Transactions[i])
			for j := 0; j < len(txInpList); j++ {
				//println("Del")
				utxo.DelFromUtxo(txInpList[j].HashAdr, txInpList[j].OutInd)
			}

			txOutList := block.Transactions[i].Outputs
			for j := 0; j < len(txOutList); j++ {
				//println("Add")
				utxo.AddToUtxo(txOutList[j].HashAdr, txOutList[j].Sum)
			}
		}
		Blockchain = append(Blockchain, block)
	}
	return isBlockValid
}

func CreateBlock(id uint, txArr []transaction.Transaction, rewardAdr string) Block {
	var newBlock Block
	newBlock.Id = id

	if len(Blockchain) > 0 {
		newBlock.HashPrevBlock = hashing.SHA1(BlockToString(Blockchain[len(Blockchain)-1]))
	} else {
		newBlock.HashPrevBlock = "0000000000000000000000000000000000000000"
	}

	newBlock.Time = time.Now()
	newBlock.Transactions = make([]transaction.Transaction, len(txArr)+1)

	var coinbaseTx transaction.Transaction
	coinbaseTx.Outputs = append(coinbaseTx.Outputs, transaction.Output{HashAdr: rewardAdr, Sum: GetCoinsForEmition()})

	txArr = append([]transaction.Transaction{coinbaseTx}, txArr...)
	copy(newBlock.Transactions, txArr)
	newBlock.Blocksize = uint(len(BlockToString(newBlock)))
	return newBlock
}

func ValidateBlock(block Block) bool {
	lastBlockHash := hashing.SHA1(BlockToString(Blockchain[len(Blockchain)-1]))
	// Transaction are verified when added to mempool, no need to double check it

	if block.HashPrevBlock != lastBlockHash {
		return false
	}

	return true
}

func InitBlockchain() {
	genesisBlock := CreateBlock(0,
		nil, "0284cbd0bcf8a34035b71c5a72e37924cb960aaa0b69df4c41d50628734b8e1408")
	genesisBlock.Time = time.Date(2022, 12, 31, 11, 54, 22, 0, time.Local)
	Blockchain = append(Blockchain, genesisBlock)
}
