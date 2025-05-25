package standard

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"

	"golang.org/x/net/html"

	"github.com/KrischanCS/go-toolbox/set"
)

type Link struct {
	Name string
	Url  string
}

func ExtractText(data *html.Node) (string, []Link) {
	sb := strings.Builder{}
	links := make([]Link, 0, 3)

	for node := data.FirstChild; node != nil; node = node.NextSibling {
		switch node.Type {
		case html.CommentNode:
			continue
		case html.TextNode:
			if isAllWhitespace(node.Data) {
				continue
			}

			sb.WriteString(node.Data)
		case html.ElementNode, html.DoctypeNode:
			extractTextFromElement(node, &links, &sb)
		default:
			log.Panic("Unexpected node type: ", node.Type)
		}
	}

	compactWhitespace(&sb)

	s := sb.String()

	// Special case currently needed for "popover" which has a trailing ;
	s = strings.Trim(s, ";")

	s, links = DistinguishLinkDuplicates(s, links)

	return s, links
}

func DistinguishLinkDuplicates(text string, links []Link) (string, []Link) {
	links = EliminateExactSameLinks(links)

	for i, link := range links {
		duplicates := 0

		for j, otherLink := range links[i+1:] {
			currentI := i + j + 1

			if link.Name != otherLink.Name {
				continue
			}

			duplicates++

			name := fmt.Sprintf("%s (%d)", link.Name, duplicates)
			links[currentI].Name = name

			index := strings.Index(text, "["+link.Name+"]")

			text = text[:index+1] +
				strings.Replace(text[index+1:], "["+link.Name+"]", "["+name+"]", 1)
		}
	}

	return text, links
}

func EliminateExactSameLinks(links []Link) []Link {
	exactSameLink := set.Of[int]()

	for i, link := range links {
		for j, otherLink := range links[i+1:] {
			currentI := i + j + 1
			if link == otherLink {
				exactSameLink.Add(currentI)
			}
		}
	}

	newLinks := make([]Link, 0, len(links)-exactSameLink.Len())

	for i, l := range links {
		if !exactSameLink.Contains(i) {
			newLinks = append(newLinks, l)
		}
	}

	return newLinks
}

func extractTextFromElement(node *html.Node, links *[]Link, sb *strings.Builder) {
	if node.Data == "a" {
		addAsLink(node, links, sb)
		return
	}

	if node.Data == "li" {
		if node.PrevSibling != nil && node.PrevSibling.Data == "li" {
			sb.WriteByte(',')
		}

		sb.WriteByte(' ')
	}

	text, ls := ExtractText(node)

	sb.WriteString(text)

	*links = append(*links, ls...)
}

func addAsLink(node *html.Node, links *[]Link, sb *strings.Builder) {
	var url string

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			url = HTMLStandardURL + attr.Val
			break
		}
	}

	text, ls := ExtractText(node)
	name := text

	*links = append(*links, Link{Name: name, Url: url})

	sb.WriteString("[" + text + "]")

	*links = append(*links, ls...)
}

func compactWhitespace(sb *strings.Builder) {
	scanner := bufio.NewScanner(strings.NewReader(sb.String()))
	sb.Reset()

	//nolint:wsl
	for scanner.Scan() {
		bs := scanner.Bytes()
		bs = bytes.Trim(bs, " \t")

		lastWasWhitespace := false
		for _, b := range bs {
			if b != ' ' && b != '\t' {
				lastWasWhitespace = false
				sb.WriteByte(b)

				continue
			}

			if lastWasWhitespace {
				continue
			}

			lastWasWhitespace = true
			sb.WriteByte(' ')
		}
	}
}

var allWhitespace = regexp.MustCompile(`^[ \t\n\r]*$`)

func isAllWhitespace(data string) bool {
	return allWhitespace.MatchString(data)
}
