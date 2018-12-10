package motest

import "fmt"

type printElem struct {
	page   *CrawledURL
	indent int
}

// PrintStream prints crawled URLs as they complete
func PrintStream(finishedCh <-chan *CrawledURL) {
	for page := range finishedCh {
		page.print()
	}
}

func (curl *CrawledURL) print() {
	fmt.Printf("\n%s\n", curl.URL)
	for _, u := range curl.Nodes {
		fmt.Printf("    -> %s\n", u)
	}
}

// Print to stdout a Domain Map of URLs in breath-first-order
func Print(crawled *CrawledDomainMap) {
	q := []printElem{printElem{page: crawled.Root, indent: 0}}
	printed := make(map[string]struct{})
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if _, ok := printed[curr.page.URL]; ok {
			continue
		}
		printed[curr.page.URL] = struct{}{}
		fmt.Printf("%s%s\n", indent(curr.indent), curr.page.URL)
		for _, u := range curr.page.Nodes {
			fmt.Printf("%s -> %s\n", indent(curr.indent), u)
			q = append(q, printElem{
				page:   crawled.URLs[u],
				indent: curr.indent + 1,
			})
		}
	}
}

func indent(count int) string {
	var s string
	for i := 0; i < count; i++ {
		s = fmt.Sprintf("  %s", s)
	}
	return s
}
