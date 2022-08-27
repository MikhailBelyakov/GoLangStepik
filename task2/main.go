package main

import (
	"fmt"
)

/**
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу.

Sample Input:
8
Sample Output:
6

*/
func main() {
	var indexInFibonacci, fibonacci, a = 0, 0, 0
	var count, b = 1, 1

	_, err := fmt.Scan(&indexInFibonacci)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; fibonacci <= indexInFibonacci; i++ {
		fibonacci = a + b
		a = b
		b = fibonacci

		count++
		if fibonacci == indexInFibonacci {
			fmt.Print(count)
			break
		} else if fibonacci > indexInFibonacci {
			fmt.Print(-1)
			break
		}
	}
}
