package motest

import "fmt"

type printElem struct {
	page   *CrawledURL
	indent int
}

// Print to stdout a Domain Map of URLs in breath-first-order
func Print(crawled *CrawledDomainMap) {
	q := []printElem{printElem{page: crawled.Root, indent: 0}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		fmtStr := fmt.Sprintf("%%%ds\n", curr.indent)
		fmt.Printf(fmtStr, curr.page.URL)
		for _, u := range curr.page.Nodes {
			fmt.Printf(fmtStr, " -> "+u)
			q = append(q, printElem{
				page:   crawled.URLs[u],
				indent: curr.indent + 1,
			})
		}
	}
}
