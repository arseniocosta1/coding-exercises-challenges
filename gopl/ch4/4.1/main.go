package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	args := os.Args[1:]

	var hash1, hash2 [32]byte
	if len(args) >= 2 {
		hash1 = sha256.Sum256([]byte(args[0]))
		hash2 = sha256.Sum256([]byte(args[1]))
	} else {
		hash1 = sha256.Sum256([]byte("x"))
		hash2 = sha256.Sum256([]byte("X"))
	}

	fmt.Printf("%#x\n%#x\n\n%08[1]b\n%08[2]b\n", hash1, hash2)

	fmt.Printf("%d different bits\n", countDiffBits(&hash1, &hash2))

	fmt.Printf("%d different bits\n", countDiffBitsWithPrecomputedTable(&hash1, &hash2))
}

func countDiffBits(b1, b2 *[sha256.Size]byte) int {
	// count the number of bits that are different between the two hashes
	// using XOR technique and counting the number of 1s
	count := 0
	for i := 0; i < sha256.Size; i++ {
		diff := b1[i] ^ b2[i]
		for diff != 0 {
			count++
			diff &= diff - 1
		}
	}
	return count
}

func countDiffBitsWithPrecomputedTable(b1, b2 *[sha256.Size]byte) int {
	// count the number of bits that are different between the two hashes
	// using XOR technique and using a precomputed PopCount table
	var sum int
	for i := 0; i < sha256.Size; i++ {
		sum += int(pc[b1[i]^b2[i]])
	}
	return sum
}
