

.PHONY: test
test:
	go test -v -race -coverprofile=coverage.out ./...

.PHONY: cover
cover:
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
