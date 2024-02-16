package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}

	fmt.Printf("Original: %v\n", array)
	rotate(array, -2)
	fmt.Printf("Rotated:  %v\n", array)
}

func rotate(s []int, n int) {
	lenS := len(s)
	if lenS == 0 {
		return // No rotation needed for an empty slice
	}

	// Normalize n to ensure it's within the range [0, lenS)
	// handles both left and right rotations by converting negative n to its positive equivalent
	n = ((n % lenS) + lenS) % lenS

	if n == 0 {
		return // No rotation needed

	}

	// Create a new slice for the rotated elements
	rotated := make([]int, lenS)

	// Calculate new positions and copy elements directly
	for i := 0; i < lenS; i++ {
		newPos := (i + n) % lenS
		rotated[newPos] = s[i]
	}

	// Copy the rotated elements back to the original slice
	copy(s, rotated)
}
