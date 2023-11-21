package main

import "fmt"

func main() {
	var i, num, even_sum, odd_sum int

	// Asking for Input
	fmt.Printf("Enter the Maximum Value: ")
	fmt.Scanf("%d", &num)

	for i = 1; i <= num; i++ {
		if i%2 == 0 {
			even_sum += i
		} else {
			odd_sum += i
		}
	}
	fmt.Printf("The Sum of all Even Numbers up to %d is %d.\n", num, even_sum)
	fmt.Printf("The Sum of all Odd Numbers up to %d is %d.\n", num, odd_sum)
}
