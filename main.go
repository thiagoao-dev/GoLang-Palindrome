package main

import (
	"fmt"
	"github.com/thiagoaugustus/GoLang-Palindrome/algorithm"
	"os"
)

// main function
func main() {
	// Read first user input string
	input := os.Args[1:2]
	fmt.Printf("Is '%s' a palindrome?\n", input[0])
	if algorithm.Palindrome(input) {
		fmt.Println("Yes, it is!")
	} else {
		fmt.Println("No, it isn't!")
	}
}
