package parse

import (
	"fmt"
	"net/url"
	"strings"
)

func absoluteURLsForDomain(domain string, URLs []string) ([]string, error) {
	var out []string
	for _, u := range filterURLs(domain, convertToURLs(URLs)) {
		if !u.IsAbs() {
			out = append(out, fmt.Sprintf("https://%s%s", domain, u.String()))
		} else {
			out = append(out, u.String())
		}
	}
	return out, nil
}

func filterURLs(domain string, urls []*url.URL) []*url.URL {
	var out []*url.URL
	for _, u := range urls {
		if u.IsAbs() {
			if strings.HasSuffix(u.Hostname(), domain) {
				out = append(out, u)
			}
		} else {
			out = append(out, u)
		}
	}
	return out
}

func convertToURLs(URLs []string) []*url.URL {
	var out []*url.URL
	for _, t := range URLs {
		u, err := url.Parse(t)
		if err != nil {
			continue
		}
		out = append(out, u)
	}
	return out
}
