package url_test

import (
	"testing"

	"github.com/danielbintar/myhttp/service/url"
)

type testCaseValid struct {
	link   string
	output bool
}

func generateValidTestCase() []testCaseValid {
	return []testCaseValid{
		{
			link:   "google.com",
			output: false,
		},
		{
			link:   "google.com/",
			output: false,
		},
		{
			link:   "http://google.com",
			output: true,
		},
		{
			link:   "https://mail.google.com",
			output: true,
		},
		{
			link:   "http:///mail.google.com",
			output: false,
		},
		{
			link:   "aa",
			output: false,
		},
		{
			link:   "aa/aaa",
			output: false,
		},
		{
			link:   "//aa.aa",
			output: false,
		},
	}
}

func TestValid(t *testing.T) {
	testCases := generateValidTestCase()

	for _, testCase := range testCases {
		resp := url.Valid(testCase.link)
		if testCase.output != resp {
			t.Fatalf("Test %s is %t, should be %t", testCase.link, resp, testCase.output)
		}
	}
}
