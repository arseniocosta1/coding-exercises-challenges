package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("(i=%d, val=%q)\n", i, arg)
		// or something like this:
		//fmt.Println("index:", i, "value:", arg)
	}
}
