package parse

import (
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

// TODO: this function is doing too many things. split out better how handle correctly the variety of URL types
func absoluteURLsForDomain(domain, baseURL string, URLs []string) ([]string, error) {
	var out []string
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrapf(err, "bad baseURL: %v", baseURL)
	}
	for _, u := range filterURLs(domain, convertToURLs(URLs)) {
		u.Fragment = "" // ignore the fragment.
		var uStr string
		if u.IsAbs() {
			uStr = u.String()
		} else {
			if strings.HasPrefix(u.Path, "/") {
				uStr = fmt.Sprintf("%s://%s%s", base.Scheme, base.Host, u.String())
			} else if strings.HasPrefix(u.Path, "../") && !strings.HasSuffix(baseURL, "/") {
				// handle path walking like: "site.com/careers../blog/a-blog-post"
				uStr = fmt.Sprintf("%s/%s", baseURL, u.String())
			} else {
				uStr = fmt.Sprintf("%s%s", baseURL, u.String())
			}
		}
		out = append(out, uStr)
	}
	return uniq(out), nil
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

func uniq(in []string) []string {
	good := make(map[string]struct{})
	for _, i := range in {
		good[i] = struct{}{}
	}
	var out []string
	for i := range good {
		out = append(out, i)
	}
	sort.Strings(out)
	return out
}
