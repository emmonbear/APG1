GO := go

build:
	$(GO) build -o readDB ./cmd/ex00/main.go
	$(GO) build -o compareDB ./cmd/ex01/main.go
	$(GO) build -o compareFS ./cmd/ex02/main.go
	

fmt:
	$(GO) fmt ./...

test:
	go test ./... -coverprofile=c.out

gcov: test
	go tool cover -html=c.out -o coverage.html
	xdg-open coverage.html

clean:
	rm -f readDB compareDB compareFS c.out coverage.html

.PHONY: build fmt clean
