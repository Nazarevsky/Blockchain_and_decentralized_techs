package account

/*
	The wallet data (list of Account structure - wallet.json) is stored
	locally on user's device. There could be a multiple ammount of accounts
	in one wallet (many users can use one device, so there should be a way to
	distinguish data). Private key is encrypted by user's own password
	using AES algorithm. The password is stored in wallet.json as a hash
	value that is created using SHA-1.
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var walletName string = "wallet.json"
var RightBoundAccNum int // Accout index of the right bound
var Wallet []Account

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

func PrintUtxo() {
	for i := 0; i < len(UtxoList); i++ {
		fmt.Printf("Pub: %s, sum: %d\n", UtxoList[i].PubKey, UtxoList[i].Sum)
	}
}

func isWalletExists(name string) bool {
	file, err := os.Open(name)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	file.Close()
	return true
}

func InitAccountsData() {
	var emptyDat []byte
	err := isWalletExists(walletName)
	if !err {
		os.WriteFile(walletName, emptyDat, 0777)
	} else {
		dataByte, _ := os.ReadFile(walletName)
		json.Unmarshal([]byte(dataByte), &Wallet)
		RightBoundAccNum, _ = strconv.Atoi(Wallet[len(Wallet)-1].Id)
	}
}

func WriteAccounts() {
	byteData, _ := json.MarshalIndent(Wallet, "", "    ")
	os.WriteFile(walletName, byteData, 0777)
}

func DelFromUtxo(id int) {
	for i := 0; i < len(UtxoList); i++ {
		if UtxoList[i].Id == id {
			UtxoList = append(UtxoList[:i], UtxoList[i+1:]...)
			return
		}
	}
}
