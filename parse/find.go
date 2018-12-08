package parse

import (
	"bytes"

	"github.com/pkg/errors"
	"golang.org/x/net/html"
)

// findURLs present in content
func findURLs(content []byte) ([]string, error) {
	// for now, just find hrefs links, we don't get urls for css and imgs or others..
	hrefs, err := findHrefs(content)
	if err != nil {
		return nil, errors.Wrap(err, "Error on findHrefs for URLs")
	}
	return hrefs, nil

}

// findHrefs uses an html parser to read all <a href=..> links on content
func findHrefs(content []byte) ([]string, error) {
	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to html.Parse() content")
	}
	return extractHrefValues(doc, []string{}), nil
}

func extractHrefValues(n *html.Node, accumulator []string) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				accumulator = append(accumulator, a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		accumulator = extractHrefValues(c, accumulator)
	}
	return accumulator
}
