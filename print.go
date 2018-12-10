package motest

import "fmt"

// PrintStream prints crawled URLs as they complete
func PrintStream(outCh <-chan *CrawledURL) {
	for page := range outCh {
		page.print()
	}
}

func (curl *CrawledURL) print() {
	fmt.Printf("\n%s\n", curl.URL)
	for _, u := range curl.Nodes {
		fmt.Printf("    -> %s\n", u)
	}
}
