GOLANGCILINT_VERSION=v2.6.1
GOLANGCILINT_ROOT=$$(go env GOPATH)/golangci-lint
GOLANGCILINT_PATH=${GOLANGCILINT_ROOT}/${GOLANGCILINT_VERSION}

.PHONY: tidy test lint gqlgen

tidy:
	go mod tidy

test:
	@mkdir -p .test
	@go test -coverprofile=./.test/coverage.out ./...

setup-lint:
	@mkdir -p ${GOLANGCILINT_ROOT}
	@mkdir -p ${GOLANGCILINT_PATH}
	@if [ -z "$$(ls -A "${GOLANGCILINT_PATH}")" ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "${GOLANGCILINT_PATH}" "${GOLANGCILINT_VERSION}"; \
	fi

lint: setup-lint
	@${GOLANGCILINT_PATH}/golangci-lint run -c tools/.golangci.yml


gqlgen:
	@go run github.com/99designs/gqlgen generate --config tools/gqlgen.yml
