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
	"os"
	"strconv"
)

var walletName string = "wallet.json"
var RightBoundAccNum int // Accout index of the right bound
var Wallet []Account

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
