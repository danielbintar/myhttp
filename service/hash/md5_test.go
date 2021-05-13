package hash_test

import (
	"testing"

	"github.com/danielbintar/myhttp/service/hash"
)

type testCaseMD5Hash struct {
	link   string
	output string
}

func generateHashMD5TestCase() []testCaseMD5Hash {
	return []testCaseMD5Hash{
		{
			link:   "google.com",
			output: "1d5920f4b44b27a802bd77c4f0536f5a",
		},
		{
			link:   "facebook.com",
			output: "2343ec78a04c6ea9d80806345d31fd78",
		},
		{
			link:   "yahoo.com",
			output: "50cd1a9a183758039b0841aa738c3f0b",
		},
		{
			link:   "yandex.com",
			output: "31aa70fc8589c52a763a2df36f304d28",
		},
		{
			link:   "twitter.com",
			output: "7905d1c4e12c54933a44d19fcd5f9356",
		},
		{
			link:   "reddit.com",
			output: "1fd7de7da0fce4963f775a5fdb894db5",
		},
	}
}

func TestMD5Hash(t *testing.T) {
	testCases := generateHashMD5TestCase()

	for _, testCase := range testCases {
		resp := hash.MD5Hash(testCase.link)
		if testCase.output != resp {
			t.Fatalf("Test %s is %s, should be %s", testCase.link, resp, testCase.output)
		}
	}
}
