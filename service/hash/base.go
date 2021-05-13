package hash

type HashFunc func(link string) string

func Hash(hashFunc HashFunc, links []string, workerCount int) map[string]string {
	if hashFunc == nil || workerCount < 1 || workerCount > 1_000_000 {
		return nil
	}

	params := make(chan string)
	resp := make(chan hashResult)
	quit := make(chan bool)
	for i := 1; i <= workerCount; i++ {
		go spawnWorker(hashFunc, params, resp, quit)
	}

	for _, link := range links {
		params <- link
	}

	result := make(map[string]string)
	for i := 1; i <= len(links); i++ {
		r := <-resp
		result[r.link] = r.result
	}

	return result
}

type hashResult struct {
	link   string
	result string
}

func spawnWorker(hashFunc HashFunc, params chan string, resp chan hashResult, quit chan bool) {
	for {
		select {
		case link := <-params:
			resp <- hashResult{
				link:   link,
				result: hashFunc(link),
			}
		case <-quit:
			return
		}
	}
}
