# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

build:
	docker compose up --build

up:
	docker compose up

down:
	docker compose down

restart:
	docker compose restart

rebuild:
	docker compose up --build

clean:
	docker stop data
	docker stop cache

	docker rm -v data
	docker rm -v cache

	docker image rm mysql
	docker image rm redis

	rm -rf .dbdata

fmt:
	@gofmt -l -w $(SRC)

lint:
	golint ./...

build:
	go build -o bin/main cmd/api/main.go

run:
	go run main.go

test:
	go test -race -cover ./...

tools:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/kisielk/errcheck
	go get github.com/golang/lint/golint
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html
	go get github.com/tools/godep
	go get github.com/mitchellh/gox

tidy: FORCE
	go mod tidy



