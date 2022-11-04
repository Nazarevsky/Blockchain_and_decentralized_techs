package cryption

// Only for latin alphabet

import (
	"strings"
)

func padKey(key string, needLen int) string {
	lenKey := len(key)
	if lenKey > needLen {
		return key[:needLen]
	}
	var intWords int = (needLen - lenKey) / lenKey
	key += strings.Repeat(key, intWords) + key[:(needLen-lenKey)%lenKey]
	return key
}

func toUpper(l byte) byte {
	return 65 + (l - 97)
}

func toLower(l byte) byte {
	return 97 + (l - 65)
}

func getEncrLetter(mesLet, keyLet byte) byte {
	if mesLet >= 65 && mesLet <= 90 { // bouds for capital letters
		return ((mesLet+keyLet)-130)%26 + 65
	}
	if mesLet >= 97 && mesLet <= 122 { // bouds for small letters
		return toLower(((toUpper(mesLet)+keyLet)-130)%26 + 65)
	}
	return mesLet
}

func absSub(a, b byte) byte {
	if a > b {
		return a - b
	}
	return b - a
}

func getDecrLetter(mesLet, keyLet byte) byte {
	if mesLet >= 65 && mesLet <= 90 { // bouds for capital letters
		return ((mesLet-keyLet)+130)%26 + 65
	}
	if mesLet >= 97 && mesLet <= 122 { // bouds for small letters
		return toLower(((toUpper(mesLet)-keyLet)+130)%26 + 65)
	}
	return mesLet
}

func VigenereEncode(mes, key string) string {
	var res string
	key = padKey(key, len(mes))

	for i := 0; i < len(mes); i++ {
		res += string(getEncrLetter(mes[i], key[i]))
	}

	return res
}

func VigenereDecode(mes, key string) string {
	var res string
	key = padKey(key, len(mes))

	for i := 0; i < len(mes); i++ {
		res += string(getDecrLetter(mes[i], key[i]))
	}

	return res
}
