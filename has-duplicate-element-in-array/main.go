package main

import "fmt"

func hasDuplicateElementInArray(input []int) (bool, []int) {
	lenNums := len(input)

	result := make([]int, 0, lenNums)

	for i := 0; i < lenNums; i++ {
		num := input[i]

		for j := 0; j < lenNums; j++ {
			if i != j {
				if num == input[j] {
					result = append(result, num)
				}
			}
		}
	}

	return len(result) == 0, result
}

func main() {
	input := []int{1, 2, 2, 21, 4, 1}

	res, odd := hasDuplicateElementInArray(input)

	fmt.Printf("%v is unique: %v, odd numbers: %v\n", input, res, odd)
}
