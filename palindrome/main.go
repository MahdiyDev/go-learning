package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {
	words := "reverse the given word - madam and it equals to given word, then input is palindrome"
	tokens := strings.Split(words, " ")

	for _, token := range tokens {
		if token == input {
			return true
		}
	}

	return false
}

func main() {
	input := "madam"
	result := isPalindrome(input)
	fmt.Println(result)
}
