package standard

import (
	"log"

	"golang.org/x/net/html"
)

func FindNodeWithId(node *html.Node, id string) *html.Node {
	for _, attr := range node.Attr {
		if attr.Key == "id" && attr.Val == id {
			return node
		}
	}

	for child := range node.ChildNodes() {
		if n := FindNodeWithId(child, id); n != nil {
			return n
		}
	}

	if node.NextSibling != nil {
		return FindNodeWithId(node.NextSibling, id)
	}

	return nil
}

func findTag(node *html.Node, tag string) (*html.Node, bool) {
	if node == nil {
		return nil, false
	}

	if node.Type == html.ElementNode || node.Type == html.DocumentNode {
		if node.Data == tag {
			return node, true
		}

		n, ok := findTag(node.FirstChild, tag)
		if ok {
			return n, true
		}
	}

	return checkNextSibling(node, tag)
}

func checkNextSibling(node *html.Node, s string) (*html.Node, bool) {
	for current := node.NextSibling; current != nil; current = current.NextSibling {
		n, ok := findTag(current, s)
		if ok {
			return n, true
		}
	}

	return nil, false
}

func Path(node *html.Node) string {
	if node.Parent == nil || node.Parent.Data == "" {
		return node.Data
	}

	return Path(node.Parent) + "." + node.Data
}

func FindTBody(table *html.Node) *html.Node {
	for current := table.FirstChild; current != nil; current = current.NextSibling {
		if current.Data == "tbody" {
			return current
		}
	}

	log.Fatal("Error finding tbody")

	return nil
}

func FindTableWithCaption(page *html.Node, caption string) (*html.Node, bool) {
	for current := page; current != nil; current = current.NextSibling {
		if current.Type != html.ElementNode && current.Type != html.DocumentNode {
			continue
		}

		if current.Data == "table" && hasCaption(current, caption) {
			return current, true
		}

		table, ok := FindTableWithCaption(current.FirstChild, caption)
		if ok {
			return table, true
		}
	}

	return nil, false
}

func hasCaption(parent *html.Node, caption string) bool {
	current := parent.FirstChild

	for current != nil {
		if current.Type != html.ElementNode {
			current = current.NextSibling
			continue
		}

		if current.Data != "caption" {
			current = current.NextSibling
			continue
		}

		child := current.FirstChild
		if child.Type != html.TextNode {
			return false
		}

		return child.Data == caption
	}

	return false
}
