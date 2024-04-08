package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(node *html.Node) (words, images int) {
	if node.Type == html.TextNode {
		words += wordCount(node.Data)
	} else if node.Type == html.ElementNode && node.Data == "img" {
		images++
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}

func wordCount(data string) int {
	n := 0

	reader := bufio.NewScanner(strings.NewReader(data))
	reader.Split(bufio.ScanWords)

	for reader.Scan() {
		n++
	}

	return n
}

func main() {
	url := "http://golang.org"

	args := os.Args[1:]

	if len(args) > 1 {
		url = args[0]
	}

	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countWordsAndImages: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("words: %d\nimages: %d\n", words, images)
}
