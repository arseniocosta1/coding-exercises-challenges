package main

import (
	"fmt"

	"gopl-solutions/ch2/2.1/tempconv"
)

func main() {
	c := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", c)
}
