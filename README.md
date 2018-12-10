# Crawl a web domain

Crawling a web domain is similar to walking a graph.
This implementation resembles a breath-first-search crawler, but because it is concurrent it isn't  strictly "breath-first".

## Goals of this implementation:

 - simple, easy to maintain, deploy and operate
 - single process (and therefore single machine)
 - concurrent to separate concerns cleanly and make scaling easier
   - Crawl: crawling engine (manages the mechanics of this program, some basic state tracking)
   - fetcher: page fetching (this is IO bound, is focuses on related problems)
   - parser: page parsing (this is cpu bound, it focuses on URL extraction logic)
 - Limit some hard resources for safety, robustness:
	- Limit maximum in flight TCP connections and parallel requests.
	- Limit pending URL fetches:
		basic upper bound on memory consumption to avoid running out of mem or swapping.
    - Print out domain map incrementally to avoid holding the entire graph in memory.
 

## Non-Goals


 - Nowadays most sites are dynamic and we need a Javascript engine to "render" and really 
	identify the URLs that would be clickable/visible to humans on a browser.
	This crawler doesn't handle this, it behaves as if we were in the good old 2000s.
 - Handle network faults very well: this is slightly non-trivial and would require:
	- extensive use of timers and retry-logic (w/ exponential backoff) around:
		- DNS reqs
		- TCP connection pool management
		- network writes and network reads (which might be streamed)
 - Handle URL pages whose body doesn't fit in memory
 - Handle other URLs besides "\<a href='*' />"
 

## Testing

The crawler and page parser have tests. The 'fetcher' component doesn't have tests.
To make this code production worthy a better 'fetcher' that handles errors is needed, at that time a good test suite for the component would be done.


## Time and Space Complexity

Time complexity is O(N) where N is number of pages. The bottleneck is going to be page fetching IO rates and not cpu time.

Space complexity: O(1) because it is bound to the number of pending and outstanding page requests. 

## Distributed Implementation

Distributed Crawler -- This isn't a multi-machine implementation.
For very large domains this would be unavoidable or we'd never complete the crawl in a reasonable time.
Additionally this would be needed to work around rate-limiting and other protections
web-services use to defend against abuse.

I can expand in person how I'd do this for "google scale" =)

The simpler way would be to maintain a singleton crawling coordinator using distributed datastructures for its state.