pretty:
	gofmt -s -w .

test:
	go test `go list ./... | grep -v cmd`
