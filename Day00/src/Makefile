GO   := go
MAIN := main.go
EXE  := main

build: 
	${GO} build -o ${EXE} ${MAIN}

clean:
	rm -f ${EXE}

run:
	${GO} run ${MAIN}

fmt:
	${GO} fmt ./...

vet:
	${GO} vet ./...

tidy:
	${GO} mod tidy

test:
	${GO} test ./...

.PHONY: build vet fmt clean run test
