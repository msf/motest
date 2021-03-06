package motest

import (
	"fmt"
	"log"
	"sync"

	"github.com/msf/motest/parse"
)

// CrawledURL holds URL, all child Nodes referenced in this URL and/or Err if this crawl errored
type CrawledURL struct {
	URL   string
	Nodes []string
	Err   error
}

// CrawlConfig holds the configurable parameters for a Domain Crawl
type CrawlConfig struct {
	Domain         string
	MaxConnections int
}

// crawlRequest is inner type for issuing crawl reqs
type crawlRequest struct {
	URL string
}

type crawlResponse struct {
	body       []byte
	err        error
	statusCode int
	req        *crawlRequest
}

// Crawl a part of the world wide web according to CrawlConfig
// 	Note that crawl will write to outCh as pages are crawled their maps are found, the user
//  is responsible for consuming from outCh so that Crawl can make progress.
func Crawl(cfg CrawlConfig, outCh chan<- *CrawledURL) {
	CrawlWithFetcher(cfg, outCh, fetch)
}

// CrawlWithFetcher is a Crawler that uses the provided fetchFn to request URLs
func CrawlWithFetcher(cfg CrawlConfig, outCh chan<- *CrawledURL, fetchFn func(r *crawlRequest) *crawlResponse) {
	reqsCh := make(chan *crawlRequest, cfg.MaxConnections)
	responsesCh := make(chan *crawlResponse, cfg.MaxConnections)
	crawlCompletedCh := make(chan *CrawledURL, cfg.MaxConnections)

	rootURL := fmt.Sprintf("https://%s/", cfg.Domain)
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

	outstandingReqs.Add(1)
	reqsCh <- &crawlRequest{URL: rootURL}
	issued[rootURL] = struct{}{}

	go waitForEnd()
	for i := 0; i < cfg.MaxConnections; i++ {
		go fetcher(fetchFn, reqsCh, responsesCh, endCh)
		go parser(cfg.Domain, responsesCh, crawlCompletedCh, endCh)
	}

	for {
		select {
		case <-endCh:
			return
		case done := <-crawlCompletedCh:
			log.Printf("completed: %v, err: %v, Nodes: %v", done.URL, done.Err, len(done.Nodes))
			for _, uri := range done.Nodes {
				if _, ok := issued[uri]; !ok {
					outstandingReqs.Add(1)
					reqsCh <- &crawlRequest{URL: uri}
					issued[uri] = struct{}{}
				}
			}
			outstandingReqs.Done()
			outCh <- done
		}
	}
}

func fetcher(fetchFn func(r *crawlRequest) *crawlResponse, reqsCh <-chan *crawlRequest, responsesCh chan<- *crawlResponse, endCh <-chan struct{}) {
	for {
		select {
		case req := <-reqsCh:
			responsesCh <- fetchFn(req)
		case <-endCh:
			break
		}
	}
}

func parser(domain string, responsesCh <-chan *crawlResponse, completedCh chan<- *CrawledURL, endCh <-chan struct{}) {

	handleResponse := func(res *crawlResponse, outCh chan<- *CrawledURL) {
		crawled := &CrawledURL{
			URL: res.req.URL,
		}
		if res.statusCode == 200 {
			urls, err := parse.URLs(domain, res.req.URL, res.body)
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
			handleResponse(res, completedCh)
		case <-endCh:
			break
		}
	}
}
