package hash

import "context"

type HashFunc func(link string) string

func Hash(hashFunc HashFunc, links []string, workerCount int) map[string]string {
	if hashFunc == nil || workerCount < 1 || workerCount > 1_000_000 {
		return nil
	}

	params := make(chan string, 100)
	resp := make(chan hashResult, 100)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 1; i <= workerCount; i++ {
		go spawnWorker(ctx, hashFunc, params, resp)
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

func spawnWorker(ctx context.Context, hashFunc HashFunc, params chan string, resp chan hashResult) {
	for {
		select {
		case <-ctx.Done():
			return
		case link := <-params:
			resp <- hashResult{
				link:   link,
				result: hashFunc(link),
			}
		}
	}
}
