package main

import "fmt"

/**
На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.

Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.

Sample Input:

2
3
56
45
21
Sample Output:

56
*/

func main() {
	var input, maxNumber int
	sortArray := [5]int{}

	for i := 0; i < 5; i++ {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Ошибка чтения данных")
			return
		}
		sortArray[i] = input
	}

	for key, val := range sortArray {
		if key == 0 {
			maxNumber = val
		} else if maxNumber < val {
			maxNumber = val
		}
	}

	fmt.Println(maxNumber)
}
