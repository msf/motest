package motest


struct crawlRequest {
	URL string
}

struct crawlResult {
	Body []byte
	statusCode int
	FetchedAt time.Time
	req crawlRequest
}

struct CrawledURL {
	URL string
	FetchedAt time.Time
	Nodes []CrawledURL
}

struct CrawledDomainMap {
	domain string
	Root CrawledURL
}


func CrawlDomain(zdomaind string) CrawledDomainMap {

}

/*

For a given crawlRequest we do a full search of all pages of that subdomain.
Print a Tree-View (or more accurately a DAG) of the pages found.
*/

struct resultProcessor {
 	completedCrawlResults chan  // inbound queue of pages to process
	pendingCrawlRequest chan
}

interface CrawlEngine {
	CrawlDomain(domain string) 
	HaveCrawlResult(crawlRequest) bool

}