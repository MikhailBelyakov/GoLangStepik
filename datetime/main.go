package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/**
Sample Input:
13.03.2018 14:00:15,12.03.2018 14:00:15
Sample Output:
24h0m0s
*/
func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	inputExplode := strings.Split(input, ",")

	dateOne, _ := time.Parse("02.01.2006 15:04:05", inputExplode[0])
	dateTwo, _ := time.Parse("02.01.2006 15:04:05", inputExplode[1])

	var result time.Duration

	if dateTwo.Before(dateOne) {
		result = time.Since(dateTwo) - time.Since(dateOne)
	} else {
		result = time.Since(dateOne) - time.Since(dateTwo)
	}

	fmt.Println(result.Round(time.Second))
}
