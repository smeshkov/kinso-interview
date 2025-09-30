.PHONY: fmt

lint:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...

fmt:
	gofmt -w=true -s $(find . -type f -name '*.go' -not -path "./vendor/*")

test:
	./_bin/test.sh

run:
	go run cmd/app/main.go

gen:
	go run cmd/generator/main.go