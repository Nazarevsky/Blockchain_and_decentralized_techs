package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"pract3/hashing"
	"runtime"
)

var execution bool
var opSys string
var scann *bufio.Scanner

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

func task1(fn func(string) string, algName string) {
	for true {
		ClearConsole()
		fmt.Println("Введіть повідомдення для гешування за алгоритмом", algName)
		var mes string
		scann.Scan()
		mes = scann.Text()

		fmt.Println("Результат:", fn(mes))
		fmt.Println("\nБудь-яка кнопка. Спробувати ще раз.")
		fmt.Println("0. Повернутися у головне меню.")

		var key string
		fmt.Scanln(&key)

		if key == "0" {
			break
		}
	}
}

func keyHandler(key string) {
	if key == "1" {
		task1(hashing.SHA1, "SHA-1")
	} else if key == "2" {
		task1(hashing.Keccak, "Keccak")
	} else if key == "0" {
		execution = false
		fmt.Println("До зустрічі!")
	}
}

func a(b []int) []int {
	b[0] = 13
	return b
}

func main() {
	opSys = runtime.GOOS
	scann = bufio.NewScanner(os.Stdin)
	execution = true
	var key string

	for execution {
		ClearConsole()
		fmt.Println("Оберіть завдання та нажміть потрібну цифру.")
		fmt.Println("1. Завдання 1: SHA-1")
		fmt.Println("1. Завдання 2: Keccak (Працює неправильно)")
		fmt.Println("0. Вийти з програми.")

		fmt.Scanln(&key)
		keyHandler(key)
	}
}
