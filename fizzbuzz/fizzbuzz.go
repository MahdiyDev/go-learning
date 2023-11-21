package main

import "fmt"

func main() {
	var count int = 10

	for input := 0; input < count; input++ {
		if input%3 == 0 && input%5 == 0 {
			fmt.Println(input, "fizz buzz")
		} else if input%3 == 0 {
			fmt.Println(input, "fizz")
		} else if input%5 == 0 {
			fmt.Println(input, "buzz")
		} else {
			fmt.Println(input)
		}
	}
}
