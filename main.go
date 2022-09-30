package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var opSys string
var execution bool

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
		fmt.Println("0. Вийти з програми")

		fmt.Scanln(&key)
		keyHandler(key)
	}
}
