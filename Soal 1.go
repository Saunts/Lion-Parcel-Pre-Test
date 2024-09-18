package main

import (
	"fmt"
)

func main() {
	word := "roger" //inputnya
	palindrome := true

	for i := 0; i < (len(word) / 2); i++ {
		if word[i] != word[len(word)-i-1] {
			palindrome = false
			break
		}
	}

	if palindrome {
		fmt.Println("True")
		return
	}

	fmt.Println("False")
}
