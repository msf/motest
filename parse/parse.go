package parse

import (
)



interface CrawlResultParser {
	parse(crawlResult) []crawlRequest
}
