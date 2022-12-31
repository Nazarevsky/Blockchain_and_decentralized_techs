package transaction

import (
	"bavovnacoin/account"
	"bavovnacoin/cryption"
	"bavovnacoin/ecdsa"
	"bavovnacoin/hashing"
	"bavovnacoin/utxo"
	"fmt"
)

type Input struct {
	HashAdr   string // Hash to hide address in DB
	ScriptSig string
	OutInd    int
}

type Output struct {
	HashAdr string // Hash value of output address. Receiver makes hash of own address and looks for his belonging coins in UTXO.
	Sum     uint64
}

type Transaction struct {
	Version  uint
	Locktime uint
	Inputs   []Input
	Outputs  []Output
}

type UtxoForInput struct {
	Address string
	Index   int
}

// Generating message for signing (SCRIPTHASH_ALL)
func GetCatTxFields(tx Transaction) string {
	message := ""
	message += fmt.Sprint(tx.Version)
	message += fmt.Sprint(tx.Locktime)
	for i := 0; i < len(tx.Inputs); i++ {
		message += tx.Inputs[i].HashAdr
		message += fmt.Sprint(tx.Inputs[i].OutInd)
	}
	for i := 0; i < len(tx.Outputs); i++ {
		message += tx.Outputs[i].HashAdr
		message += fmt.Sprint(tx.Outputs[i].Sum)
	}
	return message
}

func genTxScriptSignatures(keyPair []ecdsa.KeyPair, passKey string, tx Transaction) Transaction {
	message := hashing.SHA1(GetCatTxFields(tx))
	// Signing message
	for i := 0; i < len(keyPair); i++ {
		tx.Inputs[i].ScriptSig = keyPair[i].PublKey + ecdsa.Sign(message, cryption.AES_decrypt(keyPair[i].PrivKey, passKey))
	}

	return tx
}

func ComputeTxSize(tx Transaction) int {
	size := 0
	size += 8 // 4 bytes for Version, 4 for locktime
	for i := 0; i < len(tx.Inputs); i++ {
		size += len(tx.Inputs[i].ScriptSig)
		size += 4 // Input out index size
		size += len(tx.Inputs[i].HashAdr)
		size += len(tx.Inputs[i].ScriptSig)
	}

	for i := 0; i < len(tx.Outputs); i++ {
		size += 8 // Output value
		size += len(tx.Outputs[i].HashAdr)
	}
	return size
}

func getNextInpIndex(addressInp string, utxoInputs []utxo.UTXO, utxoInd int) int {
	ind := -1
	for i := 0; i <= utxoInd; i++ {
		if utxoInputs[i].Address == addressInp {
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
func GetTransInputs(sum uint64, accUtxo []utxo.UTXO) ([]UtxoForInput, []utxo.UTXO, uint64) {
	if accUtxo == nil {
		accUtxo = account.GetAccUtxo()
	}

	accUtxo = append(accUtxo, utxo.UTXO{Address: "", Sum: 0}) // Stub value for searching
	var utxoInput []UtxoForInput
	tempSum := uint64(0)

	if len(accUtxo) == 1 && accUtxo[0].Sum >= sum {
		return append(utxoInput, UtxoForInput{accUtxo[0].Address, getNextInpIndex(accUtxo[0].Address, accUtxo, 0)}),
			accUtxo, accUtxo[0].Sum
	}

	for i := 1; i < len(accUtxo); i++ {
		if accUtxo[i-1].Sum >= sum-tempSum {
			if sum-tempSum > accUtxo[i].Sum {
				utxoInput = append(utxoInput, UtxoForInput{Address: accUtxo[i-1].Address, Index: getNextInpIndex(accUtxo[i-1].Address, accUtxo, i-1)})
				return utxoInput, accUtxo, accUtxo[i-1].Sum + tempSum
			} else {
				continue
			}
		}
		utxoInput = append(utxoInput, UtxoForInput{accUtxo[i-1].Address, getNextInpIndex(accUtxo[i-1].Address, accUtxo, i-1)})
		tempSum += accUtxo[i-1].Sum
	}
	return nil, accUtxo, tempSum
}

// Creates transaction
func CreateTransaction(passKey string, outAdr []string, outSumVals []uint64, feePerByte int, locktime uint) (Transaction, string) { // return Transaction
	var tx Transaction
	tx.Locktime = locktime
	txSize := 0
	tx.Version = 0
	genSum := uint64(0)
	for i := 0; i < len(outSumVals); i++ {
		genSum += outSumVals[i]
	}

	// Genereting outputs
	var output []Output
	for i := 0; i < len(outAdr); i++ {
		var outVal Output
		outVal.HashAdr = hashing.SHA1(outAdr[i])
		outVal.Sum = uint64(outSumVals[i])
		output = append(output, outVal)
	}

	// Genereting inputs (including tx fee)
	var input []Input
	kpAcc := make([]ecdsa.KeyPair, len(account.CurrAccount.KeyPairList))
	copy(kpAcc, account.CurrAccount.KeyPairList)
	outTxSum := uint64(0)
	needSum := genSum + uint64(txSize)*uint64(feePerByte)

	var kpForSign []ecdsa.KeyPair
	for outTxSum < needSum { // Looking for tx fee
		kpForSign = []ecdsa.KeyPair{}
		inputs, _, outSum := GetTransInputs(needSum, nil)

		if needSum > outSum {
			return tx, "Not enough coins for payment. You need: " + fmt.Sprint(needSum) + ", you have: " + fmt.Sprint(account.GetBalance())
		}

		outTxSum = outSum
		for i := 0; i < len(inputs); i++ {
			var inpVal Input
			inpVal.HashAdr = hashing.SHA1(inputs[i].Address)
			inpVal.OutInd = inputs[i].Index

			// Get private and public key for scriptSig generation
			for j := 0; j < len(kpAcc); j++ {
				if hashing.SHA1(kpAcc[j].PublKey) == inputs[i].Address {
					kpForSign = append(kpForSign, ecdsa.KeyPair{PrivKey: kpAcc[j].PrivKey, PublKey: kpAcc[j].PublKey})
				}
			}
			input = append(input, inpVal)
		}
		tx.Inputs = input
		tx.Outputs = output
		txSize = ComputeTxSize(tx)
		needSum = genSum + uint64(txSize)*uint64(feePerByte)
	}

	//Generating change output
	if outTxSum > genSum+uint64(txSize)*uint64(feePerByte) {
		account.AddKeyPairToAccount(passKey) // generate new keypair for the change
		kpLen := len(account.CurrAccount.KeyPairList)
		tx.Outputs = append(tx.Outputs, Output{HashAdr: hashing.SHA1(account.CurrAccount.KeyPairList[kpLen-1].PublKey),
			Sum: uint64(outTxSum - (genSum + uint64(txSize)*uint64(feePerByte)))})
	}
	tx = genTxScriptSignatures(kpForSign, passKey, tx)
	return tx, ""
}

func GetInputSum(inp []Input) uint64 {
	var sum uint64 = 0
	for i := 0; i < len(inp); i++ {
		sum += account.GetBalByKeyHash(inp[i].HashAdr, inp[i].OutInd)
	}
	return sum
}

func GetOutputSum(out []Output) uint64 {
	var sum uint64 = 0
	for i := 0; i < len(out); i++ {
		sum += out[i].Sum
	}
	return sum
}

func GetTxFee(tx Transaction) uint64 {
	return GetInputSum(tx.Inputs) - GetOutputSum(tx.Outputs)
}

/*
Just to show that everything works fine.

Some information is not stored in the transaction structure,
but received in this function.
*/
func PrintTransaction(tx Transaction) {
	println("Transaction")
	fmt.Printf("Version: %d\n Locktime %d\n", tx.Version, tx.Locktime)
	println("Inputs:")
	var inpSum uint64
	for i := 0; i < len(tx.Inputs); i++ {
		curVal := account.GetBalByKeyHash(tx.Inputs[i].HashAdr, tx.Inputs[i].OutInd)
		inpSum += curVal
		fmt.Printf("%d. HashAddress: %s (Bal: %d)\nOut index: %d\nScriptSig: %s\n", i, tx.Inputs[i].HashAdr, curVal,
			tx.Inputs[i].OutInd, tx.Inputs[i].ScriptSig)
	}
	println("\nOutputs:")
	var outSum uint64
	for i := 0; i < len(tx.Outputs); i++ {
		outSum += tx.Outputs[i].Sum
		fmt.Printf("%d. HashAddress: %s. Sum: %d\n", i, tx.Outputs[i].HashAdr, tx.Outputs[i].Sum)
	}
	print("(Fee: ")
	println(inpSum-outSum, ")")
}

// Verifies transaction
func VerifyTransaction(tx Transaction) bool {
	if tx.Version == 0 {
		var inpSum uint64
		hashMesOfTx := hashing.SHA1(GetCatTxFields(tx))

		// Checking signatures and unique inputs
		for i := 0; i < len(tx.Inputs); i++ {
			pubKey := tx.Inputs[i].ScriptSig[:66]
			sign := tx.Inputs[i].ScriptSig[66:]
			if !ecdsa.Verify(pubKey, sign, hashMesOfTx) {
				return false
			}
			curVal := account.GetBalByKeyHash(tx.Inputs[i].HashAdr, tx.Inputs[i].OutInd)
			inpSum += curVal

			for j := i + 1; j < len(tx.Inputs); j++ {
				if tx.Inputs[j].ScriptSig == tx.Inputs[i].ScriptSig {
					return false
				}
			}
		}

		var outSum uint64
		for i := 0; i < len(tx.Outputs); i++ {
			inpSum += tx.Outputs[i].Sum
		}

		// Checking presence of coins to be spent
		if inpSum < outSum {
			return false
		}
	}
	return true
}
