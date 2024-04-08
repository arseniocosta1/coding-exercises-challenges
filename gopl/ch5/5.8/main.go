package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var elem *html.Node

	findElem := func(n *html.Node) bool {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				elem = n
				return false
			}
		}

		return true
	}

	forEachNode(doc, findElem, nil)

	return elem
}

func main() {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		panic(err)
	}

	id := "myid"
	args := os.Args[1:]
	if len(args) > 0 {
		id = args[0]
	}

	elem := ElementByID(doc, id)

	if elem == nil {
		fmt.Printf("Element with id %s not found", id)
		return
	}

	fmt.Printf("Element with id %s found: %+v", id, *elem)
}
