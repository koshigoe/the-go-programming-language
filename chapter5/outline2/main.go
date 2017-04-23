package main

import (
	"golang.org/x/net/html"
	"fmt"
	"os"
)

// forEachNode は n から始まるツリー内の個々のノード x に対して
// 関数 pre(x) と post(x) を呼び出します。その二つの関数はオプションです。
// pre は子ノードを訪れる前に呼び出され(前順: preorder)、
// post は子ノードを訪れた後に呼び出されます(後順: postorder)。
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

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}
