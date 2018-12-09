package motest

import (
	"log"
	"sync"

	"github.com/msf/motest/parse"
)

type crawlRequest struct {
	URL string
}

type crawlResponse struct {
	body       []byte
	err        error
	statusCode int
	req        *crawlRequest
}

// CrawledURL holds URL, all child Nodes referenced in this URL and/or Err if this crawl errored
type CrawledURL struct {
	URL   string
	Nodes []string
	Err   error
}

// CrawledDomainMap holds the complete Domain Graph of URLs of and their child URLs
type CrawledDomainMap struct {
	Domain string
	Root   *CrawledURL
	URLs   map[string]*CrawledURL
}

// CrawlConfig holds the configurable parameters for a Domain Crawl
type CrawlConfig struct {
	Domain         string
	MaxConnections int
}

// Crawl a part of the world wide web according to CrawlConfig
func Crawl(cfg CrawlConfig) *CrawledDomainMap {

	reqsCh := make(chan *crawlRequest, cfg.MaxConnections)
	responsesCh := make(chan *crawlResponse, cfg.MaxConnections)
	crawlCompletedCh := make(chan *CrawledURL, cfg.MaxConnections)

	rootURL := "https://" + cfg.Domain
	issued := make(map[string]struct{})

	// termination condition is tricky, we don't know how many pages we'll have.
	// We use a WaitGroup to signal for every outstanding crawl request we issue
	// we consume them when we get a completed crawl
	outstandingReqs := sync.WaitGroup{}
	endCh := make(chan struct{})
	waitForEnd := func() {
		outstandingReqs.Wait()
		// signal to all workers to end, crawl is complete.
		close(endCh)
	}
	go waitForEnd()

	outstandingReqs.Add(1)
	reqsCh <- &crawlRequest{URL: rootURL}
	issued[rootURL] = struct{}{}

	for i := 0; i < cfg.MaxConnections; i++ {
		go fetcher(reqsCh, responsesCh, endCh)
		go processPage(cfg.Domain, responsesCh, crawlCompletedCh, endCh)
	}

	visited := make(map[string]*CrawledURL)
	for {
		select {
		case <-endCh:
			return &CrawledDomainMap{
				Domain: cfg.Domain,
				Root:   visited[rootURL],
				URLs:   visited,
			}
		case done := <-crawlCompletedCh:
			log.Printf("completed for: %v, code: %v, childs: %v", done.URL, done.Err, len(done.Nodes))
			for _, uri := range done.Nodes {
				if _, ok := issued[uri]; !ok {
					outstandingReqs.Add(1)
					reqsCh <- &crawlRequest{URL: uri}
					issued[uri] = struct{}{}
				}
			}
			outstandingReqs.Done()
			visited[done.URL] = done
		}
	}
}

func fetcher(reqsCh <-chan *crawlRequest, responsesCh chan<- *crawlResponse, endCh <-chan struct{}) {
	for {
		select {
		case req := <-reqsCh:
			log.Printf("req for: %v", req.URL)
			responsesCh <- fetch(req)
		case <-endCh:
			break
		}
	}
}

func processPage(domain string, responsesCh <-chan *crawlResponse, completedCh chan<- *CrawledURL, endCh <-chan struct{}) {

	handleResponse := func(res *crawlResponse, outCh chan<- *CrawledURL) {
		crawled := &CrawledURL{
			URL: res.req.URL,
		}
		if res.statusCode == 200 {
			urls, err := parse.URLs(domain, res.body)
			crawled.Err = err
			crawled.Nodes = urls
		} else {
			crawled.Err = res.err
		}
		// don't block
		go func() { outCh <- crawled }()
	}

	for {
		select {
		case res := <-responsesCh:
			log.Printf("resp for: %v, code: %v", res.req.URL, res.statusCode)
			handleResponse(res, completedCh)
		case <-endCh:
			break
		}
	}
}
