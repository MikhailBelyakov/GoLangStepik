package main

import (
	"fmt"
	"unicode"
)

const wrongText = "Wrong password"
const successText = "Ok"

/**
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:
fdsghdfgjsdDD1
Sample Output:
Ok
*/
func main() {
	var input string

	fmt.Scan(&input)
	runePasswordArr := []rune(input)
	existWrongSymbol := false

	for _, runePassword := range runePasswordArr {
		if !unicode.Is(unicode.Latin, runePassword) && !unicode.IsNumber(runePassword) {
			existWrongSymbol = true
		}
	}

	if len(input) < 5 || existWrongSymbol {
		fmt.Println(wrongText)
	} else {
		fmt.Println(successText)
	}
}
