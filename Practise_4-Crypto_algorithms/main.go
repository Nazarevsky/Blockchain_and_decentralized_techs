package main

import (
	"bufio"
	"crypt/cryption"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
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

func task1() {
	for true {
		ClearConsole()
		fmt.Println("Введіть значення для шифрування (Vigenere, тільки латинська абетка)")
		var mes string
		scann.Scan()
		mes = scann.Text()

		fmt.Println("Введіть ключ")
		var key string
		scann.Scan()
		key = scann.Text()

		cr := cryption.VigenereEncode(mes, key)
		dcr := cryption.VigenereDecode(cr, key)
		fmt.Println("Закодоване значення:", cr)
		fmt.Println("Декодоване значення:", dcr)

		fmt.Println("\nБудь-яка кнопка. Спробувати ще раз.")
		fmt.Println("0. Повернутися у головне меню.")

		var key2 string
		fmt.Scanln(&key2)

		if key2 == "0" {
			break
		}
	}
}

func task2() {
	var _error bool = false
	for true {
		ClearConsole()
		if _error {
			fmt.Println("Ви ввели неправильне значення (тільки ціле беззнакове число)")
			_error = false
		}
		fmt.Println("Введіть значення (ціле беззнакове число) для шифрування (RSA)")
		var mes string
		scann.Scan()
		mes = scann.Text()

		val, err := strconv.ParseUint(mes, 10, 64)
		if err != nil {
			_error = true
			continue
		}

		n, E, D := cryption.RSA_keygen()
		fmt.Println("\nn =", strconv.FormatUint(n, 10))
		fmt.Println("Згенеровані ключі: E: " + strconv.FormatUint(E, 10) + ", D: " + strconv.FormatUint(D, 10))

		cr := cryption.RSA(val, E, n)
		fmt.Println("Зашифроване значення:", cr)
		dcr := cryption.RSA(cr, D, n)
		fmt.Println("Розшифроване значення:", dcr)

		fmt.Println("\nБудь-яка кнопка. Спробувати ще раз.")
		fmt.Println("0. Повернутися у головне меню.")

		var key2 string
		fmt.Scanln(&key2)

		if key2 == "0" {
			break
		}
	}
}

func task3() {
	for true {
		ClearConsole()
		fmt.Println("Введіть значення для шифрування AES")
		var mes string
		scann.Scan()
		mes = scann.Text()

		fmt.Println("Введіть ключ (довжина - 16 байтів)")
		var key string
		err := false
		for true {
			if err {
				fmt.Println("Помилка: Довжина ключу повинна дорівнювати 16. Введіть ключ ще раз")
				err = false
			}

			scann.Scan()
			key = scann.Text()

			if len(key) != 16 {
				err = true
				continue
			}
			break
		}

		cr := cryption.AES_encrypt(mes, key)
		dcr := cryption.AES_decrypt(cr, key)
		fmt.Println("Закодоване значення:", cr)
		fmt.Println("Декодоване значення:", dcr)

		fmt.Println("\nБудь-яка кнопка. Спробувати ще раз.")
		fmt.Println("0. Повернутися у головне меню.")

		var key2 string
		fmt.Scanln(&key2)

		if key2 == "0" {
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
		task3()
	} else if key == "0" {
		execution = false
		fmt.Println("До зустрічі!")
	}
}

func main() {
	opSys = runtime.GOOS
	scann = bufio.NewScanner(os.Stdin)
	execution = true
	var key string

	for execution {
		ClearConsole()
		fmt.Println("Оберіть завдання та нажміть потрібну цифру.")
		fmt.Println("1. Завдання 1: Vigenere")
		fmt.Println("2. Завдання 2: RSA (тільки числа)")
		fmt.Println("3. Завдання 3: AES")
		fmt.Println("0. Вийти з програми.")

		fmt.Scanln(&key)
		keyHandler(key)
	}
}
