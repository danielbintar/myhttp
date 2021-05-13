package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/danielbintar/myhttp/service/hash"
	"github.com/danielbintar/myhttp/service/url"
)

func main() {
	var workerCount int
	flag.IntVar(&workerCount, "parallel", 10, "number of worker")

	flag.Parse()

	links := flag.Args()
	validlinks := make([]string, 0)
	invalidLinks := make([]string, 0)
	for _, link := range links {
		if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
			link = "http://" + link
		}

		if url.Valid(link) {
			validlinks = append(validlinks, link)
		} else {
			invalidLinks = append(invalidLinks, link)
		}
	}

	resp := hash.Hash(hash.MD5Hash, validlinks, workerCount)
	for k, v := range resp {
		fmt.Printf("%s %s\n", k, v)
	}

	for _, link := range invalidLinks {
		fmt.Printf("%s is not a valid link\n", link)
	}
}
