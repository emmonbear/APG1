GO := go

build:
	$(GO) build -o readDB ./cmd/ex00/main.go
	$(GO) build -o compareDB ./cmd/ex01/main.go
	$(GO) build -o compareFS ./cmd/ex02/main.go
	

fmt:
	$(GO) fmt ./...

clean:
	rm -f readDB compareDB compareFS

.PHONY: build fmt clean
