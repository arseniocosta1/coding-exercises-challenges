package main

import "fmt"

const arraySize = 25

func main() {
	array := [arraySize]int{20: 1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", array)
	reverse(&array)
	fmt.Printf("Reversed: %v\n", array)
}

func reverse(input *[arraySize]int) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}
