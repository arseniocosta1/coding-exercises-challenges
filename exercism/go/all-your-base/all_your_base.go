package allyourbase

import (
	"fmt"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	}

	if outputBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}

	if len(inputDigits) == 0 {
		return []int{0}, nil
	}

	var num int
	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		// Horner's method
		num = num*inputBase + d
	}

	if num == 0 {
		return []int{0}, nil
	}

	res := make([]int, 0)
	for num > 0 {
		rem := num % outputBase
		res = append([]int{rem}, res...)
		num /= outputBase
	}

	return res, nil
}
