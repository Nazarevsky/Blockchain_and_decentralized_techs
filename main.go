package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var opSys string
var execution bool

var keyCount = make(map[int]string)

func ClearConsole() {
	if opSys == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if opSys == "linux" || opSys == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func fillKeyCount() {
	for num := 8; num <= 4096; num *= 2 {
		keyCount[num] = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(num)), nil).String()
	}
}

func task1() {
	ClearConsole()
	var task1Key string

	for true {
		if _, exist := keyCount[8]; exist {
			for num := 8; num <= 4096; num *= 2 {
				fmt.Println(strconv.FormatInt(int64(num), 10) + "-біт: " + keyCount[num])
			}
		} else {
			fillKeyCount()
			for num := 8; num <= 4096; num *= 2 {
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

func task2() {
	var task2Key string
	rand.Intn(100)

	if _, exist := keyCount[8]; !exist {
		fillKeyCount()
	}

	for true {
		ClearConsole()
		for num := 8; num <= 4096; num *= 2 {
			val, err := big.NewInt(0).SetString(keyCount[num], 10)
			if !err {
				panic(err)
			}

			val.Rand(rand.New(rand.NewSource(int64(time.Now().UnixMilli()))), val)
			fmt.Println(strconv.FormatInt(int64(num), 10) + "-біт значення: " + val.String())
		}

		fmt.Println("Будь-яка кнопка. Згенерувати нові значення.")
		fmt.Println("0. Повернутися у меню.")
		fmt.Scanln(&task2Key)
		if task2Key == "0" {
			break
		}
	}
}

func keyHandler(key string) {
	if key == "1" {
		task1()
	} else if key == "2" {
		task2()
	} else if key == "3" {

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
		fmt.Println("2. Завдання 2: Випадкові значення ключів.")
		fmt.Println("3. Завдання 3: Брутфорс.")
		fmt.Println("0. Вийти з програми.")

		fmt.Scanln(&key)
		keyHandler(key)
	}
}
