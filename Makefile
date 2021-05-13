compile:
	go build -o myhttp cmd/script/main.go

pretty:
	gofmt -s -w .

test:
	go test `go list ./... | grep -v cmd`
