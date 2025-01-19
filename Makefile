tidy:
	go mod tidy

test:
	@mkdir -p .test
	@go test -coverprofile=./.test/coverage.out ./...

lint:
	@golangci-lint run -c tools/.golangci.yml
