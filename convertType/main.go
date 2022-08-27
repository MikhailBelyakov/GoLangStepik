package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Sample Input:
1 149,6088607594936;1 179,0666666666666

Sample Output:
0.9750
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, ",", ".", -1)

	splitTextArr := strings.Split(text, ";")

	oneParseNumber, _ := strconv.ParseFloat(splitTextArr[0], 64)
	twoParseNumber, _ := strconv.ParseFloat(splitTextArr[1], 64)

	result := oneParseNumber / twoParseNumber
		
	fmt.Printf("%.4f", result)
}
