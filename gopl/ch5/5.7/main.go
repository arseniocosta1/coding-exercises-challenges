package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
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

var depth int

func isVoidElement(a atom.Atom) bool {
	var voidElements = map[atom.Atom]bool{
		atom.Area:   true,
		atom.Base:   true,
		atom.Br:     true,
		atom.Col:    true,
		atom.Embed:  true,
		atom.Hr:     true,
		atom.Img:    true,
		atom.Input:  true,
		atom.Link:   true,
		atom.Meta:   true,
		atom.Param:  true,
		atom.Source: true,
		atom.Track:  true,
		atom.Wbr:    true,
	}
	return voidElements[a]
}

func isForeignElement(name string) bool {
	var foreignElements = map[string]bool{
		"svg":  true,
		"math": true,
	}

	return foreignElements[name]
}

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		attributes := ""
		for _, a := range n.Attr {
			attr := ""

			if a.Namespace != "" {
				attr = fmt.Sprintf("%s:%s=%q", a.Namespace, a.Key, a.Val)
			} else {
				attr = fmt.Sprintf("%s=%q", a.Key, a.Val)
			}

			if len(n.Attr) > 3 {
				attributes += fmt.Sprintf("\n%*s%s", (depth+1)*2, "", attr)
			} else {
				attributes += fmt.Sprintf(" %s=%q", a.Key, a.Val)
			}
		}

		fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attributes)

		depth++
	case html.TextNode:
		trimData := strings.TrimSpace(n.Data)
		if trimData != "" {
			fmt.Printf("%*s%s\n", (depth+1)*2, "", trimData)
		}
	case html.CommentNode:
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	case html.DoctypeNode:
		fmt.Printf("<!doctype html>\n")
	default:

	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if !isVoidElement(n.DataAtom) {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func main() {
	url := "http://golang.org"

	args := os.Args[1:]
	if len(args) > 1 {
		url = args[0]
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	forEachNode(doc, startElement, endElement)
}
