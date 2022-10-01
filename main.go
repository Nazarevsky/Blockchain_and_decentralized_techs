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
var bits []int = []int{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}

func keyHandler(key string) {
	if key == "0" {
		task1()
	} else if key == "1" {
		task2()
	} else if key == "2" {
		task3()
	} else if key == "*" {
		execution = false
		fmt.Println("До зустрічі!")
	}
}

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

func fillKeyCount() {
	for i := 0; i < len(bits); i++ {
		keyCount[bits[i]] = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(bits[i])), nil).String()
	}
}

func task1() {
	ClearConsole()
	var task1Key string

	for true {
		if _, exist := keyCount[8]; exist {
			for i := 0; i < len(bits); i++ {
				fmt.Println(strconv.FormatInt(int64(bits[i]), 10) + " біт: " + keyCount[bits[i]])
			}
		} else {
			fillKeyCount()
			for i := 0; i < len(bits); i++ {
				fmt.Println(strconv.FormatInt(int64(bits[i]), 10) + " біт: " + keyCount[bits[i]])
			}
		}

		fmt.Println("*. Повернутися у меню.")
		fmt.Scanln(&task1Key)
		if task1Key == "*" {
			break
		}
	}
}

func task2() {
	var task2Key string

	if _, exist := keyCount[8]; !exist {
		fillKeyCount()
	}

	for true {
		ClearConsole()
		for i := 0; i < len(bits); i++ {
			val, err := big.NewInt(0).SetString(keyCount[bits[i]], 10)
			if !err {
				panic(err)
			}

			val.Rand(rand.New(rand.NewSource(int64(time.Now().UnixMilli()))), val)
			fmt.Println(strconv.FormatInt(int64(bits[i]), 10) + " біт значення: " + val.String())
		}

		fmt.Println("Будь-яка кнопка. Згенерувати нові значення.")
		fmt.Println("*. Повернутися у меню.")
		fmt.Scanln(&task2Key)
		if task2Key == "*" {
			break
		}
	}
}

func bruteForce(bitCount int) {
	var bruteForceKey string
	if _, exist := keyCount[8]; !exist {
		fillKeyCount()
	}

	randNumMax, err := big.NewInt(0).SetString(keyCount[bitCount], 10)
	randNum := big.NewInt(randNumMax.Int64())
	if !err {
		panic(err)
	}
	for true {
		ClearConsole()
		randNum.Rand(rand.New(rand.NewSource(int64(time.Now().UnixMilli()))), randNumMax)
		fmt.Println("Шукається значення (" + randNum.String() + ")...")

		start := time.Now().UnixMilli()
		i := big.NewInt(0)
		for ; i.Cmp(randNum) < 0; i.Add(i, big.NewInt(1)) {
		}
		end := time.Now().UnixMilli()

		fmt.Println("Значення знайдено: " + i.String() + " знайдено за " + strconv.FormatInt(end-start, 10) + " мс.")
		fmt.Println("Будь-яка кнопка. Згенерувати нове значення.")
		fmt.Println("*. Повернутися у меню.")
		fmt.Scanln(&bruteForceKey)
		if bruteForceKey == "*" {
			break
		}
	}
}

func task3() {
	ClearConsole()
	var task3Key string

	for true {
		ClearConsole()
		fmt.Println("Оберіть кількість біт для генерації ключа та натисніть на відповідну кнопку")

		for i := 0; i < len(bits); i++ {
			fmt.Println(strconv.FormatInt(int64(i), 10) + ". " + strconv.FormatInt(int64(bits[i]), 10) + " біт")
		}

		fmt.Println("*. Повернутися у меню.")
		fmt.Scanln(&task3Key)
		if task3Key == "*" {
			break
		} else if task3Key[0] >= 48 && task3Key[0] <= 57 {
			bruteForce(bits[task3Key[0]%48])
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
		fmt.Println("0. Завдання 1: Кількість варіантів ключів.")
		fmt.Println("1. Завдання 2: Випадкові значення ключів.")
		fmt.Println("2. Завдання 3: Брутфорс.")
		fmt.Println("*. Вийти з програми.")

		fmt.Scanln(&key)
		keyHandler(key)
	}
}
