package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type ProcessNodeFun func(*html.Node)
type ContinueFun func(*html.Node) bool

// traverseDOM traverses the DOM tree rooted at n, calling the optional pre and post functions at each node.
// the skip function can be used to skip the traversal of a node and its children.
func traverseDOM(n *html.Node, proc ProcessNodeFun, skip ContinueFun) {
	if skip != nil && !skip(n) {
		return
	}

	if proc != nil {
		proc(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseDOM(c, proc, skip)
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

	counts := make([]string, 0)
	countFun := func(n *html.Node) {
		if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
			counts = append(counts, strings.TrimSpace(n.Data))
		}
	}
	traverseDOM(doc, countFun, func(n *html.Node) bool {
		return !(n.Type ==
			html.ElementNode && n.Data == "script" || n.Data == "style")
	})

	fmt.Println("Element\t Count")
	for _, value := range counts {
		fmt.Printf("%s\n", value)
	}
}
