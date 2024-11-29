GO := go

build:
	$(GO) build -o readDB ./cmd/ex00/main.go
	$(GO) build -o compareDB ./cmd/ex01/main.go
	$(GO) build -o compareFS ./cmd/ex02/main.go
	

fmt:
	$(GO) fmt ./...

test: clean
	go test ./... -coverprofile=c.out

gcov: test
	go tool cover -html=c.out -o coverage.html
	xdg-open coverage.html

dvi:
	godoc -http=:6060 &
	@echo "Documentation server running at http://localhost:6060/pkg"
	@echo "Press Ctrl+C to stop the server"

clean:
	rm -f readDB compareDB compareFS c.out coverage.html

.PHONY: build fmt clean
