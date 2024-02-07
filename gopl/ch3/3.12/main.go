package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Printf("KB: %e\n", float64(KB))
	fmt.Printf("MB: %e\n", float64(MB))
	fmt.Printf("GB: %e\n", float64(GB))
	fmt.Printf("TB: %e\n", float64(TB))
	fmt.Printf("PB: %e\n", float64(PB))
	fmt.Printf("EB: %e\n", float64(EB))
	fmt.Printf("ZB: %e\n", float64(ZB))
	fmt.Printf("YB: %e\n", float64(YB))
}
