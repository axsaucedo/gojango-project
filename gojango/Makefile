
test:
	go test -v ./...

cover:
	go test -coverprofile=coverage.out ./... && \
		go tool cover -html=coverage.out -o coverage.html

coverage:
	go test -cover ./...

build_cli:
	go build -o ../gojango-myapp/gojango ./cmd/cli
