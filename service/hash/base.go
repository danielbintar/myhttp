package hash

import "context"

type HashFunc func(datum string) string

func Hash(hashFunc HashFunc, data []string, workerCount int) map[string]string {
	if hashFunc == nil || workerCount < 1 || workerCount > 1_000_000 {
		return nil
	}

	params := make(chan string, len(data))
	resp := make(chan hashResult, len(data))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 1; i <= workerCount; i++ {
		go spawnWorker(ctx, hashFunc, params, resp)
	}

	for _, datum := range data {
		params <- datum
	}

	result := make(map[string]string)
	for i := 1; i <= len(data); i++ {
		r := <-resp
		result[r.datum] = r.result
	}

	return result
}

type hashResult struct {
	datum  string
	result string
}

func spawnWorker(ctx context.Context, hashFunc HashFunc, params chan string, resp chan hashResult) {
	for {
		select {
		case <-ctx.Done():
			return
		case datum := <-params:
			resp <- hashResult{
				datum:  datum,
				result: hashFunc(datum),
			}
		}
	}
}
