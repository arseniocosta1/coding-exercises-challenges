package main

import (
	"fmt"
	"os"
	"slices"

	"golang.org/x/net/html"
)

// visit appends to links each link found in n and returns the
// result.
func visit(links []string, n *html.Node) []string {
	tags := []string{"a", "link", "script", "img", "iframe", "frame", "area", "base", "form"}
	if n.Type == html.ElementNode && slices.Contains(tags, n.Data) {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

// you can run this program together with the fetch one by piping the output of fetch to it, like this:
// go run gopl.io/ch1/fetch http://golang.org | go run main.go
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
