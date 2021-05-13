package hash_test

import (
	"fmt"
	"testing"

	"github.com/danielbintar/myhttp/service/hash"
)

type testCaseHash struct {
	workerCount int
	links       []string
	output      map[string]string
	hashFunc    hash.HashFunc

	name string
}

func hashFunc(link string) string {
	if link == "google.com" {
		return "encodedgoogle"
	}

	return "default"
}

func TestHash(t *testing.T) {
	testCases := []testCaseHash{
		{
			workerCount: 5,
			name:        "no hash function",
		},
		{
			workerCount: 0,
			name:        "low worker count",
			hashFunc:    hashFunc,
		},
		{
			workerCount: 1_000_001,
			name:        "high worker count",
			hashFunc:    hashFunc,
		},
		{
			workerCount: 5,
			name:        "normal",
			hashFunc:    hashFunc,
			links:       []string{"google.com", "facebook.com"},
			output: map[string]string{
				"google.com":   "encodedgoogle",
				"facebook.com": "default",
			},
		},
	}

	for _, testCase := range testCases {
		resp := hash.Hash(testCase.hashFunc, testCase.links, testCase.workerCount)

		if testCase.output == nil {
			if resp != nil {
				t.Fatalf("Test %s resp is not nil", testCase.name)
			}
			continue
		}

		if resp == nil {
			t.Fatalf("Test %s resp is nil", testCase.name)
		}

		if len(testCase.output) != len(resp) {
			t.Fatalf("Test %s resp len is %d, should be %d", testCase.name, len(resp), len(testCase.output))
		}

		for k, v := range testCase.output {
			if resp[k] != v {
				t.Fatalf("Test %s hash %s is %s, should be %s", testCase.name, k, resp[k], v)
			}
		}
	}

	fmt.Println(testCases)
}
