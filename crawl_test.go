package motest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fakeFetcher(req *crawlRequest) *crawlResponse {
	if req.URL == "https://domain.com/" {
		return &crawlResponse{
			body: []byte(`<html><body>
			<a href="https://domain.com/">root</a>
			<a href="/about">about</a>
			</body></html>`),
			statusCode: 200,
			req:        req,
		}
	}
	return &crawlResponse{
		body: []byte(`<html><body>
			some bad html..
			<a href="https://domain.com/">root</a>
			</body></html>`),
		statusCode: 200,
		req:        req,
	}
}

func TestCrawlBasic(t *testing.T) {
	outCh := make(chan *CrawledURL)
	go func() {
		var out []*CrawledURL
		for r := range outCh {
			out = append(out, r)
		}
		assert.Equal(t,
			[]*CrawledURL{
				&CrawledURL{
					URL: "https://domain.com/",
					Nodes: []string{
						"https://domain.com/",
						"https://domain.com/about",
					},
				},
				&CrawledURL{
					URL: "https://domain.com/about",
					Nodes: []string{
						"https://domain.com/",
					},
				},
			},
			out,
		)
	}()

	CrawlWithFetcher(
		CrawlConfig{
			MaxConnections: 1,
			Domain:         "domain.com",
		},
		outCh,
		fakeFetcher,
	)
}
