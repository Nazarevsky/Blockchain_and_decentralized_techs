package blockchain

import (
	"bavovnacoin/hashing"
	"bavovnacoin/transaction"
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

func AddBlockToBlockchain(block Block) bool {
	isBlockValid := ValidateBlock(block)
	if isBlockValid {
		Blockchain = append(Blockchain, block)
	}
	return isBlockValid
}

func CreateBlock(id uint, hashPrevBlock string, txArr []transaction.Transaction, rewardAdr string) Block {
	var newBlock Block
	newBlock.Id = id
	newBlock.HashPrevBlock = hashPrevBlock
	newBlock.Time = time.Now()
	newBlock.Transactions = make([]transaction.Transaction, len(txArr)+1)
	coinbaseTx, _ := transaction.CreateTransaction("", []string{rewardAdr},
		[]uint64{GetCoinsForEmition(len(Blockchain))}, 0)
	txArr = append([]transaction.Transaction{coinbaseTx}, txArr...)
	copy(newBlock.Transactions, txArr)
	return newBlock
}

func ValidateBlock(block Block) bool {
	lastBlockHash := fmt.Sprint(block.Id) + fmt.Sprint(block.Blocksize) +
		fmt.Sprint(block.Version) + block.HashPrevBlock + block.Time.String()

	for i := 0; i < len(block.Transactions); i++ {
		lastBlockHash += transaction.GetCatTxFields(block.Transactions[i])

		if !transaction.VerifyTransaction(block.Transactions[i]) {
			return false
		}
	}
	lastBlockHash = hashing.SHA1(lastBlockHash)

	if block.HashPrevBlock != lastBlockHash {
		return false
	}

	return true
}

func InitBlockchain() {
	genesisBlock := CreateBlock(0, "0000000000000000000000000000000000000000",
		nil, "0284cbd0bcf8a34035b71c5a72e37924cb960aaa0b69df4c41d50628734b8e1408")
	Blockchain = append(Blockchain, genesisBlock)
}
