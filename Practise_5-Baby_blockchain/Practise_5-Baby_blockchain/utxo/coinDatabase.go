package utxo

import (
	"fmt"
)

type UTXO struct {
	Id     int
	PubKey string
	Sum    uint64
}

var UtxoList []UTXO

// Future realization: get all UTXOs from DB
func InitUTXOValues() {
	UtxoList = append(UtxoList, UTXO{Id: 0, PubKey: "0284cbd0bcf8a34035b71c5a72e37924cb960aaa0b69df4c41d50628734b8e1408", Sum: 50000000})
	UtxoList = append(UtxoList, UTXO{Id: 1, PubKey: "03265af825325d3c07156142a5906235547855eda0499ca2e4cad729eebd67898d", Sum: 10040020})
	UtxoList = append(UtxoList, UTXO{Id: 2, PubKey: "02db1c6c791978448ef671746fab54797495597730b9a68721fca780db002162f0", Sum: 300200})
	UtxoList = append(UtxoList, UTXO{Id: 3, PubKey: "02487318bd34dc641708fd3300e2f5cca03cebe9ce7bea8973a40fc5b383de951d", Sum: 5160897})
	UtxoList = append(UtxoList, UTXO{Id: 4, PubKey: "02db1c6c791978448ef671746fab54797495597730b9a68721fca780db002162f0", Sum: 230000000})
	UtxoList = append(UtxoList, UTXO{Id: 5, PubKey: "03c66667101e5aecdcd5a0b86d21445b98d6862165b0cd4599a58a2f87f9d14cc5", Sum: 150000})
	UtxoList = append(UtxoList, UTXO{Id: 6, PubKey: "03265af825325d3c07156142a5906235547855eda0499ca2e4cad729eebd67898d", Sum: 781178})
	UtxoList = append(UtxoList, UTXO{Id: 7, PubKey: "02487318bd34dc641708fd3300e2f5cca03cebe9ce7bea8973a40fc5b383de951d", Sum: 871054})
	UtxoList = append(UtxoList, UTXO{Id: 8, PubKey: "0284cbd0bcf8a34035b71c5a72e37924cb960aaa0b69df4c41d50628734b8e1408", Sum: 9712054})
	UtxoList = append(UtxoList, UTXO{Id: 9, PubKey: "02db1c6c791978448ef671746fab54797495597730b9a68721fca780db002162f0", Sum: 7912})
	UtxoList = append(UtxoList, UTXO{Id: 10, PubKey: "03265af825325d3c07156142a5906235547855eda0499ca2e4cad729eebd67898d", Sum: 79874132})
	UtxoList = append(UtxoList, UTXO{Id: 11, PubKey: "03c66667101e5aecdcd5a0b86d21445b98d6862165b0cd4599a58a2f87f9d14cc5", Sum: 7912})
}

func DelFromUtxo(id int) {
	for i := 0; i < len(UtxoList); i++ {
		if UtxoList[i].Id == id {
			UtxoList = append(UtxoList[:i], UtxoList[i+1:]...)
			return
		}
	}
}

func PrintUtxo() {
	for i := 0; i < len(UtxoList); i++ {
		fmt.Printf("Pub: %s, sum: %d\n", UtxoList[i].PubKey, UtxoList[i].Sum)
	}
}
