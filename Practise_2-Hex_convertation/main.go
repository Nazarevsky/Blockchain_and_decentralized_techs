package main

import (
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"practise2/endian"
	"runtime"
)

var execution bool
var opSys string

func ClearConsole() {
	if opSys == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if opSys == "linux" || opSys == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Println("\n\n")
	}
}

func keyHandler(key string) {
	if key == "1" {
		taskFunction12(endian.HexToLittleEndian, "Little Endian")
	} else if key == "2" {
		taskFunction12(endian.HexToBigEndian, "Big Endian")
	} else if key == "3" {
		taskFunction34(endian.LittleEndianToHex, "Little Endian --> HEX")
	} else if key == "4" {
		taskFunction34(endian.BigEndianToHex, "Big Endian --> HEX")
	} else if key == "0" {
		execution = false
		fmt.Println("До зустрічі!")
	}
}

func isHexNum(num string) (string, bool) {
	if len(num) > 1 && (num[:2] == "0x" || num[:2] == "0X") {
		num = num[2:]
	}
	val, err := big.NewInt(0).SetString(num, 16)

	if !err || big.NewInt(0).Cmp(val) == 1 {
		return num, false
	}
	return num, true
}

func taskFunction12(fn func(string) *big.Int, _type string) {
	err := true

	for true {
		ClearConsole()
		fmt.Println(_type + ".")
		if !err {
			fmt.Println("Помилка. Введене значення - не ціле беззнакове число.")
			err = false
		}
		fmt.Println("Введіть HEX значення. Приклад: 0xff00a або ff00a. Регістр значення не має.")
		var str string
		fmt.Scanln(&str)
		str, err = isHexNum(str)

		if err {
			num := fn(str)
			fmt.Println("\nЗначення " + _type + ": " + num.String() + "\n")
			fmt.Println("Будь-яка кнопка. Спробувати ще раз.")
			fmt.Println("0. Повернутися у головне меню.")

			var key string
			fmt.Scanln(&key)

			if key == "0" {
				break
			}
		}
	}
}

func taskFunction34(fn func(*big.Int) string, _type string) {
	err := true
	val := big.NewInt(0)
	for true {
		ClearConsole()
		fmt.Println(_type + ".")
		if !err || (big.NewInt(0).Cmp(val) == 1) {
			fmt.Println("Помилка. Введене значення - не ціле беззнакове число.")
			err = false
		}
		fmt.Println("Введіть значення у десятковому вигляді.")
		var str string
		fmt.Scanln(&str)
		val, err = big.NewInt(0).SetString(str, 10)

		if err && !(big.NewInt(0).Cmp(val) == 1) {
			num := fn(val)
			fmt.Println("\nЗначення 0x" + _type + ": " + num + "\n")
			fmt.Println("Будь-яка кнопка. Спробувати ще раз.")
			fmt.Println("0. Повернутися у головне меню.")

			var key string
			fmt.Scanln(&key)

			if key == "0" {
				break
			}
		}
	}
}

func main() {
	opSys = runtime.GOOS
	execution = true
	var key string
	for execution {
		ClearConsole()
		fmt.Println("Оберіть завдання та нажміть потрібну цифру.")
		fmt.Println("1. Завдання 1: HEX --> Little Endian")
		fmt.Println("2. Завдання 2: HEX --> Big Endian")
		fmt.Println("3. Завдання 3: Little Endian --> HEX")
		fmt.Println("4. Завдання 3: Big Endian --> HEX")
		fmt.Println("0. Вийти з програми.")

		fmt.Scanln(&key)
		keyHandler(key)
	}
}
