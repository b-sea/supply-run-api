.PHONY: tidy test lint gqlgen certs

tidy:
	go mod tidy

cert:
	@mkdir -p ./.cert
	@openssl genrsa -out ./.cert/id_rsa 4096
	@openssl rsa -in ./.cert/id_rsa -pubout -out ./.cert/id_rsa.pub

test:
	@mkdir -p .test
	@go test -coverprofile=./.test/coverage.out ./...

lint:
	@golangci-lint run -c tools/.golangci.yml

gqlgen:
	@go run github.com/99designs/gqlgen generate --config tools/gqlgen.yml

certs:
	@mkdir -p .certs
	@openssl genrsa -out .certs/key.pem 2048
	@openssl rsa -in .certs/key.pem -pubout > .certs/key.pub