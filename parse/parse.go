package parse

import (
	"github.com/pkg/errors"
)

// URLs returns list of request URLs for the same hostname as baseURL has present in responseBody
func URLs(domain, baseURL string, responseBody []byte) ([]string, error) {
	urls, err := findURLs(responseBody)
	if err != nil {
		return nil, errors.Wrap(err, "failure on findURLs()")
	}

	return absoluteURLsForDomain(domain, baseURL, urls)
}
