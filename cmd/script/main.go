package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/danielbintar/myhttp/service/hash"
)

func main() {
	var workerCount int
	flag.IntVar(&workerCount, "parallel", 1, "number of worker")

	flag.Parse()

	links := flag.Args()
	resp := hash.Hash(md5hash, links, workerCount)
	for k, v := range resp {
		fmt.Printf("%s %s\n", k, v)
	}
}

func md5hash(link string) string {
	hashed := md5.Sum([]byte(link))
	return hex.EncodeToString(hashed[:])
}
