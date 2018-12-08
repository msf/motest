package parse

import (
	"github.com/pkg/errors"
)

// findURLs present in content (possibly an html page) that match a given domain
func findURLs(content []byte, domain string) ([]string, error) {
	// for now, just find hrefs links, other urls like cssand imgs:qa!
	hrefs, err := findHrefs(content)
	if err != nil {
		return nil, errors.Wrapf(err, "Error on findHrefs for URLs")
	}

	return filterByDomain(domain, hrefs)
}

const baseHrefPrefix  = 'href="'

// FindHrefs 
func FindHrefs(content []byte) ([]string, error) {

}