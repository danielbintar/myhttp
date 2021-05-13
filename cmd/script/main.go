package main

import (
	"flag"
	"fmt"

	"github.com/danielbintar/myhttp/service/hash"
)

func main() {
	var workerCount int
	flag.IntVar(&workerCount, "parallel", 10, "number of worker")

	flag.Parse()

	links := flag.Args()
	resp := hash.Hash(hash.MD5Hash, links, workerCount)
	for k, v := range resp {
		fmt.Printf("%s %s\n", k, v)
	}
}
