
Goals of this implementation:
 - simple, easy to maintain, deploy and operate
 - single process (and therefore single machine)
 - concurrent to separate concerns cleanly and make scaling easier
	crawling engine (manages the mechanics of this program, some basic state tracking)
	page fetching (this is IO bound, is focuses on related problems)
	page parsing (this is cpu bound, it focuses on URL extraction logic)
 - Limit some hard resources for safety, robustness:
	page fetching:
		limit maximum outstanding TCP connecitons and parallel requests
		this is mostly a internal protection to allow us to be robust in the presence of:
			- request rate limiting on the target domain
			- network flow restrictions on OS or network
			(max tcp connections or file descriptors for example)
	pending URL fetches:
		basic upper bound on memory consumption to avoid running out of mem or swapping.
 - Provide an upper limit on how long a crawl can take.

Non-Goals:
 - nowadays most sites are dynamic and we need a Javascript engine to "render" and really 
	identify the URLs that would be clickable/visible to humans on a browser.
	This crawler doesn't handle this, it behaves as if we were in the good old 2000s.
 - handle network faults very well: this is slightly non-trivial and would require:
	- extensive use of timers and retry-logic (w/ exponential backoff) around:
		- DNS reqs
		- TCP connection pool management
		- network writes and network reads (which might be streamed)
 - handle URL pages whose body doesn't fit in memory
 
 - Distributed Crawler -- This isn't a multi-machine implementation.
	For very large domains this would be unavoidable or we'd never complete the crawl in a reasonable time.
	Additionally this would be needed to work around rate-limiting and other protections
	web-services use to defend against abuse.
	I can expand separately how I'd do this for "google scale" =)