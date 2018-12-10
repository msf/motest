package main

import (
	"flag"
	"log"

	"github.com/msf/motest"
)

func main() {
	maxConnsPtr := flag.Int("MaxConns", 10, "MaxConnections")
	domainPtr := flag.String("domain", "monzo.com", "domain to crawl")
	flag.Parse()

	cfg := motest.CrawlConfig{
		MaxConnections: *maxConnsPtr,
		Domain:         *domainPtr,
	}
	log.Printf("motest, config: %+v", cfg)
	out := make(chan *motest.CrawledURL)
	go motest.Crawl(cfg, out)
	motest.PrintStream(out)
}
