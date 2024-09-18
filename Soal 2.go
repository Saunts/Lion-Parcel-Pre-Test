package main

import (
	"fmt"
)

func main() {
	arrNums := []int{2, 3, 6, 1, 89, 1} //Inputnya
	maxNum := 0

	for _, num := range arrNums {
		if num > maxNum {
			maxNum = num
		}
	}

	fmt.Println(maxNum)
}
