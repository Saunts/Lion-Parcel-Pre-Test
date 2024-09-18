package main

import (
	"fmt"
)

func main() {
	number := 5 //inputnya

	result := getFactorial(number)

	fmt.Println(result)
}

func getFactorial(num int) int {
	if num == 1 {
		return 1
	}

	return num * getFactorial(num-1)
}
