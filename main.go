package main

import (
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

var opSys string
var execution bool

var keyCount = make(map[int]string)

func ClearConsole() {
	if opSys == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if opSys == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func task1() {
	ClearConsole()
	var task1Key string
	result := big.NewInt(0)
	l := big.NewInt(2)

	for true {
		if val, exist := keyCount[8]; exist {
			fmt.Println("8-біт: " + val)
			for num := 16; num <= 4096; num *= 2 {
				fmt.Println(strconv.FormatInt(int64(num), 10) + "-біт: " + keyCount[num])
			}
		} else {
			for num := 8; num <= 4096; num *= 2 {
				keyCount[num] = result.Exp(l, big.NewInt(int64(num)), nil).String()
				fmt.Println(strconv.FormatInt(int64(num), 10) + "-біт: " + keyCount[num])
			}
		}

		fmt.Println("0. Повернутися у меню.")
		fmt.Scanln(&task1Key)
		if task1Key == "0" {
			break
		}
	}
}

func keyHandler(key string) {
	if key == "1" {
		task1()
	} else if key == "0" {
		execution = false
		fmt.Println("До зустрічі!")
	}
}

func main() {
	opSys = runtime.GOOS
	execution = true
	var key string

	for execution {
		ClearConsole()
		fmt.Println("Оберіть завдання та нажміть потрібну цифру.")
		fmt.Println("1. Завдання 1: Кількість варіантів ключів.")
		fmt.Println("0. Вийти з програми.")

		fmt.Scanln(&key)
		keyHandler(key)
	}
}
