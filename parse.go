package motest

import (
	"github.com/pkg/errors"
)



interface CrawlResultParser {
	parse(crawlResult) []crawlRequest
}

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

// findHrefs 
func findHrefs(content []byte) ([]string, error) {

}

// filterByDomain will exclude from URLs entries that aren't from domain.
func filterByDomain(domain string, URLs []string) ([]string, error) {

}
