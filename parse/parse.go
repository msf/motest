package parse

import (
	"github.com/pkg/errors"
)

// URLs returns list of request URLs for domain present in responseBody
func URLs(domain string, responseBody []byte) ([]string, error) {
	urls, err := findURLs(responseBody)
	if err != nil {
		return nil, errors.Wrap(err, "failure on findURLs()")
	}

	return filterByDomain(domain, urls)
}
