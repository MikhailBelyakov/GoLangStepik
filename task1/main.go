package main

import (
	"fmt"
	"strconv"
)

/**
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.

Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181

Sample Input:

9119
Sample Output:

811181

*/

func main() {
	var input, result string

	fmt.Scan(&input)

	inputSlice := []rune(input)

	for _, val := range inputSlice {
		intVal, _ := strconv.Atoi(string(val))

		result += strconv.Itoa(intVal * intVal)
	}
	fmt.Print(result)
}
