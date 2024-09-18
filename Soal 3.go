package main

import (
	"fmt"
)

func main() {
	size := 7 //inputnya
	row := 1

	for i := 0; i < size; i++ {
		for j := 0; j < row; j++ {
			fmt.Print("*")
		}
		row += 1
		fmt.Println("")
	}
}
