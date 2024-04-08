// counts the number of each element in an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type ProcessNodeFun func(*html.Node)

func traverseDOM(n *html.Node, pre, pos ProcessNodeFun) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseDOM(c, pre, pos)
	}

	if pos != nil {
		pos(n)
	}
}

// you can run this program together with the fetch one by piping the output of fetch to it, like this:
// go run gopl.io/ch1/fetch http://golang.org | go run main.go
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	pre := func() (ProcessNodeFun, map[string]int) {
		freq := make(map[string]int)

		adder := func(n *html.Node) {
			if n.Type == html.ElementNode {
				freq[n.Data]++
			}
		}

		return adder, freq
	}

	countFun, counts := pre()

	traverseDOM(doc, countFun, nil)

	fmt.Println("Element\t Count")
	for key, value := range counts {
		fmt.Printf("%-15v:%6d\n", key, value)
	}
}
