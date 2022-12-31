package utxo

import "fmt"

type UTXO struct {
	Address string
	Sum     uint64
}

var UtxoList []UTXO

// Future realization: get all UTXOs from DB
func InitUTXOValues() {
	UtxoList = append(UtxoList, UTXO{Address: "ce390e1683d0fa81568d2042f6f84ac0b0bf0bb9", Sum: 50000000})
	UtxoList = append(UtxoList, UTXO{Address: "41b8ec0dbbe0c29f15982d4d6155d4424adabcf6", Sum: 10040020})
	UtxoList = append(UtxoList, UTXO{Address: "5b0d98f7e17107def8c9723aedfe34fcd8aef80e", Sum: 300200})
	UtxoList = append(UtxoList, UTXO{Address: "52ccca777a1ab46ddd07257b7def18f3c33e1fec", Sum: 5160897})
	UtxoList = append(UtxoList, UTXO{Address: "5b0d98f7e17107def8c9723aedfe34fcd8aef80e", Sum: 230000000})
	UtxoList = append(UtxoList, UTXO{Address: "0accd57e23dfd44edc1e8a60f8a0889aa4b277a8", Sum: 150000})
	UtxoList = append(UtxoList, UTXO{Address: "41b8ec0dbbe0c29f15982d4d6155d4424adabcf6", Sum: 781178})
	UtxoList = append(UtxoList, UTXO{Address: "52ccca777a1ab46ddd07257b7def18f3c33e1fec", Sum: 871054})
	UtxoList = append(UtxoList, UTXO{Address: "ce390e1683d0fa81568d2042f6f84ac0b0bf0bb9", Sum: 9712054})
	UtxoList = append(UtxoList, UTXO{Address: "5b0d98f7e17107def8c9723aedfe34fcd8aef80e", Sum: 7912})
	UtxoList = append(UtxoList, UTXO{Address: "41b8ec0dbbe0c29f15982d4d6155d4424adabcf6", Sum: 79874132})
	UtxoList = append(UtxoList, UTXO{Address: "0accd57e23dfd44edc1e8a60f8a0889aa4b277a8", Sum: 7912})
}

func DelFromUtxo(address string, outind int) {
	ind := 0
	for i := 0; i < len(UtxoList); i++ {
		if UtxoList[i].Address == address {
			if ind == outind {
				UtxoList = append(UtxoList[:i], UtxoList[i+1:]...)
				return
			}
			ind++
		}
	}
}

func AddToUtxo(address string, sum uint64) {
	UtxoList = append(UtxoList, UTXO{Address: address, Sum: sum})
}

func ShowCoinDatabase() {
	for i := 0; i < len(UtxoList); i++ {
		fmt.Printf("Address: %s, sum: %d\n", UtxoList[i].Address, UtxoList[i].Sum)
	}
}
