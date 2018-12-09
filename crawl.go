package motest

type crawlRequest struct {
	URL string
}

type crawlResult struct {
	Body       []byte
	statusCode int
	req        crawlRequest
}

type CrawledURL struct {
	URL   string
	Nodes []*CrawledURL
}

type CrawledDomainMap struct {
	domain string
	Root   *CrawledURL
}

type resultParser interface {
	FindURLs(crawlResult) []crawlRequest
}

type fetcher interface {
	Do(crawlRequest) crawlResult
}

type cache interface {
	HasResult(crawlRequest) bool
	SaveResult(crawlResult) bool
}

type crawlerState struct {
	resultParser resultParser
	fetcher      fetcher
	resultCache  cache

	completedCrawlResults chan crawlResult // inbound queue of pages to process
	pendingCrawlRequest   chan crawlRequest
}

type CrawlEngine interface {
	CrawlDomain(domain string) CrawledDomainMap
	HaveCrawlResult(crawlRequest) bool
}
